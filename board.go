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

type BoardMotion int32

const (
	MotionNone BoardMotion = iota
	MotionLeft
	MotionRight
	MotionUp
	MotionDown
)

type Board struct {
	CellCount int // number of rows, same as number of cols
	Array     [][]Tile
	Motion    BoardMotion
}

func (b *Board) Init() {
	new_array := make([][]Tile, b.CellCount)
	for c := 0; c < b.CellCount; c++ {
		new_array[c] = make([]Tile, b.CellCount)
		for r := 0; r < b.CellCount; r++ {
			if (c == 1 && r == 1) || (c == 3 && r == 3) || (c == 1 && r == 3) {
				new_array[c][r] = Tile{Value: 2}
			} else {
				new_array[c][r] = Tile{Value: 0}
			}
		}
	}
	b.Array = new_array
}

var move_count = 0

const MAX_MOVE_COUNT = 3

func (board *Board) MoveTiles() {
	if move_count >= MAX_MOVE_COUNT {
		move_count = 0
		board.Motion = MotionNone
	} else {
		move_count++
		if board.Motion == MotionLeft || board.Motion == MotionUp {
			for c := 0; c < board.CellCount; c++ {
				for r := 0; r < board.CellCount; r++ {
					switch board.Motion {
					case MotionLeft:
						{
							if c-1 >= 0 && board.Array[c-1][r].Value == 0 {
								board.Array[c-1][r].Value = board.Array[c][r].Value
								board.Array[c][r].Value = 0
							} else if c-1 >= 0 && board.Array[c-1][r].Value == board.Array[c][r].Value {
								board.Array[c-1][r].Value *= 2
								board.Array[c][r].Value = 0
							}
						}
					case MotionUp:
						{
							if r-1 >= 0 && board.Array[c][r-1].Value == 0 {
								board.Array[c][r-1].Value = board.Array[c][r].Value
								board.Array[c][r].Value = 0
							} else if r-1 >= 0 && board.Array[c][r-1].Value == board.Array[c][r].Value {
								board.Array[c][r-1].Value *= 2
								board.Array[c][r].Value = 0
							}
						}
					}
				}
			}
		} else if board.Motion == MotionRight || board.Motion == MotionDown {
			for c := board.CellCount - 1; c >= 0; c-- {
				for r := board.CellCount - 1; r >= 0; r-- {
					switch board.Motion {
					case MotionRight:
						{
							if c+1 < board.CellCount && board.Array[c+1][r].Value == 0 {
								board.Array[c+1][r].Value = board.Array[c][r].Value
								board.Array[c][r].Value = 0
							} else if c+1 < board.CellCount && board.Array[c+1][r].Value == board.Array[c][r].Value {
								board.Array[c+1][r].Value *= 2
								board.Array[c][r].Value = 0
							}
						}
					case MotionDown:
						{
							if r+1 < board.CellCount && board.Array[c][r+1].Value == 0 {
								board.Array[c][r+1].Value = board.Array[c][r].Value
								board.Array[c][r].Value = 0
							} else if r+1 < board.CellCount && board.Array[c][r+1].Value == board.Array[c][r].Value {
								board.Array[c][r+1].Value *= 2
								board.Array[c][r].Value = 0
							}
						}
					}
				}
			}
		}
	}
}

func (board *Board) Update() {
	if board.Motion == MotionNone {
		switch rl.GetKeyPressed() {
		case rl.KeyLeft:
			board.Motion = MotionLeft
		case rl.KeyRight:
			board.Motion = MotionRight
		case rl.KeyUp:
			board.Motion = MotionUp
		case rl.KeyDown:
			board.Motion = MotionDown
		}
	}
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
	for c := 0; c < board.CellCount; c++ {
		cell_x := x + float32(c)*(float32(cell_size)+cell_border_offset)

		for r := 0; r < board.CellCount; r++ {
			cell_y := y + float32(r)*(float32(cell_size)+cell_border_offset)

			cell := rl.NewRectangle(cell_x, cell_y, float32(cell_size), float32(cell_size))
			rl.DrawRectangleLinesEx(cell, cell_thickness, rl.Black)

			// draw tile with value
			if board.Array[c][r].Value != 0 {
				DrawTileInCell(&board.Array[c][r], cell_x, cell_y, cell_size, cell_border_offset)
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
