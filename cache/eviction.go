package cache

import (
	"time"
)

type Eviction struct {
	key          string
	creationTime time.Time
}

const TTL = 10 * time.Second

func (e *Eviction) IsExpired() bool {
	return time.Since(e.creationTime) >= TTL
}

func NewEviction(key string) *Eviction {
	return &Eviction{
		key:          key,
		creationTime: time.Now(),
	}
}

func HandleEviction(shard *Shard, evictions []*Eviction) {
	for {
		// check if the oldest eviction is expired
		oldest := evictions[0]

		// if not expired, break the loop because ofc newer ones are not expired yet
		if !oldest.IsExpired() {
			break
		}
		// if the oldest eviction is expired, delete it and move to the next one
		delete(shard.data, oldest.key)
		evictions = evictions[1:]
	}
}
