package instruction

import (
	"gym-map/api"
	"gym-map/media"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
)

func PostMedia(c echo.Context) error {
	cc := c.(*api.DbContext)
	if isOwned, err := checkOwner(cc); err != nil {
		return err
	} else if !isOwned {
		return cc.NoContent(http.StatusForbidden)
	}

	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	fileId, err := media.SaveFile(file, cc.Config.MediaFileRepository)
	if err != nil {
		return err
	}

	err = cc.InstructionCrud.CreateFile(id, fileId, file.Filename)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusCreated)
}

func GetMedia(c echo.Context) error {
	cc := c.(*api.DbContext)

	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	model, err := cc.InstructionCrud.GetById(id)
	if err != nil {
		return cc.BadRequest(err)
	}
	if model.FileId == nil {
		return cc.NoContent(http.StatusNotFound)
	}

	fileData, err := os.ReadFile(filepath.Join(cc.Config.MediaFileRepository, *model.FileId))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not read file")
	}

	mediaType := mime.TypeByExtension(filepath.Ext(*model.FileName))
	return cc.Blob(http.StatusOK, mediaType, fileData)
}
