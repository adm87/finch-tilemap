package tilesets

import (
	"encoding/json"
	"sync"

	"github.com/adm87/finch-core/types"
	"github.com/adm87/finch-resources/storage"
)

var (
	tilesetAssetTypes = types.MakeSetFrom(".tileset")
	cacheInstance     = &TilesetCache{
		mu:    sync.RWMutex{},
		store: storage.NewStore[*Tileset](),
	}
)

type TilesetCache struct {
	mu    sync.RWMutex
	store *storage.Store[*Tileset]
}

func Cache() *TilesetCache {
	return cacheInstance
}

func (c *TilesetCache) Get(key string) (*Tileset, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	tileset, err := c.store.Get(key)
	if err != nil {
		return nil, err
	}

	return tileset, nil
}

func (c *TilesetCache) Allocate(key string, data []byte) error {
	var tileset *Tileset

	if err := json.Unmarshal(data, &tileset); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if err := c.store.Add(key, tileset); err != nil {
		return err
	}

	return nil
}

func (c *TilesetCache) Deallocate(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.store.Remove(key)
}

func (c *TilesetCache) AssetTypes() types.HashSet[string] {
	return tilesetAssetTypes
}

func (c *TilesetCache) SetDefault(key string) error {
	return nil
}
