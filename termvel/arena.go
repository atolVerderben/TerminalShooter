package termvel

import tl "github.com/JoelOtter/termloop"

//Arena represents the arena gamestate... the actual game
type Arena struct {
	level         *tl.BaseLevel
	arena         *World
	bullets       *BulletController
	explosions    *ExplosionController
	playermanager *PlayerManager
	camera        *Camera
}

//CreateArena returns an Arena pointer
func CreateArena() *Arena {
	a := &Arena{
		level: tl.NewBaseLevel(tl.Cell{
			Bg: tl.ColorWhite,
			Fg: tl.ColorWhite,
			Ch: ' ',
		})}
	return a
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

	a.playermanager = NewPlayerManager(a.arena.BaseLevel)
	level.AddEntity(TermGame.player)
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			npc := CreateNPC(30, 50, tl.ColorMagenta, level)
			level.AddEntity(npc)
			a.playermanager.AddPlayer(npc)
			break
		case 1:
			npc := CreateNPC(190, 80, tl.ColorMagenta, level)
			level.AddEntity(npc)
			a.playermanager.AddPlayer(npc)
			break
		case 2:
			npc := CreateNPC(50, 75, tl.ColorMagenta, level)
			level.AddEntity(npc)
			a.playermanager.AddPlayer(npc)
			break
		case 3:
			npc := CreateNPC(190, 75, tl.ColorMagenta, level)
			level.AddEntity(npc)
			a.playermanager.AddPlayer(npc)
			break
		}
	}
	a.playermanager.AddPlayer(TermGame.player)
	a.camera = CreateCamera(-100, 0, 40, 10, a.arena.BaseLevel, 1)
	TermGame.Screen().SetLevel(a.arena.BaseLevel)
	TermGame.player.SetLevel(a.arena.BaseLevel)

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
	a.playermanager = NewPlayerManager(a.arena.BaseLevel)
	level.AddEntity(TermGame.player)
	npc := CreateNPC(40, 45, tl.ColorMagenta, a.arena.BaseLevel)
	level.AddEntity(npc)
	a.playermanager.AddPlayer(npc)

	a.playermanager.AddPlayer(TermGame.player)
	TermGame.Screen().SetLevel(a.arena.BaseLevel)
	a.camera = CreateCamera(-100, 0, 40, 10, a.arena.BaseLevel, 0)
	TermGame.player.SetLevel(a.arena.BaseLevel)

}

//Update meets the requirements for a gamestate
func (a *Arena) Update(g *Game) GameMessage {
	if a.playermanager != nil {
		a.playermanager.Update()
	}
	a.camera.Update(g.Screen(), g.player.Character)
	if a.bullets != nil {
		a.bullets.Update()
	}
	if a.explosions != nil {
		a.explosions.Update()
	}
	return MsgNone
}

//ReturnLevel returns the base level
func (a *Arena) ReturnLevel() *tl.BaseLevel {
	return a.arena.BaseLevel
}
