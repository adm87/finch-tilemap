package tilemaps

import "github.com/adm87/finch-core/ecs"

var TilemapRenderComponentType = ecs.NewComponentType[*TilemapRenderComponent]()

type TilemapRenderComponent struct {
	ZOrder int
}

func NewTilemapRenderComponent(zOrder int) *TilemapRenderComponent {
	return &TilemapRenderComponent{
		ZOrder: zOrder,
	}
}

func (c *TilemapRenderComponent) Type() ecs.ComponentType {
	return TilemapRenderComponentType
}
