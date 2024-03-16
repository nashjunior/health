package controllers

import (
	"encoding/json"
	"health/core/clients/application/services"
	signin "health/core/clients/application/use-cases/sign-in"
	"net/http"

	"github.com/gorilla/csrf"
)

type AuthController struct {
}

func (authController *AuthController) RefreshToken(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
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

	response.Header().Set("X-CSRF-Token", csrf.Token(request))

	jsonResp, err := json.Marshal(token)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.Write(jsonResp)

	return
}

func (controller *AuthController) SignIn(response http.ResponseWriter, request *http.Request) {
	body := map[string]any{}
	response.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	csrfToken := csrf.Token(request)

	if csrfToken == "" {
		http.Error(response, "Invalid Csrf token", http.StatusForbidden)
		return
	}

	signInUserUsecase := signin.New()
	token, err := signInUserUsecase.Execute(signin.Input{
		Email:    body["email"].(string),
		Password: body["password"].(string),
	})

	if err != nil {
		http.Error(response, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("X-CSRF-Token", csrf.Token(request))

	if body["callbackUrl"] != nil {
		http.Redirect(response, request, body["callbackUrl"].(string), 301)
		return
	}

	jsonResp, err := json.Marshal(token)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.Write(jsonResp)

	return
}

func (controller *AuthController) GetCode(response http.ResponseWriter, request *http.Request) {

	// if _, ok := config.TokenMap.Load(token); !ok {
	// 	http.Error(w, "state did not match", http.StatusBadRequest)
	// 	return
	// }

	newToken := csrf.Token(request)

	response.Header().Set("X-CSRF-Token", newToken)
	response.Write([]byte("could get a token"))
	return
}

func (controller *AuthController) ExchangeAccessToken(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var requestData map[string]string
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(response, "Erro ao decodificar os dados da solicitação", http.StatusBadRequest)
		return
	}

	authorizationCode := requestData["token"]

	response.Header().Set("X-CSRF-Token", csrf.Token(request))

	accessToken, err := services.ExchanteToken(authorizationCode)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(accessToken)

	response.Write(jsonResp)
	return
}
