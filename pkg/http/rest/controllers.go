package rest

import (
	"encoding/json"
	"net/http"

	"github.com/hakankaan/go-rest-inmemory/pkg/setting"
)

type ctxKey struct{}

func getField(r *http.Request, i int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[i]
}

// flushDB returns a handler for DELETE /data requests
func (rs *restService) flushDB(w http.ResponseWriter, r *http.Request) {

	err := rs.flushingService.FlushDB()
	if err != nil {
		InternalServerErrorResponse(w, err)
		return
	}
	NoContentResponse(w)
}

// getValue returns a handler for GET /data requests
func (rs *restService) getValue(w http.ResponseWriter, r *http.Request) {

	k := getField(r, 0)
	v, err := rs.gettingService.GetValue(k)
	if err != nil {
		InternalServerErrorResponse(w, err)
		return
	}
	OkResponseWithPair(w, k, v)
}

// setValue returns a handler for POST /data requests
func (rs *restService) setValue(w http.ResponseWriter, r *http.Request) {
	var p setting.Pair

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		InternalServerErrorResponse(w, err)
		return
	}

	err = rs.settingService.SetValue(p)
	if err != nil {
		InternalServerErrorResponse(w, err)
		return
	}

	OkResponseWithPair(w, p.Key, p.Value)

}
