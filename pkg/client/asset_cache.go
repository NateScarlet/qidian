package client

import (
	"context"
	"sync"
)

type AssetCache interface {
	Set(ctx context.Context, key string, body []byte) (err error)
	Get(ctx context.Context, key string) (body []byte, ok bool, err error)
}

type contextKeyAssetCache struct{}

func WithAssetCache(ctx context.Context, v AssetCache) context.Context {
	return context.WithValue(ctx, contextKeyAssetCache{}, v)
}

var DefaultAssetCache AssetCache = NewInMemoryAssetCache()

func ContextAssetCache(ctx context.Context) AssetCache {
	var ret, _ = ctx.Value(contextKeyAssetCache{}).(AssetCache)
	if ret == nil {
		return DefaultAssetCache
	}
	return ret
}

type inMemoryAssetCache struct {
	mu sync.Mutex
	m  map[string][]byte
}

// Get implements AssetCache
func (c *inMemoryAssetCache) Get(ctx context.Context, key string) (body []byte, ok bool, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	body, ok = c.m[key]
	return
}

// Set implements AssetCache
func (c *inMemoryAssetCache) Set(ctx context.Context, key string, body []byte) (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = body
	return
}

func NewInMemoryAssetCache() AssetCache {
	return &inMemoryAssetCache{
		m: make(map[string][]byte),
	}
}
