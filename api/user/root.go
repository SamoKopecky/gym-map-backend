package user

import (
	"gym-map/api"
	"gym-map/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	users, err := cc.IAMFetcher.GetUsers()
	if err != nil {
		return cc.BadRequest(err)
	}

	userModels := make([]model.User, len(users))
	for i := range users {
		userModels[i] = users[i].ToUserModel()
	}

	return cc.JSON(http.StatusOK, userModels)
}

func Post(c echo.Context) (err error) {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[userPostRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	userId, err := cc.UserService.RegisterUser(params.Email)
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusCreated, struct {
		UserId string `json:"user_id"`
	}{UserId: userId})
}

func Delete(c echo.Context) (err error) {
	cc := c.(*api.DbContext)

	id := cc.Param("id")

	err = cc.UserService.UnregisterUser(id)
	if err != nil {
		return
	}
	return cc.NoContent(http.StatusOK)
}
