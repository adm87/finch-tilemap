package tilemaps

import (
	"github.com/adm87/finch-resources/images"
	"github.com/hajimehoshi/ebiten/v2"
)

func new_tilemap_buffer(tilemapID string, width, height int) *ebiten.Image {
	if images.Storage().Has(tilemapID) {
		img, _ := images.Storage().Get(tilemapID)
		return img
	}

	img := ebiten.NewImage(width, height)
	if err := images.Storage().PutValue(tilemapID, img); err != nil {
		panic(err)
	}
	return img
}

func get_tilemap_buffer(tilemapID string) (*ebiten.Image, bool) {
	if images.Storage().Has(tilemapID) {
		img, _ := images.Storage().Get(tilemapID)
		return img, true
	}
	return nil, false
}

func delete_tilemap_buffer(tilemapID string) {
	if !images.Storage().Has(tilemapID) {
		return
	}
	images.Storage().Deallocate(tilemapID)
}
