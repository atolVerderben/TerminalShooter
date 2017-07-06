package termvel

import tl "github.com/JoelOtter/termloop"

//MainMenu is the beginning of the game
type MainMenu struct {
	*Menu
	background *tl.BaseLevel
	msg        GameMessage
}

//CreateMainMenu returns a pointer for the MainMenu Game State
func CreateMainMenu() *MainMenu {
	m := &MainMenu{
		background: tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorBlack,
			Fg: tl.ColorWhite,
			Ch: ' ',
		}),
		Menu: NewBasicMenu(),
		msg:  MsgNone,
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
	m.background.AddEntity(m)
	return m
}

//SetMessage satisfies the interface
func (m *MainMenu) SetMessage(msg GameMessage) {
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

//ReturnLevel returns the background of the menu (also includes all other stuff)
func (m *MainMenu) ReturnLevel() *tl.BaseLevel {
	return m.background
}

//Tick processes input and reactes accordingly
func (m *MainMenu) Tick(event tl.Event) {
	if event.Type == tl.EventMouse {
		switch event.Key {
		case tl.MouseRelease:
			for _, option := range m.options {
				mx, my := event.MouseX, event.MouseY
				x, y := option.Position()
				w, h := option.Size()
				if mx >= x && mx <= x+w {
					if my >= y && my <= y+h {
						option.Action()
					}
				}
			}
		}
	}
}
