package fetcher

import (
	"bytes"
	"context"
	"fmt"
	"gym-map/config"
	"io"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const MASTER_TOKEN_ENDPOINT = "realms/master/protocol/openid-connect/token"

type IAM struct {
	AppConfig  *config.Config
	AuthConfig clientcredentials.Config
}

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

func (i IAM) GetUsers() ([]KeycloakUser, error) {
	resp, err := i.authedRequest(http.MethodGet, i.userUrl(), nil)

	if err != nil {
		return nil, err
	}

	return responseData[[]KeycloakUser](resp)
}
