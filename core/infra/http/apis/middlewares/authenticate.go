package middlewares

import (
	"context"
	"fmt"
	"health/core/infra/config"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/gorilla/csrf"
)

type AuthenticateMiddleware func(http.Handler) http.Handler

func EnsureAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		rawAccessToken := r.Header.Get("Authorization")
		tokenAccess := csrf.Token(r)
		config.TokenMap.Store(tokenAccess, true)

		if rawAccessToken == "" {
			http.Redirect(w, r, config.Oauth2Config.AuthCodeURL(tokenAccess), http.StatusFound)
			return
		}

		parts := strings.Split(rawAccessToken, " ")
		if len(parts) != 2 {
			w.WriteHeader(400)
			return
		}

		token, err := config.Verifier.Verify(config.CtxAuth, parts[1])
		if err != nil {
			fmt.Println(err.Error())
			http.Redirect(w, r, config.Oauth2Config.AuthCodeURL(tokenAccess), http.StatusFound)
			return
		}

		var info *oidc.UserInfo

		err = token.Claims(&info)

		if err != nil {
			http.Redirect(w, r, config.Oauth2Config.AuthCodeURL(tokenAccess), http.StatusFound)
			return
		}

		ctx := context.WithValue(r.Context(), "userInfo", info)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
