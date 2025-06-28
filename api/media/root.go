package media

import (
	"gym-map/api"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetMedia(c echo.Context) error {
	cc := c.(*api.DbContext)

	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	mediaMetadata, err := cc.MediaCrud.GetById(id)
	if err != nil {
		return cc.BadRequest(err)
	}

	videoPath := filepath.Join(cc.Config.MediaFileRepository, mediaMetadata.DiskFileName)
	file, err := os.Open(videoPath)
	defer file.Close()

	fileInfo, _ := file.Stat()
	http.ServeContent(c.Response().Writer, c.Request(), fileInfo.Name(), fileInfo.ModTime(), file)
	return nil
}

func GetMetadata(c echo.Context) error {
	cc := c.(*api.DbContext)

	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	mediaMetadata, err := cc.MediaCrud.GetById(id)
	if err != nil {
		return cc.BadRequest(err)
	}

	return cc.JSON(http.StatusOK, mediaMetadata)
}
