package termvel

import tl "github.com/JoelOtter/termloop"

type GameState interface {
	Update(*Game)
}

type Arena struct {
	level      *tl.BaseLevel
	arena      *World
	bullets    *BulletController
	explosions *ExplosionController
}

func CreateArena() *Arena {
	m := &Arena{
		level: tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorWhite,
			Fg: tl.ColorWhite,
			Ch: ' ',
		})}
	return m
}

func (a *Arena) createLargeArena() {

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})

	//Add clickable terrain
	for i := 0; i < 200; i++ {
		row := []*Clickable{}
		for j := 0; j < 100; j++ {
			click := NewClickable(i, j, 1, 1, tl.ColorGreen, level)
			level.AddEntity(click)
			row = append(row, click)
		}
		Clickables = append(Clickables, row)
	}

	//Level "Bounds"
	level.AddEntity(tl.NewRectangle(0, 0, 200, 1, tl.ColorDefault))
	level.AddEntity(tl.NewRectangle(0, 99, 200, 1, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(0, 1, 1, 99, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(199, 1, 1, 99, tl.ColorBlack))

	//Lake
	level.AddEntity(NewWater(100, 10, 75, 35))

	a.arena = &World{
		BaseLevel: level,
		Grid:      ReadAStarFile("testmap.txt"),
		Height:    100,
		Width:     200,
	}
	a.bullets = CreateBulletController(a.arena.BaseLevel)
	a.explosions = CreateExplosionController(a.arena.BaseLevel)
	TermGame.Screen().SetLevel(a.arena.BaseLevel)
}

func (a *Arena) createSmallArena() {

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})
	//Add clickable terrain
	for i := 0; i < 100; i++ {
		row := []*Clickable{}
		for j := 0; j < 50; j++ {
			click := NewClickable(i, j, 1, 1, tl.ColorGreen, level)
			level.AddEntity(click)
			row = append(row, click)
		}
		Clickables = append(Clickables, row)
	}

	//Level "Bounds"
	level.AddEntity(tl.NewRectangle(0, 0, 100, 1, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(0, 49, 100, 1, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(0, 1, 1, 49, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(99, 1, 1, 49, tl.ColorBlack))
	a.arena = &World{
		BaseLevel: level,
		Grid:      ReadAStarFile("testmap.txt"),
		Height:    50,
		Width:     100,
	}
	a.bullets = CreateBulletController(a.arena.BaseLevel)
	a.explosions = CreateExplosionController(a.arena.BaseLevel)
	TermGame.Screen().SetLevel(a.arena.BaseLevel)
}

func (a *Arena) Update(g *Game) {
	if g.playermanager != nil {
		g.playermanager.Update()
	}
	if a.bullets != nil {
		a.bullets.Update()
	}
	if a.explosions != nil {
		a.explosions.Update()
	}
}
