package storage

import (
	"sync"

	"github.com/hakankaan/go-rest-inmemory/pkg/getting"
	"github.com/hakankaan/go-rest-inmemory/pkg/setting"
)

type KeyValueStore map[string]string

// Memory storage keeps data in memory
type Storage struct {
	datas KeyValueStore
	sync.Mutex
}

// New is a factory function to generate a new in memory storage
func New() *Storage {
	return &Storage{
		datas: make(KeyValueStore),
	}
}

// Flush flushes the whole db
func (m *Storage) Flush() (err error) {
	m.Lock()
	defer m.Unlock()

	m.datas = make(KeyValueStore)

	return
}

// Set sets the value for given key
func (m *Storage) Set(p setting.Pair) (err error) {
	m.Lock()
	defer m.Unlock()

	m.datas[p.Key] = p.Value

	return
}

// Get returns value of given key
func (m *Storage) Get(k string) (v string, err error) {
	m.Lock()
	defer m.Unlock()

	v, ok := m.datas[k]
	if !ok {
		err = getting.ErrNotFound
		return
	}

	return
}

// GetAll returns all datas from storage
func (m *Storage) GetAll() (list KeyValueStore, err error) {
	list = m.datas

	return
}

// SetAll sets all given data as storage.KeyValueStore
func (m *Storage) SetAll(kvs KeyValueStore) {
	m.Lock()
	defer m.Unlock()

	m.datas = kvs
}
