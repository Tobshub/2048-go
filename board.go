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
	for r := 0; r < b.CellCount; r++ {
		new_array[r] = make([]Tile, b.CellCount)
		for c := 0; c < b.CellCount; c++ {
			if (r == 1 && c == 1) || (r == 3 && c == 3) {
				new_array[r][c] = Tile{Value: 2}
			} else {
				new_array[r][c] = Tile{Value: 0}
			}
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

	cell_size := float32(BOARD_SIZE-board_thickness/4) / float32(board.CellCount) // account for board borders
	cell_border_offset := float32(cell_thickness) / 8

	// draw cell
	for r := 0; r < board.CellCount; r++ {
		cell_x := x + float32(r)*(float32(cell_size)+cell_border_offset)

		for c := 0; c < board.CellCount; c++ {
			cell_y := y + float32(c)*(float32(cell_size)+cell_border_offset)

			cell := rl.NewRectangle(cell_x, cell_y, float32(cell_size), float32(cell_size))
			rl.DrawRectangleLinesEx(cell, cell_thickness, rl.Black)

			// draw tile with value
			if board.Array[r][c].Value != 0 {
				DrawTileInCell(&board.Array[r][c], cell_x, cell_y, cell_size, cell_border_offset)
			}
		}
	}
}

func DrawTileInCell(tile *Tile, cell_x float32, cell_y float32, cell_size float32, cell_border_offset float32) {
	tile_offset := cell_border_offset * 4

	tile.X = cell_x + tile_offset
	tile.Y = cell_y + tile_offset

	tile.Draw(cell_size - tile_offset*3)
}
