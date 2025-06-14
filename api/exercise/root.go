package exercise

import (
	"gym-map/api"
	"gym-map/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PostModel[exercisePostRequest](cc, cc.ExerciseCrud)
}

func Patch(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PatchModel[exercisePatchRequest](cc, cc.ExerciseCrud)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.ExerciseCrud)
}

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[exerciseGetRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	exercises := []model.Exercise{}
	if params.MachineId == nil {
		exercises, err = cc.ExerciseCrud.Get()
		if err != nil {
			return err
		}
	} else {
		exercises, err = cc.ExerciseCrud.GetByMachineId(*params.MachineId)
		if err != nil {
			return err
		}
	}

	if exercises == nil {
		exercises = []model.Exercise{}
	}

	return cc.JSON(http.StatusOK, exercises)
}
