package app

import (
	"fmt"
	"gym-map/api"
	"gym-map/api/exercise"
	"gym-map/api/instruction"
	"gym-map/api/machine"
	"gym-map/api/user"
	"gym-map/config"
	"gym-map/crud"
	"gym-map/fetcher"
	"gym-map/schema"
	"gym-map/service"
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
			instructionCrud := crud.NewInstruction(db)
			iamFetcher := fetcher.IAM{
				AppConfig:  cfg,
				AuthConfig: fetcher.CreateAuthConfig(cfg),
			}

			cc := &api.DbContext{Context: c,
				Config:          *cfg,
				MachineCrud:     crud.NewMachine(db),
				ExerciseCrud:    crud.NewExercise(db),
				InstructionCrud: instructionCrud,
				IAMFetcher:      iamFetcher,
				InstructionService: service.Instruction{
					IAM:             iamFetcher,
					InstructionCrud: instructionCrud,
				},
				UserService: service.User{
					IAM: iamFetcher,
				},
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

func trainerOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*api.DbContext)
		if !cc.Claims.IsTrainer() && !cc.Claims.IsAdmin() {
			return cc.NoContent(http.StatusForbidden)
		}

		return next(c)
	}
}

func adminOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*api.DbContext)
		if !cc.Claims.IsAdmin() {
			return cc.NoContent(http.StatusForbidden)
		}

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
	jwtMachines.Use(adminOnlyMiddleware)
	jwtMachines.POST("", machine.Post)
	jwtMachines.PATCH("/:id", machine.Patch)
	jwtMachines.DELETE("/:id", machine.Delete)

	positions := jwtMachines.Group("/:id/positions")
	positions.PATCH("", machine.PatchPositions)

	exercises := e.Group("/exercises")
	exercises.GET("", exercise.Get)

	jwtExercises := exercises.Group("")
	jwtExercises.Use(jwtMiddleware(appConfig))
	jwtExercises.Use(claimContextMiddleware)
	jwtExercises.Use(adminOnlyMiddleware)
	jwtExercises.POST("", exercise.Post)
	jwtExercises.PATCH("/:id", exercise.Patch)
	jwtExercises.DELETE("/:id", exercise.Delete)

	instructions := e.Group("/instructions")
	instructions.GET("", instruction.Get)
	instructions.GET("/:id/media", instruction.GetMedia)

	jwtInstructions := instructions.Group("")
	jwtInstructions.Use(jwtMiddleware(appConfig))
	jwtInstructions.Use(claimContextMiddleware)
	jwtInstructions.Use(trainerOnlyMiddleware)
	jwtInstructions.POST("", instruction.Post)
	jwtInstructions.PATCH("/:id", instruction.Patch)
	jwtInstructions.DELETE("/:id", instruction.Delete)
	jwtInstructions.POST("/:id/media", instruction.PostMedia)

	jwtUsers := e.Group("/users")
	jwtUsers.Use(jwtMiddleware(appConfig))
	jwtUsers.Use(claimContextMiddleware)
	jwtUsers.Use(adminOnlyMiddleware)
	jwtUsers.GET("", user.Get)
	jwtUsers.POST("", user.Post)
	jwtUsers.DELETE("/:id", user.Delete)

	e.Logger.Fatal(e.Start(":2001"))
}

func pong(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
