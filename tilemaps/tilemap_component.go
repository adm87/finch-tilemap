package tilemaps

import (
	"github.com/adm87/finch-core/ecs"
)

var TilemapComponentType = ecs.NewComponentType[*TilemapComponent]()

type TilemapComponent struct {
	ZOrder    int
	TilemapID string
	IsDirty   bool
}

func NewTilemapComponent(tilemapID string, zOrder int) *TilemapComponent {
	return &TilemapComponent{
		ZOrder:    zOrder,
		TilemapID: tilemapID,
		IsDirty:   true,
	}
}

func (c *TilemapComponent) Type() ecs.ComponentType {
	return TilemapComponentType
}
