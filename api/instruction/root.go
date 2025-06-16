package instruction

import (
	"gym-map/api"
	"gym-map/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[instructionGetRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	instructions := []model.Instruction{}
	if params.ExerciseId != nil {
		instructions, err = cc.InstructionCrud.GetByExerciseId(*params.ExerciseId)
		if err != nil {
			return err
		}
	} else if params.UserId != nil {
		instructions, err = cc.InstructionCrud.GetByUserId(*params.UserId)
		if err != nil {
			return err
		}
	} else {
		instructions, err = cc.InstructionCrud.Get()
		if err != nil {
			return err
		}
	}

	if instructions == nil {
		instructions = []model.Instruction{}
	}

	return cc.JSON(http.StatusOK, instructions)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PostModel[instructionPostRequest](cc, cc.InstructionCrud)
}

func Patch(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PatchModel[instructionPatchRequest](cc, cc.InstructionCrud)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.InstructionCrud)
}
