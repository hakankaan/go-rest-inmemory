package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// InternalServerErrorResponse returns a response for internal server error
// returns with error code = 500
func InternalServerErrorResponse(w http.ResponseWriter, err error) {
	json.NewEncoder(w).Encode(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
}

// OkResponseWithPair returns a response for successful request
// returns with error code = 200
func OkResponseWithPair(w http.ResponseWriter, k, v string) {
	r := fmt.Sprintf("%s=%s", k, v)
	json.NewEncoder(w).Encode(r)
	w.WriteHeader(http.StatusOK)
}

// OkResponseWithValue returns a response for successful request
// returns with error code = 200
func OkResponseWithValue(w http.ResponseWriter, v string) {
	r := fmt.Sprintf("%s", v)
	json.NewEncoder(w).Encode(r)
	w.WriteHeader(http.StatusOK)
}

// NoContentResponse returns a response if data not found or deleted
// returns with error code = 204
func NoContentResponse(w http.ResponseWriter) {
	json.NewEncoder(w).Encode("no_content")
	w.WriteHeader(http.StatusNoContent)
}
