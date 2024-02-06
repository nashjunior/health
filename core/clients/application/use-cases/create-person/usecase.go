package createperson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"health/core/clients/application/services"
	"health/core/infra/config"

	"io"
	"net/http"
	"time"
)

type Input struct {
	Name         string
	Cpf          *string
	Gender       *string
	Email        *string
	BirthdayDate *time.Time
}

type Output struct {
	Id           *string `json:"id"`
	Name         string  `json:"name"`
	Cpf          *string `json:"cpf"`
	Gender       *string `json:"gender"`
	Validate     *bool   `json:"validated"`
	birthdayDate *time.Time
	createdAt    time.Time
}

type Usecase struct {
}

func (usecase *Usecase) Execute(input Input) (*Output, error) {

	var output Output

	err := services.CreateKeycloakUser(services.KeycloakUser{
		Username:      input.Name,
		Email:         *input.Email,
		EmailVerified: true,
		FirstName:     input.Name,
		LastName:      input.Name,
		Enabled:       true,
		Attributes: map[string]*string{
			"cpf":    input.Cpf,
			"gender": input.Gender,
		},
	})

	if err != nil {
		return nil, err
	}
	now := time.Now()

	output.Name = input.Name
	output.Cpf = input.Cpf
	output.Gender = input.Gender
	output.birthdayDate = &now

	return &output, nil
}

func (usecase *Usecase) createKeycloakUser(input Input) error {
	url := config.KeycloakUrl + "/admin/realms/" + config.KeycloakCustomerRealm + "/users"

	client := &http.Client{}

	var groups *[]string
	var requiredActions *[]interface{}

	attributes := map[string]*string{
		"cpf":    input.Cpf,
		"gender": input.Gender,
	}

	body := map[string]interface{}{
		"username": input.Name,

		"email":         *input.Email,
		"emailVerified": true,

		"attributes": attributes,

		"firstName": input.Name,
		"lastName":  input.Name,

		"groups":          groups,
		"requiredActions": requiredActions,

		"enabled": true,
	}

	jsonValue, err := json.Marshal(body)

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

	if resp.StatusCode != 201 {
		return fmt.Errorf(bodyResponse["error"].(string))
	}

	return nil

}

func New() Usecase {
	return Usecase{}
}
