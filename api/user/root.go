package user

import (
	"gym-map/api"
	"gym-map/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	cc := c.(*api.DbContext)

	id := cc.Param("id")
	if id != cc.Claims.Subject {
		return cc.NoContent(http.StatusForbidden)
	}

	user, err := cc.IAMFetcher.GetUsersById(id)
	if err != nil {
		return cc.BadRequest(err)
	}

	return cc.JSON(http.StatusOK, user.ToUser())
}

func Get(c echo.Context) error {
	cc := c.(*api.DbContext)

	users, err := cc.UserService.GetUsers()
	if err != nil {
		return cc.BadRequest(err)
	}

	userModels := make([]model.User, len(users))
	for i := range users {
		userModels[i] = users[i].ToUser()
	}

	return cc.JSON(http.StatusOK, userModels)
}

func Post(c echo.Context) (err error) {
	cc := c.(*api.DbContext)

	params, err := api.BindParams[userPostRequest](cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	_, err = cc.UserService.RegisterUser(params.Email)
	if err != nil {
		return err
	}

	return cc.NoContent(http.StatusNoContent)
}

func Delete(c echo.Context) (err error) {
	cc := c.(*api.DbContext)

	id := cc.Param("id")

	err = cc.UserService.UnregisterUser(id)
	if err != nil {
		return
	}
	return cc.NoContent(http.StatusNoContent)
}

func PatchProfile(c echo.Context) error {
	cc := c.(*api.DbContext)
	userId := cc.Claims.Subject

	newMedia, err := api.CreateFileFromRequest(cc)
	if err != nil {
		return cc.BadRequest(err)
	}

	err = cc.UserService.UpdateAvatarId(userId, strconv.Itoa(newMedia.Id))
	if err != nil {
		return err
	}

	return cc.JSON(http.StatusOK, userPatchResponse{
		AvatarId: strconv.Itoa(newMedia.Id),
	})
}
