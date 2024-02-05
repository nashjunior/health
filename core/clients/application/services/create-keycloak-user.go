package services

import (
	"bytes"
	"encoding/json"
	"health/core/infra/config"
	"io"
	"net/http"
)

type KeycloakUser struct {
	Username        string             `json:"username"`
	Email           string             `json:"email"`
	FirstName       string             `json:"firstName"`
	LastName        string             `json:"lastName"`
	RequiredActions *[]string          `json:"requiredActions"`
	EmailVerified   bool               `json:"emailVerified"`
	Groups          *[]string          `json:"groups"`
	Enabled         bool               `json:"enabled"`
	Attributes      map[string]*string `json:"attributes"`
}

func CreateKeycloakUser(input KeycloakUser) error {
	url := config.KeycloakUrl + "/admin/realms/" + config.KeycloakCustomerRealm + "/users"

	client := &http.Client{}

	jsonValue, err := json.Marshal(input)

	if err != nil {
		return nil
	}

	requestNewUser, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	requestNewUser.Header.Set("Content-Type", "application/json")
	requestNewUser.Header.Set("Authorization", "Bearer "+*config.AdminAccessToken)

	resp, err := client.Do(requestNewUser)

	if err != nil {

		return err
	}

	bodyResponseReader, err := io.ReadAll(resp.Body)

	if err != nil {

		return err
	}

	var bodyResponse map[string]interface{}

	json.Unmarshal(bodyResponseReader, &bodyResponse)

	if resp.StatusCode > 204 {
		return http.ErrAbortHandler
	}

	return nil
}
