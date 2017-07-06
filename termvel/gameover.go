package termvel

import (
	"os"

	tl "github.com/JoelOtter/termloop"
)

//GameOver is how you will use menus....
type GameOver struct {
	*Menu
	background *tl.BaseLevel
	msg        GameMessage
}

//CreateGameOver returns a pointer for the MainMenu Game State
func CreateGameOver(win bool) *GameOver {
	m := &GameOver{
		background: tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorBlack,
			Fg: tl.ColorWhite,
			Ch: ' ',
		}),
		Menu: NewBasicMenu(),
		msg:  MsgNone,
	}
	if win {
		m.textElements = append(m.textElements, tl.NewText(10, 10, "Congratulations! Your Foes are Defeated!", tl.ColorWhite, tl.ColorBlack))
	} else {
		m.textElements = append(m.textElements, tl.NewText(10, 10, "Game Over. Try Again?", tl.ColorWhite, tl.ColorBlack))
	}

	m.options = append(m.options, &MenuOption{
		Text: tl.NewText(10, 12, "Restart", tl.ColorWhite, tl.ColorBlack),
		Action: func() {
			m.msg = MsgMainMenu
		},
	},
		&MenuOption{
			Text: tl.NewText(10, 14, "Exit", tl.ColorWhite, tl.ColorBlack),
			Action: func() {
				os.Exit(0)
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
func (m *GameOver) SetMessage(msg GameMessage) {
	m.msg = msg
}

//ShowGameOver displays the main menu of the game
func (m *GameOver) ShowGameOver(g *Game) {
	g.Screen().SetLevel(m.background)
}

//Update the menu
func (m *GameOver) Update(g *Game) GameMessage {
	return m.msg
}

//Tick processes input and reactes accordingly
func (m *GameOver) Tick(event tl.Event) {
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

//ReturnLevel returns the background
func (m *GameOver) ReturnLevel() *tl.BaseLevel {
	return m.background
}
