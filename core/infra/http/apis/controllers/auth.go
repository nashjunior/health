package controllers

import (
	"encoding/json"
	"health/core/clients/application/services"
	signin "health/core/clients/application/use-cases/sign-in"
	"net/http"
)

type AuthController struct {
}

func (authController *AuthController) RefreshToken(response http.ResponseWriter, request *http.Request) {
	body := map[string]any{}

	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := services.RefreshToken(
		services.Input{Token: body["token"].(string)},
	)

	if err != nil {
		http.Error(response, "Invalid Credentials", http.StatusForbidden)
	}

	response.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(token)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.Write(jsonResp)

	return
}

func (controller *AuthController) SignIn(response http.ResponseWriter, request *http.Request) {
	body := map[string]any{}

	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	signInUserUsecase := signin.New()
	token, err := signInUserUsecase.Execute(signin.Input{
		Email:    body["email"].(string),
		Password: body["password"].(string),
	})

	if err != nil {
		http.Error(response, "Invalid Credentials", http.StatusForbidden)
	}

	response.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(token)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.Write(jsonResp)

	return
}
