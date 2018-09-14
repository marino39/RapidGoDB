package db

import "sync"

type RapidGoDB interface {
	Store
}

type rapidGoDB struct {
	store Store
}

var rapidGoDBInstance *rapidGoDB = nil

func GetRapidGoDB(config Config) RapidGoDB {
	if rapidGoDBInstance == nil {
		var store Store = nil
		var locker sync.Locker = nil

		if config.Concurrency == "none" {
			locker = newBrokenLocker()
		} else {
			panic("Invalid Configuration: Concurrency param has incorrect value.")
		}

		if config.BaseStore == "keyvalue" {
			store = newKeyValueStore(locker)
		} else {
			panic("Invalid Configuration: BaseStore param has incorrect value.")
		}

		rapidGoDBInstance = &rapidGoDB{store}
	}
	return rapidGoDBInstance
}

func (rgdb *rapidGoDB) Set(key string, value interface{}) {
	rgdb.store.Set(key, value)
}

func (rgdb *rapidGoDB) Get(key string) interface{} {
	return rgdb.store.Get(key)
}
