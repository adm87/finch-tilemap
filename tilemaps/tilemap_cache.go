package tilemaps

import (
	"sync"

	"github.com/adm87/finch-core/types"
	"github.com/adm87/finch-resources/storage"
	"gopkg.in/yaml.v3"
)

var (
	tilemapAssetTypes = types.MakeSetFrom(".tilemap")
	cacheInstance     = &TilemapCache{
		mu:    sync.RWMutex{},
		store: storage.NewStore[*Tilemap](),
	}
)

type TilemapCache struct {
	mu    sync.RWMutex
	store *storage.Store[*Tilemap]
}

func Cache() *TilemapCache {
	return cacheInstance
}

func (c *TilemapCache) Get(key string) (*Tilemap, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	tilemap, err := c.store.Get(key)
	if err != nil {
		return nil, err
	}

	return tilemap, nil
}

func (c *TilemapCache) Allocate(key string, data []byte) error {
	var tilemap *Tilemap

	if err := yaml.Unmarshal(data, &tilemap); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if err := c.store.Add(key, tilemap); err != nil {
		return err
	}

	return nil
}

func (c *TilemapCache) Deallocate(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.store.Remove(key)
}

func (c *TilemapCache) AssetTypes() types.HashSet[string] {
	return tilemapAssetTypes
}

func (c *TilemapCache) SetDefault(key string) error {
	return nil
}
