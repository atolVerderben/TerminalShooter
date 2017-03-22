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

//Draw the player at each frame
func (player *Player) Draw(screen *tl.Screen) {
	//screenWidth, screenHeight := screen.Size()
	player.time += screen.TimeDelta()
	x, y := player.Position()
	//offSetX, offSetY := 0, 0

	//Have a small deadzone

	//offSetX = screenWidth/2 - x

	//offSetY = screenHeight/2 - y

	//player.level.SetOffset(offSetX, offSetY)
	GameCamera.Update(screen, player.Character)
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
		if GameOverInfo == nil {
			GameOverInfo = NewEventInfo(0, 10)
			GameOverInfo.center = true
			player.level.AddEntity(GameOverInfo)
		}
		GameOverInfo.text.SetText("Game Over! You Died! Press Ctrl+C to Exit")
	}
	player.Character.Update()
}

//Tick processes input and reactes accordingly
func (player *Player) Tick(event tl.Event) {
	GameInput.Tick(event, player)
	/*player.PrevX, player.PrevY = player.Position()

	if event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Ch {
		case 'a':
			player.Face(Left)
			break
		case 's':
			player.Face(Down)
			break
		case 'd':
			player.Face(Right)
			break
		case 'w':
			player.Face(Up)
			break
		case 'q':
			GameBullets.DetonateAllBullets(player.Character)
			break
		case 'e':
			GameBullets.DetonateBullet(player.Character)
			break
		}

		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.Move(Right)
			player.ClearPath()
			break
		case tl.KeyArrowLeft:
			player.Move(Left)
			player.ClearPath()
			break
		case tl.KeyArrowUp:
			player.Move(Up)
			player.ClearPath()
			break
		case tl.KeyArrowDown:
			player.Move(Down)
			player.ClearPath()
			break
		case tl.KeyBackspace2:
			GameCamera.CenterOn(player.Character)
		case tl.KeyBackspace:
			GameCamera.CenterOn(player.Character)
		case tl.KeyTab:
			GameCamera.CenterOn(player.Character)
			break
		case tl.KeySpace:
			if player.shootCoolDown == 0 {
				switch player.Facing {
				case Up:
					GameBullets.ShootBullet(player.PrevX, player.PrevY-2, player.shotColor, player.Facing, player.Character)
				case Down:
					GameBullets.ShootBullet(player.PrevX, player.PrevY+2, player.shotColor, player.Facing, player.Character)
				case Left:
					GameBullets.ShootBullet(player.PrevX-2, player.PrevY, player.shotColor, player.Facing, player.Character)
				case Right:
					GameBullets.ShootBullet(player.PrevX+2, player.PrevY, player.shotColor, player.Facing, player.Character)
				}
				player.bulletCount++
				player.shootCoolDown = 1
			}
			break
		}
	}

	if event.Type == tl.EventMouse {

		switch event.Key {
		case tl.MouseLeft:
		//	Information.text.SetText(fmt.Sprintf("%s @ [%d, %d]", "Mouse Left", event.MouseX, event.MouseY))
		case tl.MouseMiddle:
		//	Information.text.SetText(fmt.Sprintf("%s @ [%d, %d]", "Mouse Middle", event.MouseX, event.MouseY))
		case tl.MouseRight:
		//	Information.text.SetText(fmt.Sprintf("%s @ [%d, %d]", "Mouse Right", event.MouseX, event.MouseY))
		case tl.MouseWheelUp:
		//	Information.text.SetText(fmt.Sprintf("%s @ [%d, %d]", "Mouse Wheel Up", event.MouseX, event.MouseY))
		case tl.MouseWheelDown:
		//	Information.text.SetText(fmt.Sprintf("%s @ [%d, %d]", "Mouse Wheel Down", event.MouseX, event.MouseY))
		case tl.MouseRelease:
			offX, offY := player.level.Offset()
			if Select(event.MouseX-offX, event.MouseY-offY) {
				if player.hasDestination {
					Deselect(player.DestX, player.DestY)
					player.Path = nil
					player.pathDestination = nil
				}
				player.DestX, player.DestY = event.MouseX-offX, event.MouseY-offY
				player.hasDestination = true
				//Information.text.SetText(fmt.Sprintf("Destination @ [%d, %d]", player.DestX, player.DestY))
			}

		default:

		}
		player.prevEvent = event
	}*/

}
