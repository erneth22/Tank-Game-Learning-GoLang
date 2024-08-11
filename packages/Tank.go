package packages

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tank struct{
	TankColor rl.Color
	BarrelColor rl.Color
	Size int32
	barrelSizeX int32
	barrelSizeY int32
	Speed uint8
}

var PTank = Tank{TankColor: rl.Red, BarrelColor: rl.DarkPurple, Size: 50, barrelSizeX: 12,barrelSizeY: 45,Speed: 5}

func lastKeyPressed() int {
	keys := []int{87, 65, 83, 68} // Keys W, A, S, D
	for _, key := range keys {
		if rl.IsKeyDown(int32(key)) {
			return key
		}
	}
	return 0
}
func (PTank Tank ) DrawTank(posX int32,posY int32 ) {
	var moveDir float32
    switch lastKeyPressed(){
	case 87:
		moveDir = 180
	case 68:
		moveDir = 270
	case 65:
		moveDir = 90
	case 83:
		moveDir = 0
	}
	
	
	rl.DrawRectangle(posX, posY, PTank.Size, PTank.Size, PTank.TankColor) //tank mainbody
	rl.DrawRectanglePro(
		rl.Rectangle{float32(posX+(PTank.Size/2)),float32(posY+(PTank.Size/2)),float32(PTank.barrelSizeX), float32(PTank.barrelSizeY)},
		rl.Vector2{float32(PTank.barrelSizeX/2),float32(0)},
	moveDir,
		PTank.BarrelColor)
}