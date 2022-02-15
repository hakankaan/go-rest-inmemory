package flushing_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hakankaan/go-rest-inmemory/pkg/flushing"
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

	for i := 1; i <= 10; i++ {
		p := setting.Pair{
			Key:   fmt.Sprintf("Key%d", i),
			Value: fmt.Sprintf("Val%d", i),
		}
		s.Set(p)
	}

	ts = &testService{r: s, l: l}

}

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	setup()
	return m.Run()
}

func TestFlushDB(t *testing.T) {
	s, err := flushing.NewService(ts.r, ts.l)

	if err != nil {
		t.Error("flushing service could not initialized")
	}

	err = s.FlushDB()
	if err != nil {
		t.Error(err)
	}

	datas, err := ts.r.GetAll()
	if err != nil {
		t.Error(err)
	}

	for i := 1; i <= 10; i++ {
		k := fmt.Sprintf("Key%d", i)
		v, ok := datas[k]
		if ok {
			t.Error(fmt.Sprintf("expected nil got %s", v))
		}
	}

}
