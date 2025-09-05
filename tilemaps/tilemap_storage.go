package tilemaps

import (
	"encoding/json"
	"sync"

	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/types"
	"github.com/adm87/finch-resources/storage"
)

var (
	assetTypes      = types.MakeSetFrom(".tilemap")
	storageInstance = &TilemapStorage{
		mu:    sync.RWMutex{},
		store: storage.NewStore[*Tilemap](),
	}
)

type TilemapStorage struct {
	mu    sync.RWMutex
	store *storage.Store[*Tilemap]
}

func Storage() *TilemapStorage {
	return storageInstance
}

func (c *TilemapStorage) Get(key string) (*Tilemap, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	tilemap, err := c.store.Get(key)
	if err != nil {
		return nil, err
	}

	return tilemap, nil
}

func (c *TilemapStorage) Allocate(key string, data []byte) error {
	var tilemap *Tilemap

	if err := json.Unmarshal(data, &tilemap); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if err := c.store.Add(key, tilemap); err != nil {
		return err
	}

	if tilemap.Size() == 0 {
		tilemap.Fill(-1)
	}

	return nil
}

func (c *TilemapStorage) PutValue(key string, value any) error {
	tilemap, ok := value.(*Tilemap)
	if !ok {
		return errors.NewInvalidArgumentError("value must be of type *Tilemap")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if err := c.store.Add(key, tilemap); err != nil {
		return err
	}

	return nil
}

func (c *TilemapStorage) Deallocate(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.store.Remove(key)
}

func (c *TilemapStorage) AssetTypes() types.HashSet[string] {
	return assetTypes
}

func (c *TilemapStorage) SetDefault(key string) error {
	return nil
}
