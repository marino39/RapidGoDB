package db

import (
	"sync"
)

type Store interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

type keyValueStore struct {
	store  map[string]interface{}
	locker sync.Locker
}

func newKeyValueStore(locker sync.Locker) *keyValueStore {
	return &keyValueStore{make(map[string]interface{}), locker}
}

func (ks *keyValueStore) Set(key string, value interface{}) {
	ks.locker.Lock()
	ks.store[key] = value
	ks.locker.Unlock()
}

func (ks *keyValueStore) Get(key string) interface{} {
	ks.locker.Lock()
	defer ks.locker.Unlock()
	return ks.store[key]
}
