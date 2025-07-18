package media

import (
	"gym-map/api"
	"gym-map/model"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
)

func checkMediaPermision(cc *api.DbContext) (bool, error) {
	// TODO: Make generic
	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return false, cc.BadRequest(err)
	}

	if isOwned, err := cc.MediaService.IsTrainerOwned(cc.Claims.Subject, id); err != nil {
		return false, err
	} else {
		return api.HasPermisions(cc, isOwned), nil
	}
}

func GetMedia(c echo.Context) error {
	cc := c.(*api.DbContext)

	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	mediaMetadata, err := cc.MediaCrud.GetById(id)
	if err != nil {
		return err
	}

	if mediaMetadata.ContentType == model.YOUTUBE_CONTENT_TYPE {
		return cc.NoContent(http.StatusNoContent)
	}

	videoPath := filepath.Join(cc.Config.MediaFileRepository, mediaMetadata.Path)
	file, err := os.Open(videoPath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	http.ServeContent(c.Response().Writer, c.Request(), fileInfo.Name(), fileInfo.ModTime(), file)
	return nil
}

func DeleteMedia(c echo.Context) error {
	cc := c.(*api.DbContext)

	if hasPermision, err := checkMediaPermision(cc); err != nil {
		return err
	} else if !hasPermision {
		return cc.NoContent(http.StatusForbidden)
	}

	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	mediaMetadata, err := cc.MediaCrud.GetById(id)
	if err != nil {
		return err
	}
	err = cc.MediaCrud.Delete(id)
	if err != nil {
		return err
	}

	if mediaMetadata.ContentType != model.YOUTUBE_CONTENT_TYPE {
		mediaPath := filepath.Join(cc.Config.MediaFileRepository, mediaMetadata.Path)
		err = os.Remove(mediaPath)
		if err != nil {
			return err
		}
	}

	return cc.NoContent(http.StatusNoContent)
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
