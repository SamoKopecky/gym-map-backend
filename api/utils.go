package api

import (
	"fmt"
	"gym-map/config"
	"gym-map/fetcher"
	"gym-map/model"
	"gym-map/schema"
	"gym-map/service"
	"gym-map/storage"
	"gym-map/store"
	"io"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/google/uuid"
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
	CategoryCrud    store.Category
	PropertyCrud    store.Property

	Storage      store.FileStorage
	FloorMapCrud storage.FloorMap

	IAMFetcher fetcher.IAM

	InstructionService service.Instruction
	MediaService       service.Media
	UserService        service.User
	CategoryService    service.Category

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

		openFile, err := file.Open()
		if err != nil {
			return newMedias, err
		}
		defer openFile.Close()

		data, err := io.ReadAll(openFile)
		if err != nil {
			return newMedias, err
		}

		name := uuid.New().String()
		err = cc.Storage.Write(store.MEDIA, data, name)
		if err != nil {
			return newMedias, err
		}

		mediaType := mime.TypeByExtension(filepath.Ext(file.Filename))
		newMedia := model.Media{
			Name:        file.Filename,
			Path:        name,
			ContentType: mediaType,
			UserId:      cc.Claims.Subject,
		}
		// Create record in media table
		err = cc.MediaCrud.Insert(&newMedia)
		newMedias = append(newMedias, newMedia)
	}

	return
}
