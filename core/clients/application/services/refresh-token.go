package services

import (
	"encoding/json"
	"health/core/infra/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Input struct {
	Token string
}

type Output struct {
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

func RefreshToken(Input Input) (*Output, error) {
	client := &http.Client{}

	urlLogin := config.KeycloakUrl + "/realms/" + config.KeycloakCustomerRealm + "/protocol/openid-connect/token"

	data := url.Values{}
	data.Add("refresh_token", Input.Token)
	data.Set("client_id", config.KeycloakCustomerClientID)
	data.Set("client_secret", config.KeycloaktCustomerClientSecret)
	data.Set("scope", "openid")
	data.Set("grant_type", "refresh_token")

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

	token := dados["id_token"]
	refreshToken := dados["refresh_token"]

	return &Output{IdToken: token, RefreshToken: refreshToken}, nil
}
