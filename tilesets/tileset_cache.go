package tilesets

import (
	"sync"

	"github.com/adm87/finch-core/types"
	"github.com/adm87/finch-resources/storage"
	"gopkg.in/yaml.v3"
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

func (c *TilesetCache) Allocate(key string, data []byte) error {
	var tileset *Tileset

	if err := yaml.Unmarshal(data, &tileset); err != nil {
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
