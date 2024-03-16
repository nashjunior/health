package main

import (
	"encoding/json"
	"fmt"
	"health/core/clients/application/services"
	routers "health/core/clients/infra/http/apis/routes"
	"health/core/infra/apis/middlewares"
	"health/core/infra/config"
	"health/core/infra/db"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"github.com/gorilla/csrf"
	"golang.org/x/oauth2"
)

func main() {

	config.Load()

	db.StartConnections()

	router := routers.GenerateRouter()
	services.GetAccessToken()

	accessToken := config.AdminAccessToken

	if accessToken == nil {
		log.Fatalf("Access token not found")
	}

	router.HandleFunc("/", middlewares.EnsureAuthenticated(func(w http.ResponseWriter, r *http.Request) {
		newToken := csrf.Token(r)
		config.TokenMap.Store(newToken, true)

		w.Header().Set("X-CSRF-Token", newToken)

		w.Write([]byte("could get a token"))
	}))

	router.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {

		if _, ok := config.TokenMap.Load(r.URL.Query().Get("state")); !ok {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Oauth2Config.Exchange(config.CtxAuth, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)

		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}
		idToken, err := config.Verifier.Verify(config.CtxAuth, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		userInfo, err := config.Provider.UserInfo(config.CtxAuth, oauth2.StaticTokenSource(oauth2Token))

		if err != nil {
			http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
			return
		}

		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
			UserInfo      *oidc.UserInfo
			Token         string
		}{oauth2Token, new(json.RawMessage), userInfo, rawIDToken}
		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	http.Handle("/", config.CsrfMiddleware(router))

	fmt.Println("iniciando servidor")
	log.Fatal(http.ListenAndServe("0.0.0.0:15000", nil))
}
