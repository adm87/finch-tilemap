package tilemaps

type Tilemap struct {
	IsDirty   bool   `yaml:"-"`
	Rows      int    `yaml:"rows"`
	Columns   int    `yaml:"columns"`
	Data      []int  `yaml:"data"`
	TilesetID string `yaml:"tileset_id"`
}
