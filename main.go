package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 640

	large_font_size  = 36
	medium_font_size = 28
	small_font_size  = 24
)

var BOARD Board

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

	rl.DrawText("Pres [R] to restart", 10, SCREEN_HEIGHT-small_font_size*2, small_font_size, rl.DarkGray)

	if HasLost {
		game_over_text := "GAME OVER! YOU HAVE NO MORE MOVES."
		restart_instructions := "press [ENTER] to restart."

		rl.DrawText(game_over_text, SCREEN_WIDTH/2-rl.MeasureText(game_over_text, large_font_size)/2, SCREEN_HEIGHT/2-large_font_size, large_font_size, rl.Red)
		rl.DrawText(restart_instructions, SCREEN_WIDTH/2-rl.MeasureText(restart_instructions, small_font_size)/2, SCREEN_HEIGHT/2+small_font_size, small_font_size, rl.DarkGray)
	}
}

func UpdateGame() {
	if HasLost {
		if rl.IsKeyPressed(rl.KeyEnter) {
			InitGame()
		}
	} else {
		if rl.IsKeyPressed(rl.KeyR) {
			InitGame()
		}
		BOARD.Update()

		if BOARD.Motion != MotionNone {
			BOARD.MoveTiles()
		}
	}
}

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "2048")

	rl.SetTargetFPS(30)

	InitGame()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		DrawGame()
		UpdateGame()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
