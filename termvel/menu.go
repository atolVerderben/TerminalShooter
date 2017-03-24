package termvel

import tl "github.com/JoelOtter/termloop"

//Menu is how you will use menus....
type Menu struct {
	background   *tl.BaseLevel
	options      []*MenuOption
	textElements []*tl.Text
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

//MainMenu is the beginning of the game
type MainMenu struct {
	*Menu
}

//CreateMainMenu returns a pointer for the MainMenu Game State
func CreateMainMenu() *MainMenu {
	m := &MainMenu{
		Menu: &Menu{
			background: tl.NewBaseLevel(tl.Cell{
				Bg: tl.ColorBlack,
				Fg: tl.ColorWhite,
				Ch: ' ',
			}),
		},
	}
	m.textElements = append(m.textElements, tl.NewText(10, 10, "Testing 1 2 3", tl.ColorWhite, tl.ColorBlack))
	for _, t := range m.textElements {
		m.background.AddEntity(t)
	}
	return m
}

//Update the menu
func (m *MainMenu) Update(g *Game) GameMessage {
	return MsgNone
}
