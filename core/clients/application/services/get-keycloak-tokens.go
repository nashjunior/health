package services

import (
	"encoding/json"
	"fmt"
	"health/core/infra/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func GetAccessToken() {
	fmt.Println("Retreiving access tokens for both customer realm and master realm")

	if config.AdminAccessToken == nil {
		token, err := makeRequest(config.KeycloakAdminRealm, config.KeycloakAdminClientId, config.KeycloakAdminClientSecret)
		if err != nil {
			panic(err)
		}

		if token != nil {
			config.AdminAccessToken = token
		}
	}

	if config.CustomerAccessToken != nil {
		token, err := makeRequest(config.KeycloakCustomerRealm, config.KeycloakCustomerClientID, config.KeycloaktCustomerClientSecret)
		if err != nil {
			panic(err)
		}

		if token != nil {
			config.AdminAccessToken = token
		}
	}
}

func makeRequest(realm string, clientId string, clientSecret string) (*string, error) {
	client := &http.Client{}

	urlLogin := config.KeycloakUrl + "/realms/" + realm + "/protocol/openid-connect/token"

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("scope", "openid")
	data.Set("grant_type", "client_credentials")

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
	token := dados["access_token"]

	return &token, nil

}

// func GetAccessToken() {

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	tokenCh := make(chan string, 2)
// 	errCh := make(chan error, 2)

// 	go makeRequest(config.KeycloakAdminRealm, config.KeycloakAdminClientId, config.KeycloakAdminClientSecret, &wg, tokenCh, errCh)
// 	go makeRequest(config.KeycloakCustomerRealm, config.KeycloakCustomerClientID, config.KeycloaktCustomerClientSecret, &wg, tokenCh, errCh)

// 	wg.Wait()

// 	close(tokenCh)
// 	close(errCh)

// 	for i := 0; i < 2; i++ {
// 		result := <-tokenCh

// 		fmt.Println(result)

// 		if i == 0 {
// 			config.AdminAccessToken = &result
// 		} else if i == 1 {
// 			config.CustomerAccessToken = &result
// 		}
// 	}

// 	// Close the channel once all goroutines are done

// }

// func makeRequest(realm string, clientId string, clientSecret string, wg *sync.WaitGroup, ch chan<- string, errChan chan<- error) {
// 	defer wg.Done()

// 	client := &http.Client{}

// 	urlLogin := config.KeycloakUrl + "/realms/" + realm + "/protocol/openid-connect/token"

// 	data := url.Values{}
// 	data.Set("client_id", clientId)
// 	data.Set("client_secret", clientSecret)
// 	data.Set("scope", "openid")
// 	data.Set("grant_type", "client_credentials")

// 	req, err := http.NewRequest("POST", urlLogin, strings.NewReader(data.Encode()))

// 	if err != nil {
// 		errChan <- err
// 	}

// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// Realiza a requisição
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		errChan <- err
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	var dados map[string]string
// 	json.Unmarshal(body, &dados)
// 	if err != nil {
// 		errChan <- err
// 	}

// 	fmt.Println(clientId)

// 	fmt.Println(dados["access_token"])

// 	ch <- dados["access_token"]

// }
