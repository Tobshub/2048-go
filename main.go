package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "2048")

	rl.SetWindowMinSize(500, 480)
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
