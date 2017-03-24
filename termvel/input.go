package termvel

import (
	"math"

	tl "github.com/JoelOtter/termloop"
)

//Input represents the player input
type Input struct {
	Keys      map[string]tl.Key
	scheme    int
	gameState GameState
	*tl.Entity
}

//Different input types
const (
	InputA int = iota
	InputB
	InputC
)

//NewInput returns the default Input struct
func NewInput() *Input {
	i := &Input{
		Keys:   map[string]tl.Key{},
		Entity: tl.NewEntity(0, 0, 0, 0),
	}

	i.Keys["Up"] = tl.KeyCtrlW
	i.Keys["Down"] = tl.KeyCtrlA
	i.Keys["Left"] = tl.KeyArrowLeft
	i.Keys["Right"] = tl.KeyArrowRight

	i.Keys["AimUp"] = tl.KeyArrowUp
	i.Keys["AimDown"] = tl.KeyArrowDown
	i.Keys["AimLeft"] = tl.KeyArrowLeft
	i.Keys["AimRight"] = tl.KeyArrowRight
	i.scheme = 1
	TermGame.Screen().AddEntity(i)
	return i
}

//Tick process when key pressed
func (input *Input) Tick(event tl.Event) {
	input.Update(event, TermGame.player)
}

//Update fires an Input Tick Event
func (input *Input) Update(event tl.Event, player *Player) {
	player.PrevX, player.PrevY = player.Position()
	switch gs := input.gameState.(type) {
	case *Arena:
		input.UpdateGamePlay(event, player, gs)
		break
	}

}

//UpdateGamePlay is used for the actual game
func (input *Input) UpdateGamePlay(event tl.Event, player *Player, gs *Arena) {
	if event.Type == tl.EventKey {
		switch event.Ch {
		case '1':
			input.scheme = InputA
		case '2':
			input.scheme = InputB
		}

		switch event.Key {
		case tl.KeyTab:
			//GameCamera.CenterOn(player.Character)
			switch input.scheme {
			case InputA:
				input.scheme = InputB
				break
			case InputB:
				input.scheme = InputA
				break
			}
			break
		}

	}
	switch input.scheme {
	case InputA: //WASD Aims Arrows Move
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
				TermGame.Arena.bullets.DetonateAllBullets(player.Character)
				break
			case 'e':
				TermGame.Arena.bullets.DetonateBullet(player.Character)
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
				gs.camera.CenterOn(player.Character)
			case tl.KeyBackspace:
				gs.camera.CenterOn(player.Character)
			case tl.KeySpace:
				if player.shootCoolDown == 0 {
					switch player.Facing {
					case Up:
						TermGame.Arena.bullets.ShootBullet(player.PrevX, player.PrevY-2, player.shotColor, player.Facing, player.Character)
					case Down:
						TermGame.Arena.bullets.ShootBullet(player.PrevX, player.PrevY+2, player.shotColor, player.Facing, player.Character)
					case Left:
						TermGame.Arena.bullets.ShootBullet(player.PrevX-2, player.PrevY, player.shotColor, player.Facing, player.Character)
					case Right:
						TermGame.Arena.bullets.ShootBullet(player.PrevX+2, player.PrevY, player.shotColor, player.Facing, player.Character)
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
		}
		break
	case InputB: // WASD Move Arrows Aim
		if event.Type == tl.EventKey { // Is it a keyboard event?
			switch event.Ch {
			case 'a':

				player.Move(Left)
				player.ClearPath()
				break
			case 's':

				player.Move(Down)
				player.ClearPath()
				break
			case 'd':
				player.Move(Right)
				player.ClearPath()
				break
			case 'w':
				player.Move(Up)
				player.ClearPath()
				break
			case 'q':
				TermGame.Arena.bullets.DetonateAllBullets(player.Character)
				break
			case 'e':
				TermGame.Arena.bullets.DetonateBullet(player.Character)
				break
			}

			switch event.Key { // If so, switch on the pressed key.
			case tl.KeyArrowRight:
				player.Face(Right)

				break
			case tl.KeyArrowLeft:
				player.Face(Left)
				break
			case tl.KeyArrowUp:
				player.Face(Up)

				break
			case tl.KeyArrowDown:
				player.Face(Down)
				break
			case tl.KeyBackspace2:
				gs.camera.CenterOn(player.Character)
			case tl.KeyBackspace:
				gs.camera.CenterOn(player.Character)

			case tl.KeySpace:
				if player.shootCoolDown == 0 {
					switch player.Facing {
					case Up:
						TermGame.Arena.bullets.ShootBullet(player.PrevX, player.PrevY-2, player.shotColor, player.Facing, player.Character)
					case Down:
						TermGame.Arena.bullets.ShootBullet(player.PrevX, player.PrevY+2, player.shotColor, player.Facing, player.Character)
					case Left:
						TermGame.Arena.bullets.ShootBullet(player.PrevX-2, player.PrevY, player.shotColor, player.Facing, player.Character)
					case Right:
						TermGame.Arena.bullets.ShootBullet(player.PrevX+2, player.PrevY, player.shotColor, player.Facing, player.Character)
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
				offX, offY := player.level.Offset()

				tx, ty := event.MouseX-offX, event.MouseY-offY //player.Target.Position()
				x, y := player.Position()
				faceX, faceY := 0, 0
				if tx < x {
					faceX = x - tx
					faceX *= -1
				} else if tx > x {
					faceX = tx - x
				}

				if ty < y {
					faceY = y - ty
					faceY *= -1
				} else if ty > y {
					faceY = ty - y
				}

				if math.Abs(float64(faceX))/2 > math.Abs(float64(faceY)) {
					if faceX < 0 {
						player.Face(Left)
					} else {
						player.Face(Right)
					}
				} else {
					if faceY < 0 {
						player.Face(Up)
					} else {
						player.Face(Down)
					}
				}

				if player.shootCoolDown == 0 {
					switch player.Facing {
					case Up:
						TermGame.Arena.bullets.ShootBullet(player.PrevX, player.PrevY-2, player.shotColor, player.Facing, player.Character)
					case Down:
						TermGame.Arena.bullets.ShootBullet(player.PrevX, player.PrevY+2, player.shotColor, player.Facing, player.Character)
					case Left:
						TermGame.Arena.bullets.ShootBullet(player.PrevX-2, player.PrevY, player.shotColor, player.Facing, player.Character)
					case Right:
						TermGame.Arena.bullets.ShootBullet(player.PrevX+2, player.PrevY, player.shotColor, player.Facing, player.Character)
					}
					player.bulletCount++
					player.shootCoolDown = 1
				}
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

			default:

			}
			player.prevEvent = event
		}
		break
	}
}
