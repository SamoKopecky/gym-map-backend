package property

import (
	"gym-map/api"

	"github.com/labstack/echo/v4"
)

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PostModel[propertyPostRequest](cc, cc.PropertyCrud)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.PropertyCrud)
}
