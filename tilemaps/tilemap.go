package tilemaps

import (
	"encoding/json"
	"fmt"

	"github.com/adm87/finch-core/errors"
	"github.com/adm87/finch-core/types"
)

type Tilemap struct {
	data       []int
	dirtyTiles types.HashSet[int]

	Rows      int    `json:"rows"`
	Columns   int    `json:"columns"`
	TilesetID string `json:"tileset_id"`
}

func NewTileMap(rows, columns int, tilesetID string) *Tilemap {
	t := &Tilemap{
		data:       make([]int, rows*columns),
		dirtyTiles: make(types.HashSet[int]),
		Rows:       rows,
		Columns:    columns,
		TilesetID:  tilesetID,
	}
	t.Fill(-1)
	return t
}

func (t *Tilemap) IsDirty() bool {
	return !t.dirtyTiles.IsEmpty()
}

func (t *Tilemap) Size() int {
	return len(t.data)
}

func (t *Tilemap) Fill(tile int) {
	t.data = make([]int, t.Rows*t.Columns)
	for i := range t.data {
		t.data[i] = tile
		t.dirtyTiles.AddUnique(i)
	}
}

func (t *Tilemap) GetTile(x, y int) (tile int) {
	tile = t.data[index(x, y, t.Rows, t.Columns)]
	if tile < 0 || tile >= len(t.data) {
		tile = -1
	}
	return
}

func (t *Tilemap) SetTile(x, y, tile int) {
	i := index(x, y, t.Rows, t.Columns)
	if i < 0 || i >= len(t.data) {
		return
	}
	t.data[i] = tile
	t.dirtyTiles.AddUnique(i)
}

func (t *Tilemap) ClearDirtyTiles() {
	t.dirtyTiles.Clear()
}

func index(x, y, rows, columns int) int {
	return y*columns + x
}

func (t *Tilemap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Data      []int  `json:"data"`
		Rows      int    `json:"rows"`
		Columns   int    `json:"columns"`
		TilesetID string `json:"tileset_id"`
	}{
		Data:      t.data,
		Rows:      t.Rows,
		Columns:   t.Columns,
		TilesetID: t.TilesetID,
	})
}

func (t *Tilemap) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Data      []int  `json:"data"`
		Rows      int    `json:"rows"`
		Columns   int    `json:"columns"`
		TilesetID string `json:"tileset_id"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	length := len(tmp.Data)
	if length != tmp.Rows*tmp.Columns {
		return errors.NewUnmarshalError(fmt.Sprintf("data length %d does not match rows*columns %d", length, tmp.Rows*tmp.Columns))
	}
	if length > 0 {
		t.data = tmp.Data
	}
	t.Rows = tmp.Rows
	t.Columns = tmp.Columns
	t.TilesetID = tmp.TilesetID
	t.dirtyTiles = make(types.HashSet[int])
	return nil
}
