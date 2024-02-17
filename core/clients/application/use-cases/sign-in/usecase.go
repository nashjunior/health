package signin

import (
	"encoding/json"
	"health/core/infra/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ResponseAuthClient struct {
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

type Input struct {
	Email    string
	Password string
}

type Output struct {
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

type Usecase struct {
}

func (usecase *Usecase) Execute(input Input) (*Output, error) {

	token := usecase.makeRequest(input)

	if token == nil {
		return nil, http.ErrAbortHandler
	}

	return &Output{
		IdToken:      token.IdToken,
		RefreshToken: token.RefreshToken,
	}, nil

}

func (usecase *Usecase) makeRequest(input Input) *ResponseAuthClient {
	client := &http.Client{}
	uri := config.KeycloakUrl + "/realms/" + config.KeycloakCustomerRealm + "/protocol/openid-connect/token"

	data := url.Values{}
	data.Set("client_id", config.KeycloakCustomerClientID)
	data.Set("client_secret", config.KeycloaktCustomerClientSecret)
	data.Set("username", input.Email)
	data.Set("password", input.Password)
	data.Set("scope", "openid")
	data.Set("grant_type", "password")

	requestAuthUser, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))

	if err != nil {
		return nil
	}
	requestAuthUser.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	q := requestAuthUser.URL.Query()
	q.Add("response_type", "code")
	requestAuthUser.URL.RawQuery = q.Encode()

	response, err := client.Do(requestAuthUser)

	if err != nil {
		return nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil
	}

	var contentResponse ResponseAuthClient

	json.Unmarshal(body, &contentResponse)

	return &contentResponse

}

func New() Usecase {
	return Usecase{}
}
