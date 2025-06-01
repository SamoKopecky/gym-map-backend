package api

import (
	"fmt"
	"gym-map/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

func BindParams[T any](c echo.Context) (T, error) {
	params := *new(T)
	if err := c.Bind(&params); err != nil {
		return params, err
	}
	return params, nil
}

type DbContext struct {
	echo.Context

	MachineCrud store.Machine
}

func (c DbContext) BadRequest(err error) error {
	errStr := fmt.Sprint(err)
	// TODO: log error too
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters", "reason": errStr})
}
