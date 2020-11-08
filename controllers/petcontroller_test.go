package controllers

import (
	"net/http"
	"net/http/httptest"
	"pet_store_rest_api/interfaces/mock_interfaces"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

/*
  Actual test functions
*/

// TestSomething is an example of how to use our test object to
// make assertions about some target code we are testing.
func TestPetController(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPetRepository := mock_interfaces.NewMockIPetRepository(ctrl)
	mockPetRepository.EXPECT().AddNewPet(gomock.Any())

	petController := PetController{mockPetRepository}

	payload := string(`{
		"id": 0,
		"category": {
		  "id": 0,
		  "name": "string"
		},
		"name": "doggie",
		"photoUrls": [
		  "string"
		],
		"tags": [
		  {
			"id": 0,
			"name": "string"
		  }
		],
		"status": "available"
	  }`)

	// call the code we are testing
	req := httptest.NewRequest("POST", "http://localhost:9000/pet", strings.NewReader(payload))
	w := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/pet", petController.CreateNewPet)

	r.ServeHTTP(w, req)

	// Check the status code is what we expect.
	if status := w.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// assert that the expectations were met
	require.JSONEq(t, w.Body.String(), payload)
}
