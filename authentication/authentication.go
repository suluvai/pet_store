package authentication

import (
	"errors"
	"net/http"
	"pet_store_rest_api/responses"
)

type Secret struct {
	// For this assignment purpose using a simple string as key.
	MySigningKey string
}

// Middleware function, which will be called for each request
func (auth *Secret) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			if r.Header["Token"][0] == auth.MySigningKey {
				next.ServeHTTP(w, r)
				return
			}
		} 
		// Write an error and stop the handler chain
		responses.ERROR(w, http.StatusForbidden, errors.New("No Token"))
	})
}
