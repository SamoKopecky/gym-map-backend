package category

import (
	"gym-map/api"
	"gym-map/schema"
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

func GetCategories(c echo.Context) error {
	cc := c.(*api.DbContext)

	categories, err := cc.CategoryService.GetCategories()
	if err != nil {
		return err
	}

	if len(categories) == 0 {
		categories = []schema.Category{}
	}

	return cc.JSON(http.StatusOK, categories)
}
