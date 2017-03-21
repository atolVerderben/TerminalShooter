package termvel

import (
	"math"
	"math/rand"
	"runtime"

	tl "github.com/JoelOtter/termloop"
	astar "github.com/beefsack/go-astar"
)

//Character is the representation of the basic character "model"
type Character struct {
	*tl.Entity
	reticle                         *tl.Entity
	Facing                          Direction
	PrevX, PrevY                    int
	hasDestination                  bool
	DestX, DestY                    int
	Speed                           int
	moveCooldown                    int
	Health                          int
	shootCoolDown, shootCoolDownMax int
	isDead                          bool
	color                           tl.Attr
	Path                            []*Point
	pathDestination                 *Point
	//This is a placeholder to just get this working for now... not very efficient
	explosionHitBy *Explosion
}

//CreateCharacter creates a new character of course
func CreateCharacter(x, y int, color tl.Attr, level *tl.BaseLevel) *Character {
	char := &Character{
		Entity:           tl.NewEntity(x, y, 1, 1),
		reticle:          tl.NewEntity(x+1, y, 1, 1),
		Facing:           Right,
		Speed:            5,
		Health:           6,
		shootCoolDownMax: 10,
		color:            color,
	}
	// Set the character at position (0, 0) on the entity. 😃 ☠
	char.reticle.SetCell(0, 0, &tl.Cell{Fg: color, Ch: '•'})
	if runtime.GOOS == "windows" {
		char.SetCell(0, 0, &tl.Cell{Fg: color, Ch: '☻'})
	} else {
		char.SetCell(0, 0, &tl.Cell{Fg: color, Ch: '옷'}) // that unicode isn't usually present on windows
	}
	return char
}

//Move in the specified direction
func (char *Character) Move(dir Direction) {
	if char.moveCooldown == 0 {
		x, y := char.Position()
		switch dir {
		case Up:
			char.SetPosition(x, y-1)
		case Down:
			char.SetPosition(x, y+1)
		case Left:
			char.SetPosition(x-1, y)
		case Right:
			char.SetPosition(x+1, y)
		}
		char.moveCooldown = 1
	}
}

//Face specified direction
func (char *Character) Face(dir Direction) {
	switch dir {
	case Up:
		char.reticle.SetPosition(char.PrevX, char.PrevY-2)
	case Down:
		char.reticle.SetPosition(char.PrevX, char.PrevY+2)
	case Left:
		char.reticle.SetPosition(char.PrevX-2, char.PrevY)
	case Right:
		char.reticle.SetPosition(char.PrevX+2, char.PrevY)
	}
	char.Facing = dir
}

//Draw the char at each frame
func (char *Character) Draw(screen *tl.Screen) {

	//char.time += screen.TimeDelta()
	//Information.text.SetText(fmt.Sprintf("Destination @ [%v, %v]", screen.TimeDelta(), char.time))
	//x, y := char.Position()
	/*if char.time > char.update {
		char.PrevX = x
		char.PrevY = y
		char.Update()
		char.time -= char.update
	}*/
	char.Entity.Draw(screen)
	if !char.isDead {
		char.reticle.Draw(screen)
	}
}

//Update char at ~60fps
func (char *Character) Update() {
	if char.Path != nil {
		//	Information.text.SetText("I HAVE PATH!")
	} else {
		//	Information.text.SetText("No Path!")
	}
	if char.isDead {
		return
	}
	//TODO: not hardcode max health
	if char.Health > 6 {
		char.Health = 6
	}
	if char.shootCoolDown > 0 {
		if char.shootCoolDown > char.shootCoolDownMax {
			char.shootCoolDown = 0
		} else {
			char.shootCoolDown++
		}
	}
	if char.moveCooldown > 0 {
		if char.moveCooldown > char.Speed {
			char.moveCooldown = 0
		} else {
			char.moveCooldown++
		}
		char.Face(char.Facing)
		return
	}
	char.PrevX, char.PrevY = char.Position()
	if char.hasDestination {
		if char.Path == nil && char.pathDestination == nil {
			x, y := char.Position()
			if char.DestX <= 0 || char.DestX >= GameArenaWidth || char.DestY <= 0 || char.DestY >= GameArenaHeight {
				char.hasDestination = false
				return
			}
			originalTile := GameWorld.Grid.PathTile(x, y).Kind

			GameWorld.Grid.SetPathTile(&PathTile{
				Kind: KindFrom,
			}, x, y)

			toTile, fromTile := GameWorld.Grid.PathTile(char.DestX, char.DestY), GameWorld.Grid.PathTile(x, y)
			if toTile == nil {
				char.hasDestination = false
				return
			}
			if toTile.Kind == KindBlocker {
				//char.hasDestination = false
				//return
			}
			p, _, found := astar.Path(toTile, fromTile)
			if found {
				//Information.text.SetText("Found")

				if len(p) > 1 {
					//Information.text.SetText("I HAVE PATH!")
					for x := 1; x < len(p); x++ {
						pt := p[x].(*PathTile)
						p := &Point{
							X: pt.X,
							Y: pt.Y,
						}

						char.Path = append(char.Path, p)

					}
					char.pathDestination = Shift(&char.Path)
				}
			}
			//Set back the original tilekind
			GameWorld.Grid.SetPathTile(&PathTile{
				Kind: originalTile,
			}, x, y)

		}

		if char.pathDestination != nil {
			if char.pathDestination.X == char.PrevX && char.pathDestination.Y == char.PrevY {
				if char.Path != nil && len(char.Path) >= 1 {
					char.pathDestination = Shift(&char.Path)
				}
				char.InDanger()
			} else { //actually move
				if char.pathDestination.X < char.PrevX {
					char.Move(Left)
				} else if char.pathDestination.X > char.PrevX {
					char.Move(Right)
				} else if char.pathDestination.Y < char.PrevY {
					char.Move(Up)
				} else if char.pathDestination.Y > char.PrevY {
					char.Move(Down)
				}

			}
		}

		/*// I'll put in some astar here later
		if char.DestX < char.PrevX {
			char.Move(Left)
			//char.Face(Left)
		} else if char.DestX > char.PrevX {
			char.Move(Right)
			//char.Face(Right)
		} else if char.DestY < char.PrevY {
			char.Move(Up)
			//char.Face(Up)
		} else if char.DestY > char.PrevY {
			char.Move(Down)
			//char.Face(Down)
		}*/

		if char.DestY == char.PrevY && char.DestX == char.PrevX {
			char.hasDestination = false
			char.Path = nil
			char.pathDestination = nil
			Deselect(char.DestX, char.DestY)
		}

	}
	char.Face(char.Facing)
	//char.InDanger()
}

//Die runs when the character's health drops below 0
func (char *Character) Die() {
	char.Health = 0
	char.isDead = true
	//☺☻☠
	char.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '☺'})
}

//Collide function to make Character DynamicPhysical and handle collision
func (char *Character) Collide(collision tl.Physical) {

	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		char.SetPosition(char.PrevX, char.PrevY)
		//Temp Fix to solve this issue
		char.hasDestination = false
		char.Path = nil
		char.pathDestination = nil
		Deselect(char.DestX, char.DestY) //This should really only be for the player
	}

	if _, ok := collision.(*Water); ok {
		char.SetPosition(char.PrevX, char.PrevY)
		//Temp Fix to solve this issue
		char.hasDestination = false
		char.Path = nil
		char.pathDestination = nil
		Deselect(char.DestX, char.DestY) //This should really only be for the player
	}

	if _, ok := collision.(*Bullet); ok {
		char.Health -= 2
		if char.Health <= 0 {
			char.Die()
		}
	}

	if e, ok := collision.(*Explosion); ok {
		if char.explosionHitBy != nil && char.explosionHitBy == e {
			return
		}
		char.explosionHitBy = e
		char.Health--
		if char.Health <= 0 {
			char.Die()
		}
	}
}

//ClearPath removes all points and destination from the character
func (char *Character) ClearPath() {
	char.hasDestination = false
	char.Path = nil
	char.pathDestination = nil
	Deselect(char.DestX, char.DestY) //This should really only be for the player
}

//InDanger checks if the character is in danger of being hit by a bullets and react.
//TODO: include explosions
func (char *Character) InDanger() {
	for _, b := range GameBullets.bullets {
		if b.owner == char {
			continue
		}
		tx, ty := b.Position()
		x, y := char.Position()
		faceX, faceY := 0, 0
		if tx < x { // To the LEFT
			faceX = x - tx
			faceX *= -1
		} else if tx > x { // To the RIGHT
			faceX = tx - x
		}

		if ty < y { // Above (UP)
			faceY = y - ty
			faceY *= -1
		} else if ty > y { // Below (DOWN)
			faceY = ty - y
		}

		//If it's kind of close let's check the trajectory
		if math.Abs(float64(faceX)) <= 5 && math.Abs(float64(faceY)) <= 5 {

			if ty == y {
				if faceX < 0 && b.Velocity == Right { // To the left and moving right
					if rand.Intn(2) == 1 {
						char.ClearPath()
						char.Move(Up)
						//char.hasDestination = true
					} else {
						char.ClearPath()
						char.Move(Down)
						//char.hasDestination = true
					}
					//NPCInformation.text.SetText("Coming in HOT!")
				} else if faceX > 0 && b.Velocity == Left { // To the right and moving left
					if rand.Intn(2) == 1 {
						char.ClearPath()
						char.Move(Up)
						//char.hasDestination = true
					} else {
						char.ClearPath()
						char.Move(Down)
						//char.hasDestination = true
					}
					//NPCInformation.text.SetText("Coming in HOT!")
				}
			}

			if tx == x {
				if faceY < 0 && b.Velocity == Down { // Above and moving down
					if rand.Intn(2) == 1 {
						char.ClearPath()
						char.Move(Left)
						char.Move(Left)
						//char.hasDestination = true
					} else {
						char.ClearPath()
						char.Move(Right)
						char.Move(Right)
						//char.hasDestination = true
					}
					//NPCInformation.text.SetText("Coming in HOT!")
				} else if faceY > 0 && b.Velocity == Up { // Below and moving Up
					if rand.Intn(2) == 1 {
						char.ClearPath()
						char.Move(Left)
						char.Move(Left)
						//char.hasDestination = true
					} else {
						char.ClearPath()
						char.Move(Right)
						char.Move(Right)
						//char.hasDestination = true
					}
					//NPCInformation.text.SetText("Coming in HOT!")
				}
			}
		} else {
			//NPCInformation.text.SetText("SAFE!")
		}

		/*if math.Abs(float64(faceX))/2 > math.Abs(float64(faceY)) {
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
		}*/
	}
}
