package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hakankaan/go-rest-inmemory/pkg/logging"
)

// InternalServerErrorResponse returns a response for internal server error
// returns with error code = 500
func InternalServerErrorResponse(w http.ResponseWriter, l logging.Service, err error) {
	err = json.NewEncoder(w).Encode(err.Error())
	if err != nil {
		l.Error("json.Encoder.Encode", err)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

// OkResponseWithPair returns a response for successful request
// returns with error code = 200
func OkResponseWithPair(w http.ResponseWriter, l logging.Service, k, v string) {
	r := fmt.Sprintf("%s=%s", k, v)
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		l.Error("json.Encoder.Encode", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// OkResponseWithValue returns a response for successful request
// returns with error code = 200
func OkResponseWithValue(w http.ResponseWriter, l logging.Service, v string) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		l.Error("json.Encoder.Encode", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// NoContentResponse returns a response if data not found or deleted
// returns with error code = 204
func NoContentResponse(w http.ResponseWriter, l logging.Service) {
	err := json.NewEncoder(w).Encode("no_content")
	if err != nil {
		l.Error("json.Encoder.Encode", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
