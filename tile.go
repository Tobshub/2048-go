package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	X, Y  float32
	Value int
	// if tile value becomes eligible to add during the move
	// tile gets added; which is unwanted behaviour
	//
	// CanAdd is needed to avoid this behaviour
	// and make sure a tile has only been added to once
	CanAdd bool
}

func (tile *Tile) Draw(size float32) {
	tile_color := GetTileColor(tile.Value)
	rl.DrawRectangleV(rl.NewVector2(tile.X, tile.Y), rl.NewVector2(size, size), tile_color)

	val := fmt.Sprintf("%d", tile.Value)
	font_size := int32(42)
	x_offset := int32(size/2) - rl.MeasureText(val, font_size)/2
	y_offset := int32(size/2) - font_size/2

	rl.DrawText(val, int32(tile.X)+x_offset, int32(tile.Y)+y_offset, font_size, rl.White)
}

func GetTileColor(val int) rl.Color {
	var color rl.Color
	switch val {
	case 2:
		color = rl.Blue
	case 4:
		color = rl.Green
	case 8:
		color = rl.Orange
	case 16:
		color = rl.Red
	case 32:
		color = rl.Yellow
	case 64:
		color = rl.Purple
	case 128:
		color = rl.Brown
	case 256:
		color = rl.Violet

		// TODO do more

	default:
		color = rl.Black
	}
	return color
}
