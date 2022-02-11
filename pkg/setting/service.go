package setting

// Service provides pair adding operations
type Service interface {
	AddPair(...Pair) error
}

// Repository defines the rules around repository has to be able to perform
type Repository interface {
	Set(Pair) error
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

// AddPair persists the given pair(s) to storage
func (s *service) AddPair(p ...Pair) (err error) {
	for _, pp := range p {
		err = s.r.Set(pp)
		if err != nil {
			return
		}
	}

	return
}
