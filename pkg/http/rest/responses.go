package rest

import (
	"encoding/json"
	"net/http"

	"github.com/hakankaan/go-rest-inmemory/pkg/logging"
)

type BaseResponse struct {
	Success string `json:"success"`
}

type ResponseWithKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	BaseResponse
}

type ResponseWithMsg struct {
	Msg string `json:"message"`
	BaseResponse
}

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
	r := ResponseWithKeyValue{
		Key:   k,
		Value: v,
		BaseResponse: BaseResponse{
			Success: "true",
		},
	}
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		l.Error("json.Encoder.Encode", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// NoContentResponse returns a response if data not found or deleted
// returns with error code = 204
func NoContentResponse(w http.ResponseWriter, l logging.Service) {
	r := ResponseWithMsg{
		Msg: "No Content",
		BaseResponse: BaseResponse{
			Success: "true",
		},
	}
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		l.Error("json.Encoder.Encode", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
