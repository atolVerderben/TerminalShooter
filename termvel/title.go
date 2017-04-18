package termvel

import tl "github.com/JoelOtter/termloop"

//TitleMain is how you will use menus....
type TitleMain struct {
	*Menu
	msg GameMessage
}

//CreateTitle returns a pointer for the MainMenu Game State
func CreateTitle() *TitleMain {
	m := &TitleMain{
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
			m.msg = MsgMainMenu
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
