package main

import (
	"net/http"

	"github.com/hakankaan/go-rest-inmemory/pkg/flushing"
	"github.com/hakankaan/go-rest-inmemory/pkg/getting"
	"github.com/hakankaan/go-rest-inmemory/pkg/http/rest"
	"github.com/hakankaan/go-rest-inmemory/pkg/loading"
	"github.com/hakankaan/go-rest-inmemory/pkg/persisting"
	"github.com/hakankaan/go-rest-inmemory/pkg/setting"
	"github.com/hakankaan/go-rest-inmemory/pkg/storage"
)

func main() {
	s := storage.New()

	ls, _ := loading.NewService(loading.WithRepository(s))
	ls.ReadFromDiskIfExists()

	sd, _ := persisting.NewService(persisting.WithRepository(s))
	sd.ScheduleWritingToDisk()

	gs, _ := getting.NewService(getting.WithRepository(s))
	ss, _ := setting.NewService(setting.WithRepository(s))
	fs, _ := flushing.NewService(flushing.WithRepository(s))

	http.HandleFunc("/", rest.New(gs, ss, fs))
	http.ListenAndServe(":8080", nil)
}
