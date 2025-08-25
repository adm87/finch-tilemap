package tilemaps

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-rendering/rendering"
	"github.com/hajimehoshi/ebiten/v2"
)

var op = &ebiten.DrawImageOptions{}

func TilemapRenderer(world *ecs.World, entity ecs.Entity) (rendering.RenderingTask, int, error) {
	tilemapComp, _, _ := ecs.GetComponent[*TilemapComponent](world, entity, TilemapComponentType)

	return func(surface *ebiten.Image, view ebiten.GeoM) {
		if tilemapComp.IsDirty {
			draw_tilemap()
		}
	}, tilemapComp.ZOrder, nil
}

func draw_tilemap() {

}
