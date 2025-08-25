package tilesets

type Tileset struct {
	ImageID  string         `json:"image_id"`
	Rows     int            `json:"rows"`
	Columns  int            `json:"columns"`
	TileSize int            `json:"tile_size"`
	Padding  TilesetPadding `json:"padding"`
}

type TilesetPadding struct {
	Top    int `json:"top"`
	Right  int `json:"right"`
	Bottom int `json:"bottom"`
	Left   int `json:"left"`
}
