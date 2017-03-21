package termvel

import tl "github.com/JoelOtter/termloop"

//Water is a simple placeholder for "water"
//The players collide with it but bullets pass through in other words
type Water struct {
	*tl.Rectangle
}

//NewWater returns basically a blue rectangle
func NewWater(x, y, w, h int) *Water {
	water := &Water{
		Rectangle: tl.NewRectangle(x, y, w, h, tl.ColorYellow),
	}
	return water
}
