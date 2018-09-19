package db

import (
	"sync"
)

type RapidGoDB interface {
	Store
}

type rapidGoDB struct {
	store Store
}

var rapidGoDBInstance *rapidGoDB = nil

func GetRapidGoDB(config Config) RapidGoDB {
	if rapidGoDBInstance == nil {
		var store Store = getStore(config)

		if config.NumberOfShards > 0 {
			if config.Concurrency == "none" {
				panic("Sharded KeyStore requires concurrency different than none.")
			}

			innerStore := getShardedStore(config)
			store = newShardedKeyValueStore(innerStore, config.NumberOfShards)
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

func getLocker(config Config) RWLocker {
	if config.Concurrency == "none" {
		return newBrokenLocker()
	} else if config.Concurrency == "mutex" {
		return newLocker()
	} else if config.Concurrency == "rwmutex" {
		return &sync.RWMutex{}
	} else if config.Concurrency == "eager" {
		return newEagerLocker()
	} else {
		panic("Invalid Configuration: Concurrency param has incorrect value.")
	}
}

func getStore(config Config) Store {
	if config.BaseStore == "keyvalue" {
		return newKeyValueStore(getLocker(config))
	} else {
		panic("Invalid Configuration: BaseStore param has incorrect value.")
	}
}

func getShardedStore(config Config) multiStore {
	if config.BaseStore == "keyvalue" {
		storeArray := make(multiStore, config.NumberOfShards)
		for i := 0; i < config.NumberOfShards; i++ {
			storeArray[i] = newKeyValueStore(getLocker(config))
		}
		return storeArray
	} else {
		panic("Invalid Configuration: BaseStore param has incorrect value.")
	}
}
