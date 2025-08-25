package tilemaps

type Tilemap struct {
	Rows    int            `yaml:"rows"`
	Columns int            `yaml:"columns"`
	Layers  []TilemapLayer `yaml:"layers"`
}

type TilemapLayer struct {
	TilesetID string   `yaml:"tileset_id"`
	Data      []string `yaml:"data"`
}
