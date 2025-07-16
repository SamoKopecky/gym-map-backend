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
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	http.ServeContent(c.Response().Writer, c.Request(), fileInfo.Name(), fileInfo.ModTime(), file)
	return nil
}

func GetMetadataMany(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[MediaGetRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	mediaMetadatas, err := cc.MediaCrud.GetByIds(params.Ids)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, mediaMetadatas)
}
