package category

import (
	"gym-map/api"
	"gym-map/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PostModel[categoryPostRequest](cc, cc.CategoryCrud)
}

func Patch(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PatchModel[categoryPatchRequest](cc, cc.CategoryCrud)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.CategoryCrud)
}

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	categories, err := cc.CategoryCrud.GetCategoryProperties()
	if err != nil {
		return err
	}

	if len(categories) == 0 {
		categories = []model.Category{}
	}

	return cc.JSON(http.StatusOK, categories)
}
