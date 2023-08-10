package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	padding         = 200
	board_thickness = 8

	cell_thickness = 4
)

var (
	BOARD_SIZE         = int32(math.Min(float64(SCREEN_WIDTH), float64(SCREEN_HEIGHT))) - padding
	tile_did_move bool = false
)

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

	SavedStates [][][]Tile
}

func (b *Board) Init() {
	new_array := make([][]Tile, b.CellCount)
	for c := 0; c < b.CellCount; c++ {
		new_array[c] = make([]Tile, b.CellCount)
		for r := 0; r < b.CellCount; r++ {
			if (c == 1 && r == 1) || (c == 3 && r == 3) /** || (c == 1 && r == 3) */ {
				new_array[c][r] = Tile{Value: 2, CanAdd: true}
			} else {
				new_array[c][r] = Tile{Value: 0, CanAdd: true}
			}
		}
	}
	b.Array = new_array
}

func (board *Board) SpawnTile() {
	pow := float64(rand.Intn(2))

	spawn_val := 2 * math.Pow(2, pow)

	var empty_cell_idx [][]int = [][]int{} // []c,r

	Score = 0
	for c := 0; c < board.CellCount; c++ {
		for r := 0; r < board.CellCount; r++ {
			// reset CanAdd for all tiles
			board.Array[c][r].CanAdd = true
			// make score sum of all tiles
			Score += board.Array[c][r].Value

			if Score > HiScore {
				HiScore = Score
			}

			// filter empty cells
			if board.Array[c][r].Value == 0 {
				empty_cell_idx = append(empty_cell_idx, []int{c, r})
			}
		}
	}

	empty_cell_count := len(empty_cell_idx)

	if empty_cell_count >= 1 {
		rand_empty_cell_idx := rand.Intn(empty_cell_count)

		c := empty_cell_idx[rand_empty_cell_idx][0]
		r := empty_cell_idx[rand_empty_cell_idx][1]

		board.Array[c][r].Value = int(spawn_val)
	}

	if empty_cell_count <= 1 {
		HasLost = !board.CheckHasMoves()
	}
}

var move_count = 0

func (board *Board) MoveTiles() {
	if move_count >= board.CellCount-1 {
		move_count = 0
		board.Motion = MotionNone
		if tile_did_move {
			board.SpawnTile()
			tile_did_move = false
		} else {
			var empty_cell_idx [][]int = [][]int{} // []c,r

			for c := 0; c < board.CellCount; c++ {
				for r := 0; r < board.CellCount; r++ {
					// filter empty cells
					if board.Array[c][r].Value == 0 {
						empty_cell_idx = append(empty_cell_idx, []int{c, r})
					}
				}
			}

			empty_cell_count := len(empty_cell_idx)

			if empty_cell_count == 0 {
				HasLost = !board.CheckHasMoves()
			}
		}
	} else {
		move_count++
		if board.Motion == MotionLeft || board.Motion == MotionUp {
			for c := 0; c < board.CellCount; c++ {
				for r := 0; r < board.CellCount; r++ {
					if board.Array[c][r].Value == 0 {
						continue
					}
					switch board.Motion {
					case MotionLeft:
						{
							if c-1 >= 0 && board.Array[c-1][r].Value == 0 {
								board.Array[c-1][r] = board.Array[c][r]
								board.Array[c][r] = Tile{Value: 0, CanAdd: true}
								tile_did_move = true
							} else if c-1 >= 0 && CheckCanAdd(&board.Array[c][r], &board.Array[c-1][r]) {
								board.Array[c-1][r].Value *= 2
								board.Array[c][r].Value = 0
								board.Array[c-1][r].CanAdd = false
								tile_did_move = true
							}
						}
					case MotionUp:
						{
							if r-1 >= 0 && board.Array[c][r-1].Value == 0 {
								board.Array[c][r-1] = board.Array[c][r]
								board.Array[c][r] = Tile{Value: 0, CanAdd: true}
								tile_did_move = true
							} else if r-1 >= 0 && CheckCanAdd(&board.Array[c][r], &board.Array[c][r-1]) {
								board.Array[c][r-1].Value *= 2
								board.Array[c][r].Value = 0
								board.Array[c][r-1].CanAdd = false
								tile_did_move = true
							}
						}
					}
				}
			}
		} else if board.Motion == MotionRight || board.Motion == MotionDown {
			for c := board.CellCount - 1; c >= 0; c-- {
				for r := board.CellCount - 1; r >= 0; r-- {
					if board.Array[c][r].Value == 0 {
						continue
					}
					switch board.Motion {
					case MotionRight:
						{
							if c+1 < board.CellCount && board.Array[c+1][r].Value == 0 {
								board.Array[c+1][r] = board.Array[c][r]
								board.Array[c][r] = Tile{Value: 0, CanAdd: true}
								tile_did_move = true
							} else if c+1 < board.CellCount && CheckCanAdd(&board.Array[c][r], &board.Array[c+1][r]) {
								board.Array[c+1][r].Value *= 2
								board.Array[c][r].Value = 0
								board.Array[c+1][r].CanAdd = false
								tile_did_move = true
							}
						}
					case MotionDown:
						{
							if r+1 < board.CellCount && board.Array[c][r+1].Value == 0 {
								board.Array[c][r+1] = board.Array[c][r]
								board.Array[c][r] = Tile{Value: 0, CanAdd: true}
								tile_did_move = true
							} else if r+1 < board.CellCount && CheckCanAdd(&board.Array[c][r], &board.Array[c][r+1]) {
								board.Array[c][r+1].Value *= 2
								board.Array[c][r].Value = 0
								board.Array[c][r+1].CanAdd = false
								tile_did_move = true
							}
						}
					}
				}
			}
		}
	}
}

func CheckCanAdd(tile1 *Tile, tile2 *Tile) bool {
	if tile1.CanAdd && tile2.CanAdd && tile1.Value == tile2.Value {
		return true
	} else {
		return false
	}
}

// check around tile
func (board *Board) CheckHasMoves() bool {
	for c := 0; c < board.CellCount; c++ {
		for r := 0; r < board.CellCount; r++ {
			tile_val := board.Array[c][r].Value
			if c > 0 && board.Array[c-1][r].Value == tile_val {
				return true
			}
			if c < board.CellCount-1 && board.Array[c+1][r].Value == tile_val {
				return true
			}
			if r > 0 && board.Array[c][r-1].Value == tile_val {
				return true
			}
			if r < board.CellCount-1 && board.Array[c][r+1].Value == tile_val {
				return true
			}
		}
	}
	return false
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

		if board.Motion != MotionNone {
			board.SaveState()
		}
	}
}

func (board *Board) SaveState() {
	board_copy := make([][]Tile, board.CellCount)

	// deep copy ???
	for c := 0; c < board.CellCount; c++ {
		for r := 0; r < board.CellCount; r++ {
			board_copy[c] = make([]Tile, board.CellCount)
			copy(board_copy[c], board.Array[c])
		}
	}

	board.SavedStates = append(board.SavedStates, board_copy)
}

func (board *Board) UndoState() {
	saved_len := len(board.SavedStates)

	if saved_len > 0 {
		board.Array = board.SavedStates[saved_len-1]
		board.SavedStates = board.SavedStates[:saved_len-1]
	}

	board.Motion = MotionNone

	// revert score
	Score = 0
	for c := 0; c < board.CellCount; c++ {
		for r := 0; r < board.CellCount; r++ {
			Score += board.Array[c][r].Value
		}
	}
}

func (board *Board) Draw() {
	x := float32(SCREEN_WIDTH/2 - BOARD_SIZE/2)
	y := float32(SCREEN_HEIGHT/2 - BOARD_SIZE/2)

	// board_rect := rl.NewRectangle(x, y, float32(BOARD_SIZE), float32(BOARD_SIZE))

	// draw board
	rl.DrawRectangle(int32(x), int32(y), BOARD_SIZE, BOARD_SIZE, rl.Gray)

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
