package termvel

import tl "github.com/JoelOtter/termloop"

//Menu is how you will use menus....
type Menu struct {
	background   *tl.BaseLevel
	options      []*MenuOption
	textElements []*tl.Text
	msg          GameMessage
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

//ReturnLevel returns the background of the menu (also includes all other stuff)
func (m *Menu) ReturnLevel() *tl.BaseLevel {
	return m.background
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
			msg: MsgNone,
		},
	}
	m.textElements = append(m.textElements, tl.NewText(10, 10, "Click to Select Arena Size:", tl.ColorWhite, tl.ColorBlack))
	m.options = append(m.options, &MenuOption{
		Text: tl.NewText(10, 12, "Small", tl.ColorWhite, tl.ColorBlack),
		Action: func() {
			m.msg = MsgStartMainSmall
		},
	}, &MenuOption{
		Text: tl.NewText(20, 12, "Large", tl.ColorWhite, tl.ColorBlack),
		Action: func() {
			m.msg = MsgStartMainLarge
		},
	})
	for _, t := range m.textElements {
		m.background.AddEntity(t)
	}
	for _, o := range m.options {
		m.background.AddEntity(o)
	}
	return m
}

//SetMessage satisfies the interface
func (m *Menu) SetMessage(msg GameMessage) {
	m.msg = msg
}

//ShowMainMenu displays the main menu of the game
func (m *MainMenu) ShowMainMenu(g *Game) {
	g.Screen().SetLevel(m.background)
}

//Update the menu
func (m *MainMenu) Update(g *Game) GameMessage {
	return m.msg
}
