package services

import (
	"encoding/json"
	"health/core/infra/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type OutputExchangeToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func ExchanteToken(input string) (*OutputExchangeToken, error) {
	client := &http.Client{}

	urlLogin := config.KeycloakUrl + "/realms/" + config.KeycloakCustomerRealm + "/protocol/openid-connect/token"

	data := url.Values{}
	data.Set("grant_type", "urn:ietf:params:oauth:grant-type:token-exchange")

	data.Set("client_id", config.KeycloakCustomerClientID)
	data.Set("client_secret", config.KeycloaktCustomerClientSecret)

	data.Set("requested_token_type", "urn:ietf:params:oauth:token-type:refresh_token")

	data.Set("subject_token", input)
	data.Set("subject_token_type", "urn:ietf:params:oauth:token-type:access_token")

	data.Set("scope", "openid")

	req, err := http.NewRequest("POST", urlLogin, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Realiza a requisição
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	var dados map[string]string
	json.Unmarshal(body, &dados)
	if err != nil {
		return nil, err
	}

	token := dados["access_token"]
	refreshToken := dados["refresh_token"]

	return &OutputExchangeToken{AccessToken: token, RefreshToken: refreshToken}, nil
}
