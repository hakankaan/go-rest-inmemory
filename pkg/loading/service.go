package loading

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hakankaan/go-rest-inmemory/pkg/logging"
	"github.com/hakankaan/go-rest-inmemory/pkg/storage"
)

const (
	defaultFilePath = "./store"
)

// Service provides loading data from file operation
type Service interface {
	ReadFromDiskIfExists() error
}

// Repository defines the rules around repository has to be able to perform
type Repository interface {
	SetAll(storage.KeyValueStore)
}

// Service is an implementation of the Service
type service struct {
	r Repository
	l logging.Service
}

// NewService takes a variable amount of Configuration functions and returns a new Service
// Each Configuration will be called in the order they are passed in
func NewService(r Repository, l logging.Service) (s *service, err error) {
	s = &service{
		r: r,
		l: l,
	}

	return
}

// ReadFromDiskIfExists reads data from disk if .json file exists
func (s *service) ReadFromDiskIfExists() {
	path := defaultFilePath

	files, err := ioutil.ReadDir(path)
	if err != nil {
		s.l.Error("ReadFromDiskIfExists", err)
		return
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())
		f, err := os.ReadFile(path)
		if err != nil {
			s.l.Error("ReadFromDiskIfExists", err)
			return
		}

		loadedData := storage.KeyValueStore{}

		err = json.Unmarshal(f, &loadedData)
		if err != nil {
			s.l.Error("ReadFromDiskIfExists", err)
			return
		}

		s.r.SetAll(loadedData)
	}
}
