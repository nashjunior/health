package api

import (
	"net/http"
)

type Route struct {
	URI                   string
	Method                string
	Callback              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}
