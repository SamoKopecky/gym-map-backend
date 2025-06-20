package fetcher

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gym-map/config"
	"io"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const TRAINER_ROLE = "trainer"

var ErrUserAlreadyExists = errors.New("iam: user already exists")
var ErrUserNotCreated = errors.New("iam: user not created due to invalid status code")
var ErrUserActionTriggerFailed = errors.New("iam: user trigger failed because of unknown status code")

var newRequiredActions = []string{"UPDATE_PROFILE", "UPDATE_PASSWORD", "VERIFY_EMAIL"}

func CreateAuthConfig(appConfig *config.Config) clientcredentials.Config {
	return clientcredentials.Config{
		ClientID:     appConfig.KeycloakAdminClientId,
		ClientSecret: appConfig.KeycloakAdminClientSecret,
		TokenURL:     fmt.Sprintf("%s/%s", appConfig.KeycloakBaseUrl, MASTER_TOKEN_ENDPOINT),
	}
}

func (i IAM) baseUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", i.AppConfig.KeycloakBaseUrl, endpoint)
}

func (i IAM) userUrl() string {
	return i.baseUrl(fmt.Sprintf("admin/realms/%s/users", i.AppConfig.KeycloakRealm))
}

func (i IAM) roleUrl(role string) string {
	return i.baseUrl(fmt.Sprintf("admin/realms/%s/roles/%s", i.AppConfig.KeycloakRealm, role))
}

func (i IAM) authedRequest(method, url string, body *bytes.Buffer) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		reqBody = body
	} else {
		reqBody = nil
	}
	request, err := http.NewRequest(method, url, reqBody)

	if err != nil {
		return nil, err
	}
	client := oauth2.NewClient(context.Background(), i.AuthConfig.TokenSource(context.Background()))
	return client.Do(request)
}

func (i IAM) editUserRoles(method string, userLocation UserLocation, kcRole KeycloakRole) error {
	buf := createParamsBuf([]KeycloakRole{kcRole})
	resp, err := i.authedRequest(method, fmt.Sprintf("%s/role-mappings/realm", userLocation), &buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return err
	}

	return nil
}

func responseData[T any](response *http.Response) (data T, err error) {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &data)
	return
}

func createParamsBuf(data any) (buf bytes.Buffer) {
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
	return
}
