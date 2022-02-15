package setting_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hakankaan/go-rest-inmemory/pkg/logging"
	"github.com/hakankaan/go-rest-inmemory/pkg/setting"
	"github.com/hakankaan/go-rest-inmemory/pkg/storage"
)

type testService struct {
	r *storage.Storage
	l logging.Service
}

var ts *testService

func setup() {
	s := storage.New()
	l := logging.NewStdoutLogging("DEBUG")

	ts = &testService{r: s, l: l}

}
func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	setup()
	return m.Run()
}

func TestSetValue(t *testing.T) {
	s, err := setting.NewService(ts.r, ts.l)

	if err != nil {
		t.Error("flushing service could not initialized")
	}

	p := setting.Pair{
		Key:   "Key1",
		Value: "Val1",
	}

	err = s.SetValue(p)
	if err != nil {
		t.Error(err)
	}

	v, err := ts.r.Get("Key1")
	if err != nil {
		t.Error(err)
	}
	if v != "Val1" {
		t.Error(fmt.Sprintf("expected Val1 got %s", v))
	}

}
