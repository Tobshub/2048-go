package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	padding         = 40
	board_thickness = 4

	cell_thickness = 2
)

var BOARD_SIZE = int32(math.Min(float64(SCREEN_WIDTH), float64(SCREEN_HEIGHT))) - padding

type Board struct {
	CellCount int // number of rows, same as number of cols
	Array     [][]Tile
}

func (b *Board) Init() {
	new_array := make([][]Tile, b.CellCount)
	for i := 0; i < b.CellCount; i++ {
		new_array[i] = make([]Tile, b.CellCount)
		for j := 0; j < b.CellCount; j++ {
			new_array[i][j] = Tile{Value: 1}
		}
	}
	b.Array = new_array
}

func (board *Board) Draw() {
	x := float32(SCREEN_WIDTH/2 - BOARD_SIZE/2)
	y := float32(SCREEN_HEIGHT/2 - BOARD_SIZE/2)

	board_rect := rl.NewRectangle(x, y, float32(BOARD_SIZE), float32(BOARD_SIZE))

	// draw board
	rl.DrawRectangleLinesEx(board_rect, board_thickness, rl.Black)

	cell_size := float32(BOARD_SIZE-board_thickness) / float32(board.CellCount) // account for board borders
	cell_border_offset := float32(cell_thickness) / 4

	// draw cell
	for r := 0; r < board.CellCount; r++ {
		cell_x := x + float32(r)*(float32(cell_size)+cell_border_offset)

		for c := 0; c < board.CellCount; c++ {
			cell_y := y + float32(c)*(float32(cell_size)+cell_border_offset)

			cell := rl.NewRectangle(cell_x, cell_y, float32(cell_size), float32(cell_size))
			rl.DrawRectangleLinesEx(cell, cell_thickness, rl.Black)

			if board.Array[r][c].Value != 0 {
				tile := &board.Array[r][c]

				tile.X = cell_x + cell_border_offset*2
				tile.Y = cell_y + cell_border_offset*2

				tile.Draw(cell_size - cell_border_offset)
			}
		}
	}
}
