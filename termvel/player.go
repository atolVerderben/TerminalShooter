package termvel

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

//Player is the main struct that the user controls
type Player struct {
	*Character
	Name         string
	level        *tl.BaseLevel
	prevEvent    tl.Event
	time, update float64
	bulletCount  int
	shotColor    tl.Attr
}

//Fun unicode for characters
//☺☻옷

//CreatePlayer returns a pointer to a Player struct
func CreatePlayer(x, y int, color tl.Attr, level *tl.BaseLevel) *Player {
	player := &Player{
		Character: CreateCharacter(x, y, color, level),
		level:     level,
		update:    .0167,
		Name:      "Player1",
		shotColor: tl.ColorCyan,
	}
	return player
}

//SetLevel for the player
func (player *Player) SetLevel(level *tl.BaseLevel) {
	player.level = level
}

//Draw the player at each frame
func (player *Player) Draw(screen *tl.Screen) {
	player.time += screen.TimeDelta()
	x, y := player.Position()
	//GameCamera.Update(screen, player.Character)
	if player.time > player.update {
		player.PrevX = x
		player.PrevY = y
		//player.Update()
		player.time = 0 //-= player.update
	}
	//Information.text.SetText(fmt.Sprintf("Destination @ [%v, %v]", screen.TimeDelta(), player.time))
	player.Character.Draw(screen)
}

//Direction represents a direction the character can move or face
type Direction string

//Four possible directions
const (
	Up    Direction = "u"
	Down  Direction = "d"
	Left  Direction = "l"
	Right Direction = "r"
)

//Update the player
func (player *Player) Update() {
	Information.text.SetText(fmt.Sprintf("Player Health: %v", player.Health))
	if player.isDead {
		/*if GameOverInfo == nil {
			GameOverInfo = NewEventInfo(0, 10)
			GameOverInfo.center = true
			player.level.AddEntity(GameOverInfo)
		}
		GameOverInfo.text.SetText("Game Over! You Died! Press Ctrl+C to Exit or Ctrl+Q to Restart")*/

	}
	player.Character.Update()
}

//Tick processes input and reactes accordingly
//func (player *Player) Tick(event tl.Event) {

//}
