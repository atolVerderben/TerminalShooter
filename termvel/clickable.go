package termvel

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

//Clickable is a structure that is selectable with the mouse -- probably will change to interface later
type Clickable struct {
	entity             *tl.Entity
	level              *tl.BaseLevel
	selected, hasGrass bool
}

//NewClickable creates a clickable object
func NewClickable(x, y, w, h int, col tl.Attr, level *tl.BaseLevel) *Clickable {
	click := &Clickable{
		entity: tl.NewEntity(x, y, w, h),
		level:  level,
	}

	num := rand.Intn(50)
	if num == 13 {
		click.entity.SetCell(0, 0, &tl.Cell{Fg: col, Bg: tl.ColorBlack, Ch: '\''})
		click.hasGrass = true
	} else {
		click.entity.SetCell(0, 0, &tl.Cell{Fg: col, Bg: tl.ColorBlack, Ch: ' '})

	}
	//fmt.Println(num)

	return click
}

//Draw clickable object
func (c *Clickable) Draw(s *tl.Screen) {

	c.entity.Draw(s)

}

//Deselect the current object
func (c *Clickable) Deselect() {
	c.selected = false
	if c.hasGrass {
		c.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlack, Bg: tl.ColorBlack, Ch: '\''})
	} else {
		c.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Bg: tl.ColorBlack, Ch: ' '})
	}
}

//Select the current object
func (c *Clickable) Select() {
	c.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlue, Bg: tl.ColorBlack, Ch: '•'})
	c.selected = true
}

//Tick the next action
func (c *Clickable) Tick(ev tl.Event) {
	/*x, y := c.entity.Position()
	offX, offy := c.level.Offset()
	if ev.Type == tl.EventMouse && ev.MouseX == x+offX && ev.MouseY == y+offy {
		//if ev.Key == tl.MouseRelease {
		//	c.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Bg: tl.ColorGreen, Ch: '•'})
		//	c.selected = true
		//}
	} else if ev.Type == tl.EventMouse {
		c.Deselect()
	}*/
}
