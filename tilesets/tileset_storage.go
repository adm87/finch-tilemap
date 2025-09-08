package tilesets

import (
	"encoding/json"
	"sync"

	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/types"
	"github.com/adm87/finch-resources/storage"
)

var (
	assetTypes      = types.MakeSetFrom(".tileset")
	storageInstance = &TilesetStorage{
		mu:    sync.RWMutex{},
		store: storage.NewStore[*Tileset](),
	}
)

type TilesetStorage struct {
	mu    sync.RWMutex
	store *storage.Store[*Tileset]
}

func Storage() *TilesetStorage {
	return storageInstance
}

func (c *TilesetStorage) Get(key string) (*Tileset, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	tileset, err := c.store.Get(key)
	if err != nil {
		return nil, err
	}

	return tileset, nil
}

func (c *TilesetStorage) Allocate(key string, data []byte) error {
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

func (c *TilesetStorage) Put(key string, value any) error {
	tileset, ok := value.(*Tileset)
	if !ok {
		return errors.NewInvalidArgumentError("value must be of type *Tileset")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if err := c.store.Set(key, tileset); err != nil {
		return err
	}

	return nil
}

func (c *TilesetStorage) Deallocate(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.store.Remove(key)
}

func (c *TilesetStorage) AssetTypes() types.HashSet[string] {
	return assetTypes
}

func (c *TilesetStorage) SetDefault(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	has, err := c.store.Has(key)
	if err != nil {
		return err
	}

	if !has {
		return errors.NewNotFoundError("default tileset not found in storage: " + key)
	}

	c.store.SetDefault(key)
	return nil
}

func (c *TilesetStorage) DefaultKey() string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.store.Default()
}

func (t *TilesetStorage) Has(key string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

	exists, _ := t.store.Has(key)
	return exists
}
