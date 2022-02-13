package getting

import "errors"

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

// WithRepository applies a given getting repository to the Service
func WithRepository(r Repository) (cfg Configuration) {
	return func(ss *service) (err error) {
		ss.r = r
		return
	}
}

// GetValue gets the value of pair by given key from storage
func (s *service) GetValue(k string) (v string, err error) {
	v, err = s.r.Get(k)

	return
}
