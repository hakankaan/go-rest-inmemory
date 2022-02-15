package persisting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/hakankaan/go-rest-inmemory/pkg/logging"
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

// ScheduleWritingToDisk starts a timer to schedule writeToDisk function
func (s *service) ScheduleWritingToDisk() {
	duration, err := time.ParseDuration(defaultDuration)
	if err != nil {
		duration, _ = time.ParseDuration(defaultDuration)
	}

	ticker := time.NewTicker(duration)

	go func() {
		for range ticker.C {
			err := s.WriteToDisk()
			if err != nil {
				s.l.Error("s.WriteToDisk", err)
			}
		}
	}()
}

// WriteToDisk writes given storage.KeyValueStore to disk as json
func (s *service) WriteToDisk() (err error) {
	p := defaultFilePath

	_, err = os.Stat(p)
	if os.IsNotExist(err) {
		err = os.Mkdir(p, os.ModePerm)
		if err != nil {
			s.l.Error("writeToDisk", err)
			return
		}
		s.l.Info("os.Mkdir", "store file has beed created")
	} else {
		err = s.CleanDisk()
		if err != nil {
			s.l.Error("writeToDisk", err)
			return
		}
	}

	allData, err := s.r.GetAll()
	if err != nil {
		s.l.Error("writeToDisk", err)
		return
	}

	bytes, err := json.Marshal(allData)
	if err != nil {
		s.l.Error("writeToDisk", err)
		return
	}

	fileName := fmt.Sprintf("%d-data.json", time.Now().Unix())

	fullPath := filepath.Join(p, fileName)
	err = ioutil.WriteFile(fullPath, bytes, 0755)
	if err != nil {
		s.l.Error("writeToDisk", err)
		return
	}

	return
}

// cleanDisk removes json file related to store from disk
func (s *service) CleanDisk() (err error) {
	path := defaultFilePath

	files, err := ioutil.ReadDir(path)
	if err != nil {
		s.l.Error("writeToDisk", err)
		return
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())
		err = os.Remove(path)
		if err != nil {
			s.l.Error("writeToDisk", err)
			return
		}
	}

	return
}
