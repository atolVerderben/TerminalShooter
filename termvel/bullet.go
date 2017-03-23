package termvel

import tl "github.com/JoelOtter/termloop"

//BulletController controls the bullets I suppose...
type BulletController struct {
	bullets []*Bullet
	level   *tl.BaseLevel
}

//CreateBulletController creates a new empty BulletController
func CreateBulletController(level *tl.BaseLevel) *BulletController {
	bc := &BulletController{
		bullets: []*Bullet{},
		level:   level,
	}
	return bc
}

//AddBullet to the BulletController
func (bc *BulletController) AddBullet(bullet *Bullet) {
	bc.bullets = append(bc.bullets, bullet)
}

//RemoveBullet removes the bullet object from the slice
func (bc *BulletController) RemoveBullet(bullet *Bullet) {
	//fmt.Println("Attempting to remove Bullet")
	delete := -1
	for index, e := range bc.bullets {
		if e == bullet {
			delete = index
			break
		}
	}
	if delete >= 0 {
		bc.level.RemoveEntity(bullet)
		bc.bullets = append(bc.bullets[:delete], bc.bullets[delete+1:]...)
	}

}

//Update all active bullets
func (bc *BulletController) Update() {
	for _, b := range bc.bullets {
		//	bc.bullets[i].Update()
		b.Update()
	}
}

//Bullet is something you shoot at someone or something(?!) else
type Bullet struct {
	*tl.Entity
	Velocity   Direction
	level      *tl.BaseLevel
	Speed      int
	speedCount int
	Type       string
	removeFlag bool
	ID         string
	owner      *Character
}

//ShootBullet creates a new bullet and adds it to the controller
func (bc *BulletController) ShootBullet(x, y int, color tl.Attr, vel Direction, owner *Character) {
	b := &Bullet{
		Entity:   tl.NewEntity(x, y, 1, 1),
		Velocity: vel,
		Speed:    2,
		Type:     "Bullet",
		owner:    owner,
	}
	b.SetCell(0, 0, &tl.Cell{Fg: color, Ch: '*'})
	bc.level.AddEntity(b)
	bc.AddBullet(b)
}

//DetonateBullet detonates the first bullet fired by the passed Character
func (bc *BulletController) DetonateBullet(char *Character) {
	for _, b := range bc.bullets {
		if b.owner == char {
			x, y := b.Position()
			TermGame.Arena.explosions.CreateExplosion(x-1, y-1)
			bc.RemoveBullet(b)
			return
		}
	}
}

//DetonateAllBullets detonates all bullets belonging to the passed Character
func (bc *BulletController) DetonateAllBullets(char *Character) {
	delList := []*Bullet{}
	for _, b := range bc.bullets {
		if b.owner == char {
			x, y := b.Position()
			TermGame.Arena.explosions.CreateExplosion(x-1, y-1)
			delList = append(delList, b)

		}
	}

	for _, b := range delList {
		bc.RemoveBullet(b)
	}
}

//Update the bullet
func (b *Bullet) Update() {
	if b.speedCount > 0 {
		if b.speedCount > b.Speed {
			b.speedCount = 0
		} else {
			b.speedCount++
			return
		}

	}
	x, y := b.Position()
	switch b.Velocity {
	case Up:
		b.SetPosition(x, y-1)
		break
	case Down:
		b.SetPosition(x, y+1)
		break
	case Right:
		b.SetPosition(x+1, y)
		break
	case Left:
		b.SetPosition(x-1, y)
		break
	default:
		break

	}
	b.speedCount = 1
}

//Collide function to make player DynamicPhysical and handle collision
func (b *Bullet) Collide(collision tl.Physical) {

	if b.Type == "Bullet" {
		// Check if it's a Rectangle we're colliding with
		if _, ok := collision.(*tl.Rectangle); ok {
			x, y := b.Position()
			TermGame.Arena.explosions.CreateExplosion(x-1, y-1)
			TermGame.Arena.bullets.RemoveBullet(b)

		}

		if _, ok := collision.(*NPC); ok {
			TermGame.Arena.bullets.RemoveBullet(b)
		}
		if _, ok := collision.(*Player); ok {
			TermGame.Arena.bullets.RemoveBullet(b)
		}
	}
}
