package main

import (
	"net/http"

	"github.com/hakankaan/go-rest-inmemory/pkg/flushing"
	"github.com/hakankaan/go-rest-inmemory/pkg/getting"
	"github.com/hakankaan/go-rest-inmemory/pkg/http/rest"
	"github.com/hakankaan/go-rest-inmemory/pkg/loading"
	"github.com/hakankaan/go-rest-inmemory/pkg/logging"
	"github.com/hakankaan/go-rest-inmemory/pkg/persisting"
	"github.com/hakankaan/go-rest-inmemory/pkg/setting"
	"github.com/hakankaan/go-rest-inmemory/pkg/storage"
)

func main() {
	l := logging.NewStdoutLogging("DEBUG")

	s := storage.New()

	ls, _ := loading.NewService(s, l)
	ls.ReadFromDiskIfExists()

	sd, _ := persisting.NewService(s, l)
	sd.ScheduleWritingToDisk()
	l.Info("qwewqe", "dddd")

	gs, _ := getting.NewService(s, l)
	ss, _ := setting.NewService(s, l)
	fs, _ := flushing.NewService(s, l)

	http.HandleFunc("/", rest.New(l, gs, ss, fs))
	l.Error("serve", http.ListenAndServe(":8080", nil))
}
