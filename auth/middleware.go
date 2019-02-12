package auth

import (
	"net/http"

	"github.com/berlingoqc/fxbe/utility"
)

// MiddlewareFile is the middleware for authentification on the files api
func MiddlewareFile(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := ValidUserCookie(r)
		if err != nil {
			utility.RespondWithError(w, http.StatusUnauthorized, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
