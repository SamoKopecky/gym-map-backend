package api

import (
	"fmt"
	"gym-map/config"
	"gym-map/fetcher"
	fileio "gym-map/file_io"
	"gym-map/model"
	"gym-map/schema"
	"gym-map/service"
	"gym-map/store"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func BindParams[T any](c echo.Context) (T, error) {
	params := *new(T)
	if err := c.Bind(&params); err != nil {
		return params, err
	}
	return params, nil
}

type DbContext struct {
	echo.Context

	MachineCrud     store.Machine
	ExerciseCrud    store.Exercise
	InstructionCrud store.Instruction
	MediaCrud       store.Media

	FloorMapCrud fileio.FloorMap

	IAMFetcher fetcher.IAM

	InstructionService service.Instruction
	MediaService       service.Media
	UserService        service.User

	Claims *schema.JwtClaims

	Config config.Config
}

func (c DbContext) BadRequest(err error) error {
	errStr := fmt.Sprint(err)
	// TODO: log error too
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters", "reason": errStr})
}

func DerefString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func DerefInt(ptr *int) int {
	if ptr == nil {
		return 0
	}
	return *ptr

}

func CreateFilesFromRequest(cc *DbContext) (newMedias []model.Media, err error) {
	i := 0
	for {
		file, err := cc.FormFile(fmt.Sprintf("file_%d", i))
		// If file not found, there are no more files, break the loop
		if file == nil {
			break
		} else {
			i += 1
		}

		if err != nil {
			return newMedias, err
		}

		fileId, err := fileio.SaveFile(file, cc.Config.MediaFileRepository)
		if err != nil {
			return newMedias, err
		}

		mediaType := mime.TypeByExtension(filepath.Ext(file.Filename))
		newMedia := model.Media{
			Name:        file.Filename,
			Path:        fileId,
			ContentType: mediaType,
			UserId:      cc.Claims.Subject,
		}
		// Create record in media table
		err = cc.MediaCrud.Insert(&newMedia)
		newMedias = append(newMedias, newMedia)
	}

	return
}
