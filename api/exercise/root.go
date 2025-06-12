package exercise

import (
	"gym-map/api"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.GetModels(cc, cc.ExerciseCrud)
}

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
