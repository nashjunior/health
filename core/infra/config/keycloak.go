package config

type KeycloakResponseClient struct {
	ID       string `json:"id"`
	ClientId string `json:"clientId"`
	Name     string `json:"name"`
}

type KeycloakResponseRealm struct {
	ID    string `json:"id"`
	Realm string `json:"realm"`
}
