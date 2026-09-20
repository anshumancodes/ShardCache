package cache

import (
	"hash/fnv"
	"sync"
)

type Shard struct {
	data map[string]string // shard data
	evictions []*Eviction
	mu   sync.RWMutex      // lock for the shard
}

// pass a cache key to get the shard
// it converts the key to a hash and uses it to determine the shard index
// now once shard index is determined, it returns the shard from Cache struct
func (c *Cache) GetShard(key string) *Shard {
	h := fnv.New32a()
	h.Write([]byte(key))
	index := h.Sum32() % uint32(len(c.shards))

	return &c.shards[index]
}
