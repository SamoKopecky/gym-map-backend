package fetcher

import (
	"fmt"
	"gym-map/model"
	"gym-map/utils"
)

type KeycloakUser struct {
	Id        string  `json:"id"`
	Email     string  `json:"email"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
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

func (ku KeycloakUser) ToUserModel() model.User {
	return model.User{
		Id:    ku.Id,
		Email: ku.Email,
		Name:  ku.FullName(),
	}
}
