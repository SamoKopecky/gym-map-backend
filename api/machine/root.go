package machine

import (
	"gym-map/api"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.GetModels(cc, cc.MachineCrud)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PostModel[machinePostRequest](cc, cc.MachineCrud)
}

func Patch(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.PatchModel[machinePatchRequest](cc, cc.MachineCrud)
}

func PatchPositions(c echo.Context) error {
	cc := c.(*api.DbContext)

	id, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return cc.BadRequest(err)
	}

	params, err := api.BindParams[machinePositionsPatchRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	model := params.ToExistingModel(id)
	err = cc.MachineCrud.UpdatePosition(&model)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusOK)
}

func Delete(c echo.Context) error {
	cc := c.(*api.DbContext)
	return api.DeleteModel(cc, cc.MachineCrud)
}
