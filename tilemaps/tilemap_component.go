package tilemaps

import (
	"github.com/adm87/finch-core/ecs"
)

var TilemapComponentType = ecs.NewComponentType[*TilemapComponent]()

type TilemapComponent struct {
	TilemapID string
}

func NewTilemapComponent(tilemapID string) *TilemapComponent {
	return &TilemapComponent{
		TilemapID: tilemapID,
	}
}

func NewEmptyTilemapComponent() *TilemapComponent {
	return &TilemapComponent{
		TilemapID: "",
	}
}

func (c *TilemapComponent) Type() ecs.ComponentType {
	return TilemapComponentType
}
