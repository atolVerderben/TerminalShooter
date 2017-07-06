package termvel

import tl "github.com/JoelOtter/termloop"

//Menu is how you will use menus....
type Menu struct {
	options      []*MenuOption
	textElements []*tl.Text
	*tl.Entity
}

//NewBasicMenu returns an empty menu with a blank entity (this is used for updates)
func NewBasicMenu() *Menu {
	m := &Menu{
		Entity: tl.NewEntity(0, 0, 0, 0),
	}
	return m
}

//MenuOption is a selectable text string
type MenuOption struct {
	*tl.Text
	Action func()
}

//NewMenuOption returns a MenuOption pointer
func NewMenuOption(x, y int, fg, bg tl.Attr, text string, action func()) *MenuOption {
	m := &MenuOption{
		Text:   tl.NewText(x, y, text, fg, bg),
		Action: action,
	}
	return m
}
