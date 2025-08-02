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
		fileHeader, err := cc.FormFile(fmt.Sprintf("file_%d", i))
		// Use the idiomatic way to check for the end of files
		if err != nil {
			if err == http.ErrMissingFile {
				break
			}
			return newMedias, err
		}
		i += 1

		openFile, err := fileHeader.Open()
		if err != nil {
			return newMedias, err
		}
		defer openFile.Close()

		name := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(fileHeader.Filename))
		err = cc.Storage.Write(store.MEDIA, openFile, name)
		if err != nil {
			return newMedias, err
		}

		mediaType := mime.TypeByExtension(filepath.Ext(fileHeader.Filename))
		newMedia := model.Media{
			Name:        fileHeader.Filename,
			Path:        name,
			ContentType: mediaType,
			UserId:      cc.Claims.Subject,
		}
		err = cc.MediaCrud.Insert(&newMedia)
		newMedias = append(newMedias, newMedia)
	}

	return newMedias, nil
}
