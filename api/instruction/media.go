package instruction

import (
	"gym-map/api"
	"gym-map/utils"
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

	newMedias, err := api.CreateFilesFromRequest(cc)
	utils.PrettyPrint(newMedias)
	if err != nil {
		return cc.BadRequest(err)
	}

	// Update instructions table
	newMediaIds := make([]int, len(newMedias))
	for i, newMedia := range newMedias {
		newMediaIds[i] = newMedia.Id
	}

	err = cc.InstructionCrud.SaveFiles(id, newMediaIds)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusCreated, instructionPostResponse{MediaIds: newMediaIds})
}
