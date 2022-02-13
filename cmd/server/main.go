package main

import (
	"net/http"

	"github.com/hakankaan/go-rest-inmemory/pkg/flushing"
	"github.com/hakankaan/go-rest-inmemory/pkg/getting"
	"github.com/hakankaan/go-rest-inmemory/pkg/http/rest"
	"github.com/hakankaan/go-rest-inmemory/pkg/setting"
	"github.com/hakankaan/go-rest-inmemory/pkg/storage"
)

func main() {
	s := storage.New()

	s.Set(setting.Pair{Key: "genesis", Value: "kaan"})

	gs, _ := getting.NewService(getting.WithRepository(s))
	ss, _ := setting.NewService(setting.WithRepository(s))
	fs, _ := flushing.NewService(flushing.WithRepository(s))

	http.HandleFunc("/", rest.New(gs, ss, fs))
	http.ListenAndServe(":8080", nil)
}
