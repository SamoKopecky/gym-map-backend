package app

import (
	"errors"
	"fmt"
	"gym-map/api"
	"gym-map/api/category"
	"gym-map/api/exercise"
	floormap "gym-map/api/floor_map"
	"gym-map/api/instruction"
	"gym-map/api/machine"
	"gym-map/api/media"
	"gym-map/api/user"
	"gym-map/config"
	"gym-map/crud"
	"gym-map/fetcher"
	fileio "gym-map/file_io"
	"gym-map/schema"
	"gym-map/service"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-contrib/echoprometheus"
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
			mediaCrud := crud.NewMedia(db)
			iamFetcher := fetcher.IAM{
				AppConfig:  cfg,
				AuthConfig: fetcher.CreateAuthConfig(cfg),
			}
			categoryCrud := crud.NewCategory(db)
			propertyCrud := crud.NewProperty(db)

			cc := &api.DbContext{Context: c,
				Config:          *cfg,
				MachineCrud:     crud.NewMachine(db),
				ExerciseCrud:    crud.NewExercise(db),
				InstructionCrud: instructionCrud,
				CategoryCrud:    categoryCrud,
				PropertyCrud:    propertyCrud,
				MediaCrud:       mediaCrud,
				IAMFetcher:      iamFetcher,
				FloorMapCrud:    fileio.FloorMap{Config: *cfg},
				InstructionService: service.Instruction{
					IAM:             iamFetcher,
					InstructionCrud: instructionCrud,
				},
				MediaService: service.Media{
					MediaCrud: mediaCrud,
				},
				UserService: service.User{
					IAM: iamFetcher,
				},
				CategoryService: service.Category{
					CategoryCrud: categoryCrud,
					PropertyCrud: propertyCrud,
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
	e.Use(echoprometheus.NewMiddleware("gym-map"))
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

	jwtUsersTrainerOnly := jwtUsers.Group("")
	jwtUsersTrainerOnly.Use(trainerOnlyMiddleware)
	jwtUsersTrainerOnly.GET("/:id", user.GetUser)
	jwtUsersTrainerOnly.PATCH("/profile", user.PatchProfile)

	jwtUsersAdminOnly := jwtUsers.Group("")
	jwtUsersAdminOnly.Use(adminOnlyMiddleware)
	jwtUsersAdminOnly.GET("", user.Get)
	jwtUsersAdminOnly.POST("", user.Post)
	jwtUsersAdminOnly.DELETE("/:id", user.Delete)

	mediaGroup := e.Group("/media")
	mediaGroup.GET("/:id", media.GetMedia)
	mediaGroup.GET("/metadata", media.GetMetadataMany)

	jwtMediaGroup := mediaGroup.Group("")
	jwtMediaGroup.Use(jwtMiddleware(appConfig))
	jwtMediaGroup.Use(claimContextMiddleware)
	jwtMediaGroup.Use(trainerOnlyMiddleware)
	jwtMediaGroup.DELETE("/:id", media.DeleteMedia)

	floorMap := e.Group("/map")
	floorMap.GET("", floormap.Get)
	floorMapJwt := floorMap.Group("")
	floorMapJwt.Use(jwtMiddleware(appConfig))
	floorMapJwt.Use(claimContextMiddleware)
	floorMapJwt.Use(adminOnlyMiddleware)
	floorMapJwt.PUT("", floormap.Put)

	categories := e.Group("/categories")
	categories.Use(jwtMiddleware(appConfig))
	categories.Use(claimContextMiddleware)
	categories.Use(adminOnlyMiddleware)
	categories.GET("", category.GetCategories)
	categories.POST("", category.Post)
	categories.PATCH("/:id", category.Patch)
	categories.DELETE("/:id", category.Delete)

	if err := e.Start(":2001"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func pong(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
