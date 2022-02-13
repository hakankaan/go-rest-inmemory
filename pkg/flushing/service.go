package flushing

import "github.com/hakankaan/go-rest-inmemory/pkg/logging"

// Service provides pair adding operations
type Service interface {
	// FlushDB removes all key-value pairs from db
	FlushDB() error
}

// Repository defines the rules around repository has to be able to perform
type Repository interface {
	// Flush flushes the whole db
	Flush() error
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

// FlushDB removes all key-value pairs from db
func (s *service) FlushDB() (err error) {
	err = s.r.Flush()
	if err != nil {
		s.l.Error("FlushDB", err)
		return
	}

	return
}
