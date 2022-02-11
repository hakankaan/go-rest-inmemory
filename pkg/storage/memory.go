package memory

import (
	"sync"

	"github.com/hakankaan/go-rest-inmemory/pkg/getting"
	"github.com/hakankaan/go-rest-inmemory/pkg/setting"
)

// Memory storage keeps data in memory
type Storage struct {
	datas map[string]string
	sync.Mutex
}

// New is a factory function to generate a new in memory storage
func New() *Storage {
	return &Storage{
		datas: make(map[string]string),
	}
}

// Set sets the value for given key
func (m *Storage) Set(p setting.Pair) (err error) {
	m.Lock()
	m.datas[p.Key] = p.Value
	m.Unlock()

	return
}

// Get returns value of given key
func (m *Storage) Get(k string) (v string, err error) {

	v, ok := m.datas[k]
	if !ok {
		err = getting.ErrNotFound
		return
	}

	return
}
