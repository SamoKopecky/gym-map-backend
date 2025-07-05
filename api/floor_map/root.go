package floormap

import (
	"gym-map/api"
	"net/http"
	"os"
	"path/filepath"

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

	file := cc.FormValue("file")
	println(file)

	// Destination
	err := os.WriteFile(filepath.Join("./files/map", "map.svg"), []byte(file), 0644)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}
