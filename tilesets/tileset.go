package tilesets

type Tileset struct {
	ImageId string         `yaml:"image_id"`
	Rows    int            `yaml:"rows"`
	Columns int            `yaml:"columns"`
	Width   int            `yaml:"tile_width"`
	Height  int            `yaml:"tile_height"`
	Padding TilesetPadding `yaml:"padding"`
}

type TilesetPadding struct {
	Top    int `yaml:"top"`
	Right  int `yaml:"right"`
	Bottom int `yaml:"bottom"`
	Left   int `yaml:"left"`
}
