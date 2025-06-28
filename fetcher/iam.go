package fetcher

import (
	"fmt"
	"gym-map/config"
	"gym-map/utils"
	"net/http"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
)

const MASTER_TOKEN_ENDPOINT = "realms/master/protocol/openid-connect/token"

type IAM struct {
	AppConfig  *config.Config
	AuthConfig clientcredentials.Config
}

func (i IAM) GetUsersByRole(role string) ([]KeycloakUser, error) {
	resp, err := i.authedRequest(http.MethodGet, fmt.Sprintf("%s/users", i.roleUrl(role)), nil)

	if err != nil {
		return nil, err
	}

	return responseData[[]KeycloakUser](resp)
}

func (i IAM) GetUsersById(id string) (KeycloakUser, error) {
	resp, err := i.authedRequest(http.MethodGet, fmt.Sprintf("%s/%s", i.userUrl(), id), nil)

	if err != nil {
		return KeycloakUser{}, err
	}

	return responseData[KeycloakUser](resp)
}

func (i IAM) GetUsers() ([]KeycloakUser, error) {
	resp, err := i.authedRequest(http.MethodGet, i.userUrl(), nil)

	if err != nil {
		return nil, err
	}

	return responseData[[]KeycloakUser](resp)
}

func (i IAM) CreateUser(email string) (userLocation UserLocation, err error) {
	username := strings.Split(email, "@")[0]
	newUser := NewKeycloakUser{
		Email:           email,
		Username:        username,
		Enabled:         true,
		EmailVerified:   false,
		RequiredActions: newRequiredActions,
	}
	buf := createParamsBuf(newUser)
	resp, err := i.authedRequest(http.MethodPost, i.userUrl(), &buf)
	if err != nil {
		return
	}
	if resp.StatusCode == http.StatusConflict {
		return "", ErrUserAlreadyExists
	}

	if resp.StatusCode != http.StatusCreated {
		return "", ErrUserNotCreated
	}

	userLocation = UserLocation(resp.Header.Get("Location"))
	return
}

func (i IAM) GetUserLocationByEmail(email string) (UserLocation, error) {
	baseUrl := fmt.Sprintf("%s", i.userUrl())
	queryParams := map[string]string{"email": email, "exact": "true"}
	resp, err := i.authedRequest(
		http.MethodGet,
		utils.AddQueryParam(baseUrl, queryParams),
		nil)

	if err != nil {
		return "", err
	}

	users, err := responseData[[]KeycloakUser](resp)
	if err != nil {
		return "", err
	}

	// TODO: Properly check len
	user := users[0]
	userLocation := i.GetUserLocation(user.Id)
	return userLocation, nil
}

func (i IAM) InvokeUserUpdate(userLocation UserLocation) error {
	buf := createParamsBuf(newRequiredActions)

	resp, err := i.authedRequest(http.MethodPut, fmt.Sprintf("%s/execute-actions-email", userLocation), &buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return ErrUserActionTriggerFailed
	}
	return nil
}

func (i IAM) GetRole(roleName string) (KeycloakRole, error) {
	resp, err := i.authedRequest(http.MethodGet, i.roleUrl(roleName), nil)
	if err != nil {
		return KeycloakRole{}, err
	}

	return responseData[KeycloakRole](resp)
}

func (i IAM) AddUserRoles(userLocation UserLocation, kcRole KeycloakRole) error {
	return i.editUserRoles(http.MethodPost, userLocation, kcRole)
}

func (i IAM) RemoveUserRoles(userLocation UserLocation, kcRole KeycloakRole) error {
	return i.editUserRoles(http.MethodDelete, userLocation, kcRole)
}

func (i IAM) UpdateUser(user KeycloakUser) error {
	buf := createParamsBuf(user)

	resp, err := i.authedRequest(http.MethodPut, i.userIdUrl(user.Id), &buf)
	if err != nil {
		return err
	}

	// TODO: return status code
	if resp.StatusCode != http.StatusNoContent {
		return ErrUserNotUpdated
	}

	return nil
}
