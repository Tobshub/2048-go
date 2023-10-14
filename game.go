package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	SCREEN_WIDTH  int32 = 800
	SCREEN_HEIGHT int32 = 680
)

const (
	large_font_size  = 36
	medium_font_size = 28
	small_font_size  = 24
)

var (
	GameSettings Settings
	BOARD        Board
)

var (
	HiScore int = 0
	Score   int = 0

	HasLost bool = false
)

func InitGame() {
	Score = 0
	HasLost = false
	BOARD = Board{
		CellCount: 4,
		Motion:    MotionNone,
	}
	BOARD.Init()
}

func DrawGame() {
	BOARD.Draw()
	rl.DrawText(fmt.Sprintf("Hi-Score: %d", HiScore), 10, 10, medium_font_size, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("Score: %d", Score), 10, 10+medium_font_size+5, medium_font_size, rl.DarkGray)

	rl.DrawText("Restart: [R]", 10, SCREEN_HEIGHT-small_font_size*4, small_font_size, rl.DarkGray)
	rl.DrawText("Settings: [S]", 10, SCREEN_HEIGHT-small_font_size*2, small_font_size, rl.DarkGray)

	if HasLost {
		game_over_text := "GAME OVER! YOU HAVE NO MORE MOVES."
		restart_instructions := "press [R] to restart."

		rl.DrawText(
			game_over_text,
			SCREEN_WIDTH/2-rl.MeasureText(game_over_text, large_font_size)/2,
			SCREEN_HEIGHT/2-large_font_size,
			large_font_size,
			rl.Red,
		)
		rl.DrawText(
			restart_instructions,
			SCREEN_WIDTH/2-rl.MeasureText(restart_instructions, small_font_size)/2,
			SCREEN_HEIGHT/2+small_font_size,
			small_font_size,
			rl.DarkGray,
		)
	}
}

func UpdateGame() {
	if rl.IsKeyPressed(rl.KeyR) {
		InitGame()
	} else if rl.IsKeyPressed(rl.KeyU) {
		BOARD.UndoState()
		if HasLost {
			HasLost = false
		}
	} else if rl.IsKeyPressed(rl.KeyS) {
		ToggleSettings(&GameSettings)
	}

	if !HasLost {
		BOARD.Update()

		if BOARD.Motion != MotionNone {
			BOARD.MoveTiles()
		}
	}

	if rl.IsWindowResized() {
		SCREEN_WIDTH = int32(rl.GetScreenWidth())
		SCREEN_HEIGHT = int32(rl.GetScreenHeight())
		BOARD_SIZE = int32(math.Min(float64(SCREEN_WIDTH), float64(SCREEN_HEIGHT))) - padding
	}
}
