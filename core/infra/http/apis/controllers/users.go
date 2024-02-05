package controllers

import (
	"bytes"
	"encoding/json"
	"health/core/clients/application/services"
	"health/core/infra/config"

	"io"
	"net/http"
)

type UsersController struct {
}

func (controller *UsersController) Create(response http.ResponseWriter, request *http.Request) {

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
