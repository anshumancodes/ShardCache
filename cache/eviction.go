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
	for _, eviction := range evictions {
		if eviction.IsExpired() {
			// delete from the shard's data map only if the eviction is expired
			delete(shard.data, eviction.key)
		}

	}
}
