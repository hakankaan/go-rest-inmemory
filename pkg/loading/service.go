package loading

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

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

// Configuration is an alias for a function that will take in a pointer to an Service and modify it
type Configuration func(as *service) error

// Service is an implementation of the Service
type service struct {
	r Repository
}

// NewService takes a variable amount of Configuration functions and returns a new Service
// Each Configuration will be called in the order they are passed in
func NewService(cfgs ...Configuration) (s *service, err error) {
	s = &service{}

	for _, cfg := range cfgs {
		err = cfg(s)
		if err != nil {
			return
		}
	}

	return
}

// WithRepository applies a given setting repository to the Service
func WithRepository(r Repository) Configuration {
	return func(ss *service) error {
		ss.r = r
		return nil
	}
}

// ReadFromDiskIfExists reads data from disk if .json file exists
func (s *service) ReadFromDiskIfExists() {
	path := defaultFilePath

	files, err := ioutil.ReadDir(path)
	if err != nil {

		return
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())
		f, err := os.ReadFile(path)
		if err != nil {
			return
		}

		loadedData := storage.KeyValueStore{}

		err = json.Unmarshal(f, &loadedData)
		if err != nil {
			return
		}

		s.r.SetAll(loadedData)
	}
}
