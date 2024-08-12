package main

import rl "github.com/gen2brain/raylib-go/raylib"
import pack "Go/packages"




func main() {
	W := pack.Window{Height: 1080, Width: 1920, RefreshRate: 60}
	rl.InitWindow(W.Width, W.Height, "Raylib Tank Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(W.RefreshRate)

	var posX = W.Width/2
	var posY = W.Height/2
	var direction float32 = 180.0

	var offset = pack.PTank.Size+20


	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		pack.DrawMap(W.Height,W.Width,offset)

		dir := pack.DetermineDirection(direction)
		if dir == 180 && posY > offset {
	        posY = posY-int32(pack.PTank.Speed)
		} else if dir == 270 && posX < (W.Width - offset) {
			posX = posX+int32(pack.PTank.Speed)
		} else if dir == 90 && posX > offset {
			posX = posX-int32(pack.PTank.Speed)
		} else if dir == 0 && posY < (W.Height - offset) {
			posY = posY+int32(pack.PTank.Speed)
		}
		
		direction = dir // save variable state to prevent overwriting in next frame

		pack.PTank.DrawTank(posX,posY,direction)
		pack.PProjectile.SpawnProjectile(posX,posY,direction)

		rl.EndDrawing()
	}
}