package main

import rl "github.com/gen2brain/raylib-go/raylib"
import pack "Go/packages"




func main() {
	W := pack.Window{Height: 1080, Width: 1920, RefreshRate: 60}
	rl.InitWindow(W.Width, W.Height, "Raylib Tank Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(W.RefreshRate)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		
		pack.PTank.DrawTank(900,500)

		rl.EndDrawing()
	}
}