package db

type multiStore []Store

type shardedStore struct {
	shardedStore   multiStore
	numberOfShards int
}

func newShardedKeyValueStore(storeType multiStore, numberOfShards int) *shardedStore {
	return &shardedStore{
		storeType,
		numberOfShards,
	}
}

func (ks *shardedStore) Set(key string, value interface{}) {
	keyValueStore := *ks.GetKeyValueStore(key)
	keyValueStore.Set(key, value)
}

func (ks *shardedStore) Get(key string) interface{} {
	keyValueStore := *ks.GetKeyValueStore(key)

	return keyValueStore.Get(key)
}

func (ks *shardedStore) GetKeyValueStore(key string) *Store {
	return &ks.shardedStore[uint(fnv32(key))%uint(ks.numberOfShards)]
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
