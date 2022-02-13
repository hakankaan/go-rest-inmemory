package getting_test

import (
	"fmt"
	"testing"

	"github.com/hakankaan/go-rest-inmemory/pkg/getting"
	"github.com/hakankaan/go-rest-inmemory/pkg/logging"
	"github.com/hakankaan/go-rest-inmemory/pkg/setting"
	"github.com/hakankaan/go-rest-inmemory/pkg/storage"
)

type testService struct {
	r *storage.Storage
	l logging.Service
}

var ts *testService

func setUp() {
	s := storage.New()
	l := logging.NewStdoutLogging("DEBUG")

	p := setting.Pair{
		Key:   "Key1",
		Value: "Val1",
	}
	s.Set(p)

	ts = &testService{r: s, l: l}

}

func TestMain(m *testing.M) {
	setUp()
	m.Run()
}

func TestGetValue(t *testing.T) {
	s, err := getting.NewService(ts.r, ts.l)

	if err != nil {
		t.Error("flushing service could not initialized")
	}

	v, err := s.GetValue("Key1")
	if err != nil {
		t.Error(err)
	}

	if v != "Val1" {
		t.Error(fmt.Sprintf("expected Val1 got %s", v))
	}

}
