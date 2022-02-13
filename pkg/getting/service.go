package getting

import (
	"errors"

	"github.com/hakankaan/go-rest-inmemory/pkg/logging"
)

// ErrNotFound is used when a key could not be found
var ErrNotFound = errors.New("getting: key not found")

// Service provides getting pair operation
type Service interface {
	GetValue(k string) (string, error)
}

// Repository defines the rules around repository has to be able to perform
type Repository interface {
	Get(k string) (string, error)
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

// GetValue gets the value of pair by given key from storage
func (s *service) GetValue(k string) (v string, err error) {
	v, err = s.r.Get(k)

	return
}
