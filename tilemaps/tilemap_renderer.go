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
	tilemap, err := Cache().Get(tilemapComp.TilemapID)
	if err != nil {
		return nil, 0, err
	}
	tileset, err := tilesets.Cache().Get(tilemap.TilesetID)
	if err != nil {
		return nil, 0, err
	}
	tilemapBuffer, exists := get_tilemap_buffer(tilemapComp.TilemapID)
	if !exists {
		tilemapBuffer = new_tilemap_buffer(tilemapComp.TilemapID, tilemap.Rows*tileset.TileSize, tilemap.Columns*tileset.TileSize)
	}
	tilesetImg, err := images.Cache().Get(tileset.ImageID)
	if err != nil {
		return nil, 0, err
	}
	if tilemap.IsDirty {
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
	}, tilemapComp.ZOrder, nil
}

func draw_tilemap(buffer *ebiten.Image, palette *ebiten.Image, tilemap *Tilemap, tileset *tilesets.Tileset) error {
	buffer.Clear()

	tsw := tileset.Columns * tileset.TileSize
	tsh := tileset.Rows * tileset.TileSize

	// Draw the tilemap onto the buffer using the palette
	op := &ebiten.DrawImageOptions{}
	for i := 0; i < len(tilemap.Data); i++ {
		tileID := tilemap.Data[i]
		if tileID < 0 {
			continue
		}

		sx := (tileID % tileset.Columns) * tileset.TileSize
		sy := (tileID / tileset.Columns) * tileset.TileSize

		if sx+tileset.TileSize > tsw || sy+tileset.TileSize > tsh {
			continue
		}

		tx := (i % tilemap.Columns) * tileset.TileSize
		ty := (i / tilemap.Columns) * tileset.TileSize

		op.GeoM.Reset()
		op.GeoM.Translate(float64(tx), float64(ty))
		buffer.DrawImage(palette.SubImage(image.Rect(sx, sy, sx+tileset.TileSize, sy+tileset.TileSize)).(*ebiten.Image), op)
	}

	tilemap.IsDirty = false
	return nil
}
