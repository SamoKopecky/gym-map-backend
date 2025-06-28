package fetcher

import (
	"fmt"
	"gym-map/model"
	"gym-map/utils"
	"path"
)

// URL path to the user
type UserLocation string

type KeycloakUser struct {
	Id         string             `json:"id"`
	Email      string             `json:"email"`
	FirstName  *string            `json:"firstName"`
	LastName   *string            `json:"lastName"`
	Attributes KeycloakAttributes `json:"attributes"`
}

type NewKeycloakUser struct {
	Email           string   `json:"email"`
	Username        string   `json:"username"`
	Enabled         bool     `json:"enabled"`
	EmailVerified   bool     `json:"emailVerified"`
	RequiredActions []string `json:"requiredActions"`
}

type KeycloakRole struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type KeycloakAttributes struct {
	AvatarId []string `json:"avatarId"`
}

func (ku KeycloakUser) FullName() *string {
	if ku.FirstName != nil && ku.LastName != nil {
		name := fmt.Sprintf("%s %s",
			utils.UpperFirstChar(*ku.FirstName),
			utils.UpperFirstChar(*ku.LastName))
		return &name
	}
	return nil
}

func (ku KeycloakUser) ToUserBase() model.UserBase {
	user := model.UserBase{
		Name:      ku.FullName(),
		FirstName: ku.FirstName,
		LastName:  ku.LastName,
		AvatarId:  nil,
	}
	if len(ku.Attributes.AvatarId) > 0 {
		user.AvatarId = &ku.Attributes.AvatarId[0]
	}
	return user
}

func (ku KeycloakUser) ToUser() model.User {
	return model.User{
		Id:       ku.Id,
		Email:    ku.Email,
		UserBase: ku.ToUserBase(),
	}
}

func (ul UserLocation) UserId() string {
	return path.Base(string(ul))
}
