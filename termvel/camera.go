package termvel

import tl "github.com/JoelOtter/termloop"

//Camera represents the game camera
type Camera struct {
	level                     *tl.BaseLevel
	X, Y, PrevX, PrevY        int
	DeadZoneX, DeadZoneY      int
	screenWidth, screenHeight int
	size                      int
}

//CreateCamera returns a new Camera struct
func CreateCamera(x, y, deadx, deady int, level *tl.BaseLevel, size int) *Camera {
	c := &Camera{
		X:         x,
		Y:         y,
		PrevX:     x,
		PrevY:     y,
		DeadZoneX: deadx,
		DeadZoneY: deady,
		level:     level,
		size:      size,
	}

	return c
}

//CenterOn on the passed Character
func (c *Camera) CenterOn(char *Character) {

	x, y := char.Position()
	c.Y = c.screenHeight/2 - y
	c.PrevY = y
	c.X = c.screenWidth/2 - x
	c.PrevX = x
	c.level.SetOffset(c.X, c.Y)
}

//Update just allows for adding a leveloffset to the game and simulating a camera
//This was added because without a deadzone I found the "camera" movement gave me a headache
func (c *Camera) Update(screen *tl.Screen, char *Character) {
	c.screenWidth, c.screenHeight = screen.Size()
	x, y := char.Position()
	if c.size == 0 {
		c.level.SetOffset(c.screenWidth/2-50, c.screenHeight/2-25)
		return
	}

	if c.PrevY+c.DeadZoneY < y || c.PrevY-c.DeadZoneY > y {
		c.Y = c.screenHeight/2 - y
		c.PrevY = y
	}

	if c.PrevX+c.DeadZoneX < x || c.PrevX-c.DeadZoneX > x {
		c.X = c.screenWidth/2 - x
		c.PrevX = x
	}
	c.level.SetOffset(c.X, c.Y)
}
