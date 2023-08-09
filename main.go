package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 640
)

var BOARD Board

func InitGame() {
	BOARD = Board{
		CellCount: 4,
	}
	BOARD.Init()
}

func DrawGame() {
	BOARD.Draw()
}

func UpdateGame() {
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
