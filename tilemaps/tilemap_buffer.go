package tilemaps

import "github.com/hajimehoshi/ebiten/v2"

var tilemapBuffers = make(map[string]*ebiten.Image)

func new_tilemap_buffer(tilemapID string, width, height int) *ebiten.Image {
	if buffer, exists := tilemapBuffers[tilemapID]; exists {
		return buffer
	}
	buffer := ebiten.NewImage(width, height)
	tilemapBuffers[tilemapID] = buffer
	return buffer
}

func get_tilemap_buffer(tilemapID string) (*ebiten.Image, bool) {
	buffer, exists := tilemapBuffers[tilemapID]
	return buffer, exists
}

func delete_tilemap_buffer(tilemapID string) {
	buffer, exists := tilemapBuffers[tilemapID]
	if !exists {
		return
	}
	buffer.Deallocate()
	delete(tilemapBuffers, tilemapID)
}
