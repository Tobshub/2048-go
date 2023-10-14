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
	rl.DrawRectangleV(
		rl.NewVector2(tile.X, tile.Y),
		rl.NewVector2(size, size),
		GetTileColor(tile.Value),
	)

	val := fmt.Sprintf("%d", tile.Value)
	font_size := int32(42)

	for float32(rl.MeasureText(val, font_size)) > size {
		font_size -= 4
	}

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

func DrawTileInCell(tile *Tile, cell_x float32, cell_y float32, cell_size float32, cell_border_offset float32) {
	tile_offset := cell_border_offset * 4

	tile.X = cell_x + tile_offset
	tile.Y = cell_y + tile_offset

	tile.Draw(cell_size - tile_offset*3)
}
