package tilemaps

type Tilemap struct {
	IsDirty   bool   `json:"-"`
	Rows      int    `json:"rows"`
	Columns   int    `json:"columns"`
	Data      []int  `json:"data"`
	TilesetID string `json:"tileset_id"`
}
