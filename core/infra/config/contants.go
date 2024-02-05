package config

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/coreos/go-oidc"
	"github.com/gorilla/csrf"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var (
	CtxAuth     = context.Background()
	KeycloakUrl string
	ConfigURL   string

	ValidationCode bool

	//MASTER REALM CONFIG
	KeycloakAdminRealm        string
	KeycloakAdminClientId     string
	KeycloakAdminClientSecret string
	AdminAccessToken          *string

	//CUSTOMER REALM CONFIG
	KeycloakCustomerRealm         string
	KeycloakCustomerClientID      string
	KeycloaktCustomerClientSecret string
	RedirectURL                   string
	CustomerAccessToken           *string

	CsrfMiddleware func(http.Handler) http.Handler
	TokenMap       = sync.Map{}
	Provider       *oidc.Provider
	OidcConfig     *oidc.Config

	Verifier     *oidc.IDTokenVerifier
	Oauth2Config oauth2.Config
)

func newProvider() *oidc.Provider {
	newProvider, err := oidc.NewProvider(CtxAuth, ConfigURL)
	if err != nil {
		panic(err)
	}

	return newProvider
}

func Load() {

	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "local"
	}

	fileEnv := ".env." + env

	err := godotenv.Load(fileEnv)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Reading " + fileEnv)

	hasValidationCode := os.Getenv("VALIDATION_CODE")

	if hasValidationCode == "true" {
		ValidationCode = true
	}

	KeycloakUrl = os.Getenv("KEYCLOAK_URL")

	//MASTER REALM
	KeycloakAdminRealm = os.Getenv("KEYCLOAK_ADMIN_REALM")
	KeycloakAdminClientId = os.Getenv("KEYCLOAK_ADMIN_CLIENT_ID")
	KeycloakAdminClientSecret = os.Getenv("KEYCLOAK_ADMIN_CLIENT_SECRET")

	//CURRENT REALM
	KeycloakCustomerRealm = os.Getenv("KEYCLOAK_REALM")
	ConfigURL = KeycloakUrl + "/realms/" + KeycloakCustomerRealm
	KeycloakCustomerClientID = os.Getenv("KEYCLOAK_CLIENT_ID")
	KeycloaktCustomerClientSecret = os.Getenv("KEYCLOAK_CLIENT_SECRET")
	RedirectURL = os.Getenv("KEYCLOAK_REDIRECT_URL")

	//CRSF
	tokenCsrf := os.Getenv("CRSF_KEY")

	CsrfMiddleware = csrf.Protect([]byte(tokenCsrf), csrf.Secure(false))

	Provider = newProvider()
	OidcConfig = &oidc.Config{ClientID: KeycloakCustomerClientID}
	Verifier = Provider.Verifier(OidcConfig)

	Oauth2Config = oauth2.Config{
		ClientID:     KeycloakCustomerClientID,
		ClientSecret: KeycloaktCustomerClientSecret,
		RedirectURL:  RedirectURL,
		// Discovery returns the OAuth2 endpoints.
		Endpoint: Provider.Endpoint(),
		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email", "role"},
	}
}
