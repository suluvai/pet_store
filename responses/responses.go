package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet_store_rest_api/models"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, models.ApiResponse{
			Code:    int32(statusCode),
			Type:    "error",
			Message: err.Error()})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
