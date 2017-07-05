package termvel

import tl "github.com/JoelOtter/termloop"

//TitleMain is how you will use menus....
type TitleMain struct {
	*Menu
	*tl.Entity
	msg GameMessage
}

//CreateTitle returns a pointer for the MainMenu Game State
func CreateTitle(g *Game) *TitleMain {
	m := &TitleMain{
		Entity: tl.NewEntity(0, 0, 0, 0),
		Menu: &Menu{
			background: tl.NewBaseLevel(tl.Cell{
				Bg: tl.ColorBlack,
				Fg: tl.ColorWhite,
				Ch: ' ',
			}),
		},
		msg: MsgNone,
	}
	m.textElements = append(m.textElements, tl.NewText(10, 10, "Terminal Arena Shooter:", tl.ColorWhite, tl.ColorBlack))
	m.options = append(m.options, &MenuOption{
		Text: tl.NewText(10, 12, "Start", tl.ColorWhite, tl.ColorBlack),
		Action: func() {
			g.Screen().RemoveEntity(m)
			m.msg = MsgMainMenu
		},
	})
	for _, t := range m.textElements {
		m.background.AddEntity(t)
	}
	for _, o := range m.options {
		m.background.AddEntity(o)
	}
	g.Screen().AddEntity(m)
	return m
}

//SetMessage satisfies the interface
func (m *TitleMain) SetMessage(msg GameMessage) {
	m.msg = msg
}

//ShowTitleMain displays the main menu of the game
func (m *TitleMain) ShowTitleMain(g *Game) {
	g.Screen().SetLevel(m.background)
}

//Update the menu
func (m *TitleMain) Update(g *Game) GameMessage {
	return m.msg
}

func (m *TitleMain) Tick(event tl.Event) {
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
