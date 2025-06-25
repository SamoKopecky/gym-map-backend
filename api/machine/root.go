package machine

import (
	"gym-map/api"
	"gym-map/schema"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	machines, err := cc.MachineCrud.GetWithCount()
	if err != nil {
		return err
	}
	if machines == nil {
		machines = []schema.Machine{}
	}

	return cc.JSON(http.StatusOK, machines)
}

func Post(c echo.Context) error {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[machinePostRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	machine := params.ToNewModel()
	err = cc.MachineCrud.Insert(&machine)
	if err != nil {
		return err
	}

	exerciseWithCount := schema.Machine{
		Machine:       machine,
		ExerciseCount: 0,
	}

	return cc.JSON(http.StatusOK, exerciseWithCount)
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
