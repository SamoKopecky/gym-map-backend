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
