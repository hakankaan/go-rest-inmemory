package setting

import "github.com/hakankaan/go-rest-inmemory/pkg/logging"

// Service provides pair adding operations
type Service interface {
	SetValue(Pair) error
}

// Repository defines the rules around repository has to be able to perform
type Repository interface {
	Set(Pair) error
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

// SetValue persists the given pair to storage
func (s *service) SetValue(p Pair) (err error) {
	err = s.r.Set(p)
	if err != nil {
		s.l.Error("SetValue", err)
		return
	}

	return
}
