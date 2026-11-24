package coalesce

import (
	"context"
	"sync"
	"time"
)

type DataLoaded struct {
	result    []byte
	err       error
	done      chan struct{}
	expiresAt time.Time
}

type Cache struct {
	mu      sync.Mutex
	entries map[string]*DataLoaded // The data reference
	loader  Loader
	ttl     time.Duration
}

type Loader func(ctx context.Context, key string) ([]byte, error)

// Goroutine with Close() starts here
func New(loader Loader, ttl time.Duration) *Cache {
	entries := make(map[string]*DataLoaded)
	return &Cache{
		mu: sync.Mutex,
		entries: entries,
		loader: loader,
		ttl: ttl,
	}
}

// Goroutine that run the Loader() starts here
func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if val, ok := c.entries[key]; ok { // The key exists, but the loader maybe still havent returned the value
		if val.result = [] {
			
		}
		return c.entries[key].result, nil
	}
	
	entry := DataLoaded{
		result: []byte{},
		err: nil,
		done: make(chan struct{}),
		ttl: time.Now().Add(c.ttl),
	}
	
	c.entries[key] = &entry
	
	go func() {
		
	}
	
	
}

func (c *Cache) Close()
