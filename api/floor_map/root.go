package floormap

import (
	"gym-map/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	floorMap, err := cc.FloorMapCrud.GetMap()
	if err != nil {
		cc.BadRequest(err)
	}

	return cc.HTML(http.StatusOK, string(floorMap))
}

func Put(c echo.Context) error {
	cc := c.(*api.DbContext)

	floorMap, err := cc.FloorMapCrud.GetMap()
	if err != nil {
		cc.BadRequest(err)
	}

	return cc.HTML(http.StatusOK, string(floorMap))

}
