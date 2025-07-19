package instruction

import (
	"gym-map/api"
	"gym-map/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func PostMedia(c echo.Context) error {
	cc := c.(*api.DbContext)
	if hasPermision, err := checkInstructionPermisions(cc); err != nil {
		return err
	} else if !hasPermision {
		return cc.NoContent(http.StatusForbidden)
	}

	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	params, err := api.BindParams[instructionMediaPostRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	var newMedias []model.Media
	if params.YoutubeVideoId == nil || params.Name == nil {
		newMedias, err = api.CreateFilesFromRequest(cc)
		if err != nil {
			return err
		}
	} else {
		newMedia := model.NewYoutubeMedia(*params.YoutubeVideoId, cc.Claims.Subject, *params.Name)
		err := cc.MediaCrud.Insert(&newMedia)
		if err != nil {
			return err
		}
		newMedias = append(newMedias, newMedia)
	}

	// Update instructions table
	newMediaIds := make([]int, len(newMedias))
	for i, newMedia := range newMedias {
		newMediaIds[i] = newMedia.Id
	}
	err = cc.InstructionCrud.SaveMedia(id, newMediaIds)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusCreated, instructionPostResponse{MediaIds: newMediaIds})
}
