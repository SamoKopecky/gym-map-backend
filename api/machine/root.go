package machine

import (
	"gym-map/api"
	"gym-map/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	machines, err := cc.MachineCrud.Get()
	if err != nil {
		return err
	}

	if machines == nil {
		machines = []model.Machine{}
	}

	return cc.JSON(http.StatusOK, machines)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PostModel[machinePostRequest](cc, cc.MachineCrud)
}

func Patch(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PatchModel[machinePatchRequest](cc, cc.MachineCrud)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.MachineCrud)
}
