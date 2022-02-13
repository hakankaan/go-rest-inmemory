package persisting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/hakankaan/go-rest-inmemory/pkg/storage"
)

const (
	defaultFilePath = "./store"
	defaultDuration = "1s"
)

// Service provides persisting data to disk operations
type Service interface {
	ScheduleWritingToDisk() error
	writeToDisk() error
}

// Repository defines the rules around repository has to be able to perform
type Repository interface {
	GetAll() (storage.KeyValueStore, error)
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

// ScheduleWritingToDisk starts a timer to schedule writeToDisk function
func (s *service) ScheduleWritingToDisk() {
	duration, err := time.ParseDuration(defaultDuration)
	if err != nil {
		duration, _ = time.ParseDuration(defaultDuration)
	}

	ticker := time.NewTicker(duration)

	go func() {
		for range ticker.C {
			s.writeToDisk()
		}
	}()
}

// writeToDisk writes given storage.KeyValueStore to disk as json
func (s *service) writeToDisk() (err error) {
	p := defaultFilePath

	_, err = os.Stat(p)
	if os.IsNotExist(err) {
		err = os.Mkdir(p, os.ModePerm)
		if err != nil {
			return
		}
	} else {
		err = s.cleanDisk()
		if err != nil {
			return
		}
	}

	allData, err := s.r.GetAll()
	if err != nil {
		return
	}

	bytes, err := json.Marshal(allData)
	if err != nil {
		return
	}

	fileName := fmt.Sprintf("%d-data.json", time.Now().Unix())

	fullPath := filepath.Join(p, fileName)
	err = ioutil.WriteFile(fullPath, bytes, 0755)
	if err != nil {
		return
	}

	return
}

// cleanDisk removes json file related to store from disk
func (s *service) cleanDisk() (err error) {
	path := defaultFilePath

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())
		err = os.Remove(path)
		if err != nil {
			return
		}
	}

	return
}
