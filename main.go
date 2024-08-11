package main

import rl "github.com/gen2brain/raylib-go/raylib"
import pack "Go/packages"




func main() {
	W := pack.Window{Height: 1080, Width: 1920, RefreshRate: 60}
	rl.InitWindow(W.Width, W.Height, "Raylib Tank Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(W.RefreshRate)

	var startX int32 = W.Width/2
	var startY int32 = W.Height/2
	var direction float32 = 180.0

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGreen)

		dir := pack.DetermineDirection(direction)
		if dir == 180 {
			startY = startY-int32(pack.PTank.Speed)
		} else if dir == 270 {
			startX = startX+int32(pack.PTank.Speed)
		} else if dir == 90 {
			startX = startX-int32(pack.PTank.Speed)
		} else if dir == 0 {
			startY = startY+int32(pack.PTank.Speed)
		}
		
		direction = dir // save variable state to prevent overwriting in next frame

		pack.PTank.DrawTank(startX,startY,direction)

		rl.EndDrawing()
	}
}