package service

import (
	"errors"
	"gym-map/fetcher"
)

type User struct {
	IAM fetcher.IAM
}

func (u User) GetUsers() (users []fetcher.KeycloakUser, err error) {
	trainers, err := u.IAM.GetUsersByRole(fetcher.TRAINER_ROLE)
	if err != nil {
		return
	}

	users = append(users, trainers...)
	return
}

func (u User) RegisterUser(email string) (userId string, err error) {
	userLocation, err := u.IAM.CreateUser(email)
	if errors.Is(err, fetcher.ErrUserAlreadyExists) {
		// If use is already created, get user id by email
		userLocation, err = u.IAM.GetUserLocationByEmail(email)
		if err != nil {
			return
		}
	} else if err == nil {
		// If user is created, invoke user update to set password etc.
		err = u.IAM.InvokeUserUpdate(userLocation)
		if err != nil {
			return
		}
	} else {
		return
	}

	trainer_role, err := u.IAM.GetRole(fetcher.TRAINER_ROLE)
	if err != nil {
		return
	}

	err = u.IAM.AddUserRoles(userLocation, trainer_role)
	if err != nil {
		return
	}

	return userLocation.UserId(), nil
}

func (u User) UnregisterUser(userId string) error {
	userLocation := u.IAM.GetUserLocation(userId)
	trainerRole, err := u.IAM.GetRole(fetcher.TRAINER_ROLE)
	if err != nil {
		return err
	}

	err = u.IAM.RemoveUserRoles(userLocation, trainerRole)
	if err != nil {
		return err
	}
	return nil
}

func (u User) updateAttributes(userId string, attributes fetcher.KeycloakAttributes) error {
	user, err := u.IAM.GetUsersById(userId)
	if err != nil {
		return err
	}

	if len(attributes.AvatarId) > 0 {
		user.Attributes.AvatarId = attributes.AvatarId
	}

	err = u.IAM.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u User) UpdateAvatarId(userId, avatarId string) error {
	attributes := fetcher.KeycloakAttributes{AvatarId: []string{avatarId}}
	err := u.updateAttributes(userId, attributes)

	if err != nil {
		return err
	}

	return nil
}
