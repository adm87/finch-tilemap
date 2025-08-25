package tilesets

type Tileset struct {
	ImageID  string         `yaml:"image_id"`
	Rows     int            `yaml:"rows"`
	Columns  int            `yaml:"columns"`
	TileSize int            `yaml:"tile_size"`
	Padding  TilesetPadding `yaml:"padding"`
}

type TilesetPadding struct {
	Top    int `yaml:"top"`
	Right  int `yaml:"right"`
	Bottom int `yaml:"bottom"`
	Left   int `yaml:"left"`
}
