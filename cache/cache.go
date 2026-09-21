package cache

import "time"

type Cache struct {
	shards []Shard // cache is made of shards (slice of shards)
}

// cache constructor
func New(numShards int) *Cache {
	shards := make([]Shard, numShards) // returns a slice of shards
	for i := range shards {
		shards[i].data = make(map[string]string)
	}
	return &Cache{shards: shards}
}

// passing key and value to set function
// then the key is used to get the shard on which the value is to be set , using the GetShard function
// once the shard is obtained , we lock it and set the value
// then unlock it
func (c *Cache) Set(key, value string) {
	shard := c.GetShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	shard.data[key] = value
	// for shards eviction queue if key already exists in the queue then pick it and update its creation time
	for _, eviction := range shard.evictions {
		if eviction.key == key {
			eviction.creationTime = time.Now()
			HandleEviction(shard, shard.evictions)
			return
		}
	}
	// if the key does not exist in the eviction queue, add it
	shard.evictions = append(shard.evictions, NewEviction(key))
	// handles evictions while setting the value , because lock is held
	HandleEviction(shard, shard.evictions)

}
func (c *Cache) Get(key string) (string, bool) {

	shard := c.GetShard(key)
	shard.mu.RUnlock()
	defer shard.mu.RUnlock()

	value, ok := shard.data[key]

	if !ok {
		return "", false
	}

	return value, true
}
