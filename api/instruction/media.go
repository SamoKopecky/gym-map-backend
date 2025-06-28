package instruction

import (
	"gym-map/api"
	"gym-map/media"
	"gym-map/model"
	"mime"
	"net/http"
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

	mediaType := mime.TypeByExtension(filepath.Ext(file.Filename))
	newMedia := model.Media{
		OriginalFileName: file.Filename,
		DiskFileName:     fileId,
		ContentType:      mediaType,
	}
	// TODO: Make a service function
	// Create record in media table
	err = cc.MediaCrud.Insert(&newMedia)

	// Update instructions table
	err = cc.InstructionCrud.SaveFile(id, newMedia.Id)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusCreated, instructionPostResponse{MediaId: newMedia.Id})
}
