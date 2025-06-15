package app

import (
	"fmt"
	"gym-map/api"
	"gym-map/api/exercise"
	"gym-map/api/machine"
	"gym-map/config"
	"gym-map/crud"
	"gym-map/schema"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
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

func jwtMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	keyFunc := func(token *jwt.Token) (any, error) {
		return getKey(cfg, token)
	}
	return echojwt.WithConfig(echojwt.Config{
		KeyFunc:       keyFunc,
		SigningMethod: "RS256",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(schema.JwtClaims)
		},
	})
}

func claimContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*api.DbContext)
		cc.Claims = cc.Get("user").(*jwt.Token).Claims.(*schema.JwtClaims)
		return next(c)
	}
}

func RunApi(db *bun.DB, appConfig *config.Config) {
	e := echo.New()
	e.HTTPErrorHandler = logError
	e.Use(contextMiddleware(db, appConfig))
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/-/ping", pong)

	machines := e.Group("/machines")
	machines.GET("", machine.Get)

	jwtMachines := machines.Group("")
	jwtMachines.Use(jwtMiddleware(appConfig))
	jwtMachines.Use(claimContextMiddleware)
	jwtMachines.POST("", machine.Post)
	jwtMachines.PATCH("/:id", machine.Patch)
	jwtMachines.DELETE("/:id", machine.Delete)

	positions := machines.Group("/:id/positions")
	positions.Use(jwtMiddleware(appConfig))
	positions.Use(claimContextMiddleware)
	positions.PATCH("", machine.PatchPositions)

	exercises := e.Group("/exercises")
	exercises.GET("", exercise.Get)

	jwtExercises := machines.Group("")
	jwtExercises.Use(jwtMiddleware(appConfig))
	jwtExercises.Use(claimContextMiddleware)
	jwtExercises.POST("", exercise.Post)
	jwtExercises.PATCH("/:id", exercise.Patch)
	jwtExercises.DELETE("/:id", exercise.Delete)

	e.Logger.Fatal(e.Start(":2001"))
}

func pong(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
