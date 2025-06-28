package instruction

import (
	"gym-map/api"
	"gym-map/schema"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func checkOwner(cc *api.DbContext) (bool, error) {
	if isAdmin := cc.Claims.IsAdmin(); isAdmin {
		return true, nil
	}

	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return false, cc.BadRequest(err)
	}

	if isOwned, err := cc.InstructionService.IsTrainerOwned(cc.Claims.Subject, id); err != nil {
		return false, err
	} else if isOwned {
		return true, nil
	}

	return false, nil
}

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[instructionGetRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	instructions := []schema.Instruction{}
	if params.ExerciseId != nil {
		instructions, err = cc.InstructionService.GetByExerciseId(*params.ExerciseId)
		if err != nil {
			return err
		}
	} else if params.UserId != nil {
		// TODO: Where is this used
		instructions, err = cc.InstructionService.GetByUserId(*params.UserId)
		if err != nil {
			return err
		}
	} else {
		instructions, err = cc.InstructionService.Get()
		if err != nil {
			return err
		}
	}

	if instructions == nil {
		instructions = []schema.Instruction{}
	}

	return cc.JSON(http.StatusOK, instructions)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[instructionPostRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	newModel := params.ToNewModel()
	userInstruction, err := cc.InstructionService.Insert(&newModel)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, userInstruction)
}

func Patch(c echo.Context) error {
	cc := c.(*api.DbContext)
	if isOwned, err := checkOwner(cc); err != nil {
		return err
	} else if !isOwned {
		return cc.NoContent(http.StatusForbidden)
	}

	return api.PatchModel[instructionPatchRequest](cc, cc.InstructionCrud)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	if isOwned, err := checkOwner(cc); err != nil {
		return err
	} else if !isOwned {
		return cc.NoContent(http.StatusForbidden)
	}

	return api.DeleteModel(cc, cc.InstructionCrud)
}
