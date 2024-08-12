package packages

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Window struct {
	Height int32
	Width int32
	RefreshRate int32
}

func DrawMap(width int32,height int32,offset int32) {
	rl.ClearBackground(rl.DarkGray)
	rl.DrawRectangle(0 +(offset/2),0 + (offset/2), height - offset, width - offset,rl.DarkGreen)
}