package controllers

import (
	"bytes"
	"encoding/json"
	"health/core/clients/application/services"
	createenterprise "health/core/clients/application/use-cases/create-enterprise"
	createperson "health/core/clients/application/use-cases/create-person"
	signup "health/core/clients/application/use-cases/sign-up"
	inmemory "health/core/clients/infra/db/in-memory"
	"health/core/infra/config"
	"time"

	"io"
	"net/http"

	"github.com/gorilla/csrf"
)

type UsersController struct {
}

func (controller *UsersController) Create(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	body := map[string]any{}

	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	var user any

	var email *string

	if body["email"] != nil {
		emailCustom := body["email"].(string)
		email = &emailCustom
	}

	var birthdayDate *time.Time

	if body["birthdayDate"] != nil {
		birthdayDateCustom := body["birthdayDate"].(time.Time)
		birthdayDate = &birthdayDateCustom
	}

	if config.ValidationCode {
		usecase := signup.NewUsecase(inmemory.NewUsersInMemoryRepository())
		user, err = usecase.Execute(signup.Input{Name: body["name"].(string)})
	} else if body["cnpj"] != nil {
		usecase := createenterprise.New()
		user, err = usecase.Execute(createenterprise.Input{
			Name:         body["name"].(string),
			Cnpj:         body["cnpj"].(string),
			Email:        email,
			SocialReason: body["social_reason"].(string),
			BirthdayDate: birthdayDate,
		})

		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(err.Error()))
			return
		}
	} else {

		var cpf *string

		if body["cpf"] != nil {
			cpfCustom := body["cpf"].(string)
			cpf = &cpfCustom
		}

		var gender *string

		if body["gender"] != nil {
			genderCustom := body["gender"].(string)
			gender = &genderCustom
		}

		usecase := createperson.New()
		user, err = usecase.Execute(createperson.Input{
			Name:         body["name"].(string),
			Cpf:          cpf,
			Email:        email,
			Gender:       gender,
			BirthdayDate: birthdayDate,
		})

		if err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}
	}

	response.Header().Set("X-CSRF-Token", csrf.Token(request))

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	encoded, err := json.Marshal(user)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusCreated)
	response.Write(encoded)
	return

}

func (controller *UsersController) Validate(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	if config.AdminAccessToken == nil {
		services.GetAccessToken()
		controller.Create(response, request)
	}

	body := map[string]any{}

	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	jsonValue, err := json.Marshal(body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("X-CSRF-Token", csrf.Token(request))

	url := config.KeycloakUrl + "/admin/realms/" + config.KeycloakCustomerRealm + "/users"

	client := &http.Client{}

	//"firstName":"Sergey","lastName":"Kargopolov", "email":"test@test.com", "enabled":"true", "username":"app-user"

	requestNewUser, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	requestNewUser.Header.Set("Content-Type", "application/json")
	requestNewUser.Header.Set("Authorization", "Bearer "+*config.AdminAccessToken)

	resp, err := client.Do(requestNewUser)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	bodyResponseReader, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	var bodyResponse any

	json.Unmarshal(bodyResponseReader, &bodyResponse)

	if resp.StatusCode > 204 {
		response.WriteHeader(http.StatusBadRequest)
		response.Write(bodyResponseReader)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Write(bodyResponseReader)
	return

}
