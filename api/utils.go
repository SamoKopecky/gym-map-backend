package api

import (
	"fmt"
	"gym-map/config"
	"gym-map/fetcher"
	"gym-map/schema"
	"gym-map/service"
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

	MachineCrud     store.Machine
	ExerciseCrud    store.Exercise
	InstructionCrud store.Instruction

	IAMFetcher fetcher.IAM

	InstructionService service.Instruction
	UserService        service.User

	Claims *schema.JwtClaims

	Config config.Config
}

func (c DbContext) BadRequest(err error) error {
	errStr := fmt.Sprint(err)
	// TODO: log error too
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid query parameters", "reason": errStr})
}

func DerefString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func DerefInt(ptr *int) int {
	if ptr == nil {
		return 0
	}
	return *ptr

}
