package middleware

import (
	"errors"
	"net/http"

	"github.com/matheus-alpe/go-coin-api/api"
	"github.com/matheus-alpe/go-coin-api/internal/tools"
)

var ErrUnauthorized = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	token := r.Header.Get("Authorization")

	if username == "" || token == "" {
	    api.RequestErrorHandler(w, ErrUnauthorized)
	    return
	}

	var database *tools.DatabaseInterface
	var err error
	database, err = tools.NewDatabase()
	if err != nil {
	    api.InternalErrorHandler(w, err)
	    return
	}

	loginDetails := (*database).GetUserLoginDetails(username)

	if loginDetails == nil || token != (*loginDetails).AuthToken {
	    api.RequestErrorHandler(w, ErrUnauthorized)
	    return
	}

	next.ServeHTTP(w, r)
    })
}
