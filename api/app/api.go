package app

import (
	"fmt"
	"gym-map/api"
	"gym-map/api/exercise"
	"gym-map/api/machine"
	"gym-map/config"
	"gym-map/crud"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

func logError(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	errMsg := "Internal server error"
	httpError, ok := err.(*echo.HTTPError)
	if ok {
		code = httpError.Code
		errMsg = fmt.Sprint(httpError.Message)
	}

	if err := c.JSON(code, map[string]string{"message": errMsg}); err != nil {
		c.Logger().Error(err)
	}
}

func contextMiddleware(db *bun.DB, cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &api.DbContext{Context: c,
				MachineCrud:  crud.NewMachine(db),
				ExerciseCrud: crud.NewExercise(db),
			}

			return next(cc)
		}
	}
}

func RunApi(db *bun.DB, appConfig *config.Config) {
	e := echo.New()
	e.HTTPErrorHandler = logError
	e.Use(contextMiddleware(db, appConfig))
	e.Use(middleware.CORS())
	e.GET("/-/ping", pong)

	machines := e.Group("/machines")
	machines.GET("", machine.Get)
	machines.POST("", machine.Post)
	machines.PATCH("/:id", machine.Patch)
	machines.DELETE("/:id", machine.Delete)

	exercises := e.Group("/exercises")
	exercises.GET("", exercise.Get)
	exercises.POST("", exercise.Post)
	exercises.PATCH("/:id", exercise.Patch)
	exercises.DELETE("/:id", exercise.Delete)

	e.Logger.Fatal(e.Start(":2001"))
}

func pong(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
