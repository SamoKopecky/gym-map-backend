package instruction

import (
	"gym-map/api"
	"net/http"
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

	newMedia, err := api.CreateFileFromRequest(cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	// Update instructions table
	err = cc.InstructionCrud.SaveFile(id, newMedia.Id)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusCreated, instructionPostResponse{MediaId: newMedia.Id})
}
