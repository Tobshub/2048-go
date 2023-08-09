package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 640
)

var BOARD Board

var SCORE int

func InitGame() {
	BOARD = Board{
		CellCount: 4,
		Motion:    MotionNone,
	}
	BOARD.Init()
}

func DrawGame() {
	BOARD.Draw()
	rl.DrawText(fmt.Sprintf("Score: %d", SCORE), 10, 10, 24, rl.DarkGray)
}

func UpdateGame() {
	BOARD.Update()

	if BOARD.Motion != MotionNone {
		BOARD.MoveTiles()
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
