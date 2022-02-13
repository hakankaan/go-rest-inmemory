package rest

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/hakankaan/go-rest-inmemory/pkg/flushing"
	"github.com/hakankaan/go-rest-inmemory/pkg/getting"
	"github.com/hakankaan/go-rest-inmemory/pkg/logging"
	"github.com/hakankaan/go-rest-inmemory/pkg/setting"
)

const (
	apiBase = "/api"
)

// newApiRoute returns a route with added /api prefix to pattern
func newApiRoute(method, pattern string, handler http.HandlerFunc) route {
	apiPattern := fmt.Sprintf("%s%s", apiBase, pattern)
	return route{method, regexp.MustCompile("^" + apiPattern + "$"), Logging(handler)}
}

// route struct for routes
type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

// restService contains services for restapi to use
type restService struct {
	gettingService  getting.Service
	settingService  setting.Service
	flushingService flushing.Service
	logger          logging.Service
}

// New creates restService and routelist then uses regex table for routing
func New(l logging.Service, gs getting.Service, ss setting.Service, fs flushing.Service) http.HandlerFunc {
	rest := &restService{
		gettingService:  gs,
		settingService:  ss,
		flushingService: fs,
		logger:          l,
	}
	var routes = []route{
		newApiRoute(http.MethodPost, "/datas", rest.setValue),
		newApiRoute(http.MethodDelete, "/datas/flush", rest.flushDB),
		newApiRoute(http.MethodGet, "/datas/([^/]+)", rest.getValue),
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var allow []string
		for _, route := range routes {
			matches := route.regex.FindStringSubmatch(r.URL.Path)
			if len(matches) > 0 {
				if r.Method != route.method {
					allow = append(allow, route.method)
					continue
				}

				w.Header().Set("Content-Type", "application/json")
				ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
				route.handler(w, r.WithContext(ctx))
				return
			}
		}
		if len(allow) > 0 {
			w.Header().Set("Allow", strings.Join(allow, ", "))
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.NotFound(w, r)
	}
}
