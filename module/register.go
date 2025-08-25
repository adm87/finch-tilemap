package module

import (
	"github.com/adm87/finch-core/ecs"
	"github.com/adm87/finch-tilemap/tilemaps"
	"github.com/adm87/finch-tilemap/tilesets"

	rendering "github.com/adm87/finch-rendering/rendering"
	resources "github.com/adm87/finch-resources/storage"
)

func RegisterModule() error {
	if err := resources.RegisterCache(
		tilemaps.Cache(),
		tilesets.Cache(),
	); err != nil {
		return err
	}
	if err := rendering.RegisterRenderers(map[ecs.ComponentType]rendering.Renderer{
		tilemaps.TilemapComponentType: tilemaps.TilemapRenderer,
	}); err != nil {
		return err
	}
	return nil
}
