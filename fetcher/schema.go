package fetcher

import (
	"fmt"
	"gym-map/model"
	"gym-map/utils"
	"path"
)

type KeycloakUser struct {
	Id        string  `json:"id"`
	Email     string  `json:"email"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type NewKeycloakUser struct {
	Email           string   `json:"email"`
	Username        string   `json:"username"`
	Enabled         bool     `json:"enabled"`
	EmailVerified   bool     `json:"emailVerified"`
	RequiredActions []string `json:"requiredActions"`
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
	return model.UserBase{
		Email:     ku.Email,
		Name:      ku.FullName(),
		FirstName: ku.FirstName,
		LastName:  ku.LastName,
	}
}

func (ku KeycloakUser) ToUser() model.User {
	return model.User{
		Id:       ku.Id,
		UserBase: ku.ToUserBase(),
	}
}

type KeycloakRole struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// URL path to the user
type UserLocation string

func (ul UserLocation) UserId() string {
	return path.Base(string(ul))
}

func (i IAM) GetUserLocation(userId string) UserLocation {
	return UserLocation(
		fmt.Sprintf("%s/%s", i.userUrl(), userId))
}
