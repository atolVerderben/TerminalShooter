package termvel

import tl "github.com/JoelOtter/termloop"

//ExplosionController controls the explosions
type ExplosionController struct {
	explosions []*Explosion
	level      *tl.BaseLevel
}

//Explosion is what happens when the bullet is detonated
type Explosion struct {
	*tl.Entity
	exCount int
}

//CreateExplosion adds an explosion to the controller
func (ec *ExplosionController) CreateExplosion(x, y int) {
	e := &Explosion{
		Entity: tl.NewEntity(x, y, 3, 3),
	}
	e.SetCell(0, 0, &tl.Cell{Fg: tl.ColorYellow, Ch: '*'})
	e.SetCell(0, 1, &tl.Cell{Fg: tl.ColorYellow, Ch: '*'})
	e.SetCell(0, 2, &tl.Cell{Fg: tl.ColorYellow, Ch: '*'})
	e.SetCell(1, 0, &tl.Cell{Fg: tl.ColorYellow, Ch: '*'})
	e.SetCell(1, 1, &tl.Cell{Fg: tl.ColorYellow, Ch: '*'})
	e.SetCell(1, 2, &tl.Cell{Fg: tl.ColorYellow, Ch: '*'})
	e.SetCell(2, 0, &tl.Cell{Fg: tl.ColorYellow, Ch: '*'})
	e.SetCell(2, 1, &tl.Cell{Fg: tl.ColorYellow, Ch: '*'})
	e.SetCell(2, 2, &tl.Cell{Fg: tl.ColorYellow, Ch: '*'})
	ec.level.AddEntity(e)
	ec.AddExplosion(e)
}

//Update the explosion... mainly just remove it after a given time
func (e *Explosion) Update() {
	e.exCount++
	if e.exCount > 40 {
		TermGame.Arena.explosions.RemoveExplosion(e)
	}
}

//CreateExplosionController creates a new empty ExplosionController
func CreateExplosionController(level *tl.BaseLevel) *ExplosionController {
	ec := &ExplosionController{
		explosions: []*Explosion{},
		level:      level,
	}
	return ec
}

//AddExplosion to the ExplosionController
func (ec *ExplosionController) AddExplosion(explosion *Explosion) {
	ec.explosions = append(ec.explosions, explosion)
}

//RemoveExplosion removes the Explosion object from the slice
func (ec *ExplosionController) RemoveExplosion(explosion *Explosion) {
	//fmt.Println("Attempting to remove Explosion")
	delete := -1
	for index, e := range ec.explosions {
		if e == explosion {
			delete = index
			break
		}
	}
	if delete >= 0 {
		ec.level.RemoveEntity(explosion)
		ec.explosions = append(ec.explosions[:delete], ec.explosions[delete+1:]...)
	}

}

//Update all active explosions
func (ec *ExplosionController) Update() {
	for _, e := range ec.explosions {
		//	ec.explosions[i].Update()
		e.Update()
	}
}

//Collide function to make explosion DynamicPhysical and handle collision
func (e *Explosion) Collide(collision tl.Physical) {

	/*if _, ok := collision.(*NPC); ok {
		GameExplosion.RemoveExplosion(e)
	}
	if _, ok := collision.(*Player); ok {
		GameExplosion.RemoveExplosion(e)
	}*/

}
