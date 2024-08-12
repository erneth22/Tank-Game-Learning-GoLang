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
type Projectile struct{
	ProjectileColor rl.Color
	ProjectileSpeed float32
	ProjectileSize float32
}
var PTank = Tank{TankColor: rl.DarkBrown, BarrelColor: rl.Black,Size: 50, barrelSizeX: 12,barrelSizeY: 45,Speed: 5}
var PProjectile = Projectile{ProjectileColor: rl.Blue,ProjectileSpeed: 2,ProjectileSize: 24}

func lastKeyPressed() int {
	keys := []int{87, 65, 83, 68} // Keys W, A, S, D
	for _, key := range keys {
		if rl.IsKeyDown(int32(key)) {
			return key
		}
	}
	return 0
}

func DetermineDirection(dir float32) float32 {
	var moveDir = dir
    switch lastKeyPressed() {
    case 87:
        moveDir = 180 //up
    case 68:
        moveDir = 270 //right
    case 65:
        moveDir = 90  //left
    case 83:
        moveDir = 0   //south
    }
	return moveDir
}

func (PTank Tank ) DrawTank(posX int32,posY int32,rotation float32 ) {
	//tank mainbody
	rl.DrawRectanglePro(
		rl.Rectangle{X: float32(posX), Y: float32(posY), Width: float32(PTank.Size), Height: float32(PTank.Size)},
		rl.Vector2{X: float32(PTank.Size/2), Y: float32(PTank.Size/2)},
		rotation,
		PTank.TankColor)
	//tank barrel
	rl.DrawRectanglePro(
		rl.Rectangle{X:float32(posX),Y: float32(posY), Width:float32(PTank.barrelSizeX),Height:float32(PTank.barrelSizeY)},
		rl.Vector2{X:float32(PTank.barrelSizeX/2),Y:float32(0)},
		rotation,
		PTank.BarrelColor)
}

func (PProjectile Projectile ) SpawnProjectile(posX int32,posY int32,rotation float32 ) {
	
	rl.DrawRectanglePro(
		rl.Rectangle{X: float32(posX),Y: float32(posY),Width: PProjectile.ProjectileSize,Height: PProjectile.ProjectileSize},
        rl.Vector2{X: PProjectile.ProjectileSize / 2, Y: PProjectile.ProjectileSize - 66},
		rotation,
		PProjectile.ProjectileColor)
}