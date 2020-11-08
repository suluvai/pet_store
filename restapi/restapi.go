package restapi

import (
	"log"
	"net/http"
	"pet_store_rest_api/authentication"
	"pet_store_rest_api/controllers"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// SetupRouter sets up the rest api handlers.
func (app *App) SetupRouter(petController *controllers.PetController, auth *authentication.Secret) {
	// Find pets by status
	app.Router.Methods("GET").Path("/pet/findByStatus").HandlerFunc(petController.ReadPetsInfoByStatus)

	// Find pet by Id
	app.Router.Methods("GET").Path("/pet/{petId}").HandlerFunc(petController.ReadPetInfo)

	// Adds new pet to the store
	app.Router.Methods("POST").Path("/pet").HandlerFunc(petController.CreateNewPet)

	// Update an existing pet with form data
	app.Router.Methods("POST").Path("/pet/{petId}").HandlerFunc(petController.UpdatePetWithFormData)

	// Update an existing pet
	app.Router.Methods("PUT").Path("/pet").HandlerFunc(petController.UpdatePetData)

	// Deletes a pet
	app.Router.Methods("DELETE").Path("/pet/{petId}").HandlerFunc(petController.DeletePetData)

	// I have chosen to implement the Pet Endpoint, therefore store and user endpoints are not implemented.
	app.Router.Methods("POST").Path("/store/order").HandlerFunc(NotImplemented)
	app.Router.Methods("GET").Path("/store/order/{orderId}").HandlerFunc(NotImplemented)
	app.Router.Methods("DELETE").Path("/store/order/{orderId}").HandlerFunc(NotImplemented)
	app.Router.Methods("GET").Path("/store/inventory").HandlerFunc(NotImplemented)

	app.Router.Methods("POST").Path("/user/createWithArray").HandlerFunc(NotImplemented)
	app.Router.Methods("POST").Path("​/user​/createWithList").HandlerFunc(NotImplemented)
	app.Router.Methods("GET").Path("/user/{username}").HandlerFunc(NotImplemented)
	app.Router.Methods("PUT").Path("/user/{username}").HandlerFunc(NotImplemented)
	app.Router.Methods("DELETE").Path("/user/{username}").HandlerFunc(NotImplemented)
	app.Router.Methods("GET").Path("/user/login").HandlerFunc(NotImplemented)
	app.Router.Methods("GET").Path("/user/logout").HandlerFunc(NotImplemented)
	app.Router.Methods("POST").Path("/user").HandlerFunc(NotImplemented)

	app.Router.Use(loggingMiddleware)
	app.Router.Use(auth.Middleware)
}

// NotImplemented function implementing the NotImplemented handler. Whenever an API endpoint is hit
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Only Pet endpoints are implemented. Store and User endpoints are not implemented"))
}
