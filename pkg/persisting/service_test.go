package persisting_test

import (
	"fmt"
	"testing"

	"github.com/hakankaan/go-rest-inmemory/pkg/loading"
	"github.com/hakankaan/go-rest-inmemory/pkg/logging"
	"github.com/hakankaan/go-rest-inmemory/pkg/persisting"
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
	setUp()
	m.Run()
	tearDown()
}

func tearDown() {
	ps, _ := persisting.NewService(ts.r, ts.l)
	ps.CleanDisk()
}

func TestWriteToDisk(t *testing.T) {
	ps, err := persisting.NewService(ts.r, ts.l)
	if err != nil {
		t.Error(err)
	}

	err = ps.WriteToDisk()
	if err != nil {
		t.Error(err)
	}

	ls, err := loading.NewService(ts.r, ts.l)
	if err != nil {
		t.Error(err)
	}

	ls.ReadFromDiskIfExists()

	datas, err := ts.r.GetAll()
	if err != nil {
		t.Error(err)
	}

	for i := 1; i <= 10; i++ {
		k := fmt.Sprintf("Key%d", i)
		v, ok := datas[k]
		if !ok {
			t.Error(fmt.Sprintf("expected %s got nil", v))
		}
	}
}
