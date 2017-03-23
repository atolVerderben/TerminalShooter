package termvel

import (
	"fmt"
	"math"
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

//NPC is the main struct that the user controls
type NPC struct {
	*Character
	level        *tl.BaseLevel
	update, time float64
	prevEvent    tl.Event
	Target       *Character
	evade        bool
} // ☢ ☣

//CreateNPC creates a new non-player character
func CreateNPC(x, y int, color tl.Attr, level *tl.BaseLevel) *NPC {
	npc := &NPC{
		Character: CreateCharacter(x, y, color, level),
		update:    .2,
		level:     level,
		Target:    GamePlayer.Character,
	}

	//npc.entity.SetCell(0, 0, &tl.Cell{Fg: color, Ch: '☻'})
	return npc
}

//Draw the npc at each frame
func (npc *NPC) Draw(screen *tl.Screen) {

	npc.time += screen.TimeDelta()
	//Information.text.SetText(fmt.Sprintf("Destination @ [%v, %v]", screen.TimeDelta(), npc.time))
	x, y := npc.Position()

	if npc.time > npc.update {
		npc.PrevX = x
		npc.PrevY = y
		//npc.Update()
		npc.time -= npc.update
	}
	npc.Character.Draw(screen)
}

//Update the NPC
func (npc *NPC) Update() {

	NPCInformation.text.SetText(fmt.Sprintf("Enemy Health: %v", npc.Health))
	if npc.isDead {
		return
	}
	/*if npc.hasDestination {

	} else {
		//Pick a random point

		if rand.Intn(10) == 2 {

			numX := rand.Intn(GameArenaWidth)
			numY := rand.Intn(GameArenaHeight)
			npc.DestX = numX
			npc.DestY = numY
			npc.hasDestination = true
		}

	}*/
	if npc.Target != nil {
		tx, ty := npc.Target.Position()
		x, y := npc.Position()
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
				npc.Face(Left)
			} else {
				npc.Face(Right)
			}
		} else {
			if faceY < 0 {
				npc.Face(Up)
			} else {
				npc.Face(Down)
			}
		}

		if !npc.hasDestination {
			npc.DestX = tx
			npc.DestY = ty
			npc.hasDestination = true
		}

		if math.Abs(float64(faceX)) <= 10 && math.Abs(float64(faceY)) <= 10 {
			if rand.Intn(10) == 2 {
				if npc.hasDestination {
					npc.ClearPath()
				}
				numX := rand.Intn(TermGame.Arena.arena.Width)
				numY := rand.Intn(TermGame.Arena.arena.Height)
				npc.DestX = numX
				npc.DestY = numY
				npc.hasDestination = true
			}
		}

		if rand.Intn(20) == 14 && npc.shootCoolDown == 0 {
			switch npc.Facing {
			case Up:
				TermGame.Arena.bullets.ShootBullet(npc.PrevX, npc.PrevY-2, tl.ColorRed, npc.Facing, npc.Character)
			case Down:
				TermGame.Arena.bullets.ShootBullet(npc.PrevX, npc.PrevY+2, tl.ColorRed, npc.Facing, npc.Character)
			case Left:
				TermGame.Arena.bullets.ShootBullet(npc.PrevX-2, npc.PrevY, tl.ColorRed, npc.Facing, npc.Character)
			case Right:
				TermGame.Arena.bullets.ShootBullet(npc.PrevX+2, npc.PrevY, tl.ColorRed, npc.Facing, npc.Character)
			}
			npc.shootCoolDown = 1
		}
	}
	npc.InDanger()
	npc.Character.Update()

}

//Tick processes input and reactes accordingly
func (npc *NPC) Tick(event tl.Event) {

}
