package middleware

import (
	"errors"
	"net/http"

	"go-simple-api/api"
	"go-simple-api/internal/tools"

	log "github.com/sirupsen/logrus"
)

var UnAuthorizeError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizeError)
			api.RequestErrorHandler(w, UnAuthorizeError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizeError)
			api.RequestErrorHandler(w, UnAuthorizeError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
