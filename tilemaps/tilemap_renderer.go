package tilemaps

import (
	"image"

	"github.com/adm87/finch-core/components/transform"
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-rendering/rendering"
	"github.com/adm87/finch-resources/images"
	"github.com/adm87/finch-tilemap/tilesets"
	"github.com/hajimehoshi/ebiten/v2"
)

var op = &ebiten.DrawImageOptions{}

func TilemapRenderer(world *ecs.World, entity ecs.Entity) (rendering.RenderingTask, int, error) {
	tilemapComp, _, _ := ecs.GetComponent[*TilemapComponent](world, entity, TilemapComponentType)
	tilemapRend, _, _ := ecs.GetComponent[*TilemapRenderComponent](world, entity, TilemapRenderComponentType)

	if tilemapComp.TilemapID == "" {
		return nil, 0, nil
	}

	tilemap, err := Storage().Get(tilemapComp.TilemapID)
	if err != nil {
		return nil, 0, err
	}

	tilemapBuffer, exists := get_tilemap_buffer(tilemapComp.TilemapID)
	if !exists || tilemap.IsDirty() {
		tileset, err := tilesets.Storage().Get(tilemap.TilesetID)
		if err != nil {
			return nil, 0, err
		}
		if !exists {
			tilemapBuffer = new_tilemap_buffer(tilemapComp.TilemapID, tilemap.Columns*tileset.TileSize, tilemap.Rows*tileset.TileSize)
		}
		tilesetImg, err := images.Storage().Get(tileset.ImageID)
		if err != nil {
			return nil, 0, err
		}
		if err := draw_tilemap(tilemapBuffer, tilesetImg, tilemap, tileset); err != nil {
			return nil, 0, err
		}
	}

	return func(surface *ebiten.Image, view ebiten.GeoM) {
		op.GeoM.Reset()
		if transformComp, ok, _ := ecs.GetComponent[*transform.TransformComponent](world, entity, transform.TransformComponentType); ok {
			op.GeoM.Concat(transformComp.WorldMatrix())
		}
		op.GeoM.Concat(view)
		surface.DrawImage(tilemapBuffer, op)
	}, tilemapRend.ZOrder, nil
}

func draw_tilemap(buffer *ebiten.Image, palette *ebiten.Image, tilemap *Tilemap, tileset *tilesets.Tileset) error {
	buffer.Clear()

	tsw := tileset.Columns * tileset.TileSize
	tsh := tileset.Rows * tileset.TileSize

	// // Draw the tilemap onto the buffer using the palette
	op := &ebiten.DrawImageOptions{}
	for y := 0; y < tilemap.Rows; y++ {
		for x := 0; x < tilemap.Columns; x++ {
			tile := tilemap.GetTile(x, y)

			sx := (tile % tileset.Columns) * tileset.TileSize
			sy := (tile / tileset.Columns) * tileset.TileSize

			if sx+tileset.TileSize > tsw || sy+tileset.TileSize > tsh {
				continue
			}

			op.GeoM.Reset()
			op.GeoM.Translate(float64(x*tileset.TileSize), float64(y*tileset.TileSize))
			buffer.DrawImage(palette.SubImage(image.Rect(sx, sy, sx+tileset.TileSize, sy+tileset.TileSize)).(*ebiten.Image), op)
		}
	}

	tilemap.ClearDirtyTiles()
	return nil
}
