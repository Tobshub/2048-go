package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	X, Y  float32
	Value int
}

func (tile *Tile) Draw(size float32) {
	rl.DrawRectangleV(rl.NewVector2(tile.X, tile.Y), rl.NewVector2(size, size), rl.Blue)

	val := fmt.Sprintf("%d", tile.Value)
	font_size := int32(42)
	x_offset := int32(size/2) - rl.MeasureText(val, font_size)/2
	y_offset := int32(size/2) - font_size/2

	rl.DrawText(val, int32(tile.X)+x_offset, int32(tile.Y)+y_offset, font_size, rl.White)
}
