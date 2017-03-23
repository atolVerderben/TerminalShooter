package termvel

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

//TODO: something about all these...
var (
	//Information mainly for debugging right now
	Information *EventInfo
	//NPCInformation shows the NPC health, mainly debugging still
	NPCInformation *EventInfo
	GameOverInfo   *EventInfo
	//Clickables is a 2d slice of Clickable objects
	Clickables = [][]*Clickable{}
)

//Deselect the currently selected object
func Deselect(x, y int) {
	if x <= 0 || x >= len(Clickables) || y <= 0 || y >= len(Clickables[0]) {
		return
	}
	Clickables[x][y].Deselect()

}

//DeselectAll clickable objects
func DeselectAll() {

	for i := 0; i < len(Clickables); i++ {
		for j := 1; j < len(Clickables[0]); j++ {
			Clickables[i][j].Deselect()
		}
	}
}

//Select object at point x, y
func Select(x, y int) bool {
	if x <= 0 || x >= len(Clickables) || y <= 0 || y >= len(Clickables[0]) {
		return false
	}
	Clickables[x][y].Select()
	return true
}

//Game will represent the overall game
type Game struct {
	*tl.Game
	ArenaSize          string
	ArenaWidth         int
	ArenaHeight        int
	inputType          int
	Msg                GameMessage
	loop               *time.Ticker
	currentState       GameState
	states             []GameState
	player             *Player
	playermanager      *PlayerManager
	input              *Input
	currTime, lastTime time.Time
	Arena              *Arena
}

type GameMessage int

const (
	MsgNone GameMessage = iota
	MsgStartMain
	MsgEndGame
)

func (g *Game) gameLoop() {
	g.loop = time.NewTicker(time.Millisecond * 17) //roughly 60fps (16.67)
	for {
		if GamePlayer == nil {
			g.loop.Stop()
			return
		}
		switch g.Msg {
		case MsgStartMain:

			break
		}
		g.currentState.Update(g)
		g.currTime = time.Now()
		//Update(time.Since(lastTime))
		g.lastTime = g.currTime
		<-g.loop.C
	}

}

//NewGame returns a new game
func NewGame() *Game {
	g := &Game{
		Game:  tl.NewGame(),
		Arena: CreateArena(),
	}
	g.currentState = g.Arena
	return g
}

//Run the game
func (g *Game) Run() {
	rand.Seed(time.Now().Unix())
	GameOverInfo = nil
	//TODO: Fix this nonsense here too
	TermGame = g

	//game := tl.NewGame()
	Information = NewEventInfo(0, 0)
	NPCInformation = NewEventInfo(50, 0)
	NPCInformation.right = true
	g.Screen().AddEntity(Information)
	g.Screen().AddEntity(NPCInformation)
	g.Screen().SetFps(60)

	//game.Screen().AddEntity(tl.NewFpsText(0, 1, tl.ColorWhite, tl.ColorBlack, .2))
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})

	if g.ArenaSize == "large" {
		g.Arena.createLargeArena()
		level = g.Arena.arena.BaseLevel
		GamePlayer = CreatePlayer(20, 20, tl.ColorWhite, level)
		for i := 0; i < 4; i++ {
			switch i {
			case 0:
				GameNPCs = append(GameNPCs, CreateNPC(30, 50, tl.ColorMagenta, level))
				break
			case 1:
				GameNPCs = append(GameNPCs, CreateNPC(190, 80, tl.ColorMagenta, level))
				break
			case 2:
				GameNPCs = append(GameNPCs, CreateNPC(50, 75, tl.ColorMagenta, level))
				break
			case 3:
				GameNPCs = append(GameNPCs, CreateNPC(190, 75, tl.ColorMagenta, level))
				break
			}
		}

		GameCamera = CreateCamera(-100, 0, 40, 10, level, 1)
	}

	if g.ArenaSize == "small" {
		g.Arena.createSmallArena()
		GamePlayer = CreatePlayer(20, 20, tl.ColorWhite, g.Arena.arena.BaseLevel)
		GameNPCs = append(GameNPCs, CreateNPC(40, 45, tl.ColorMagenta, g.Arena.arena.BaseLevel))
		GameCamera = CreateCamera(-100, 0, 40, 10, g.Arena.arena.BaseLevel, 0)
		level = g.Arena.arena.BaseLevel
	}
	g.playermanager = NewPlayerManager(level)
	level.AddEntity(GamePlayer)
	for _, npc := range GameNPCs {
		level.AddEntity(npc)
		g.playermanager.AddPlayer(npc)
	}
	g.playermanager.AddPlayer(GamePlayer)
	//GameBullets = CreateBulletController(level)

	//g.Screen().SetLevel(level)
	screenWidth, screenHeight := g.Screen().Size()
	x, y := GamePlayer.Position()
	level.SetOffset(screenWidth/2-x, screenHeight/2-y)

	//GameExplosion = CreateExplosionController(level)
	/*GameWorld = &World{
		BaseLevel: level,
		Grid:      ReadAStarFile("testmap.txt"),
	}*/
	g.player = GamePlayer
	g.input = NewInput()

	g.Start()
}

//Start the game... overrides termloop
func (g *Game) Start() {
	go g.gameLoop()
	g.Game.Start()
}

//These are all the global actors of the game.
//TODO: fit these into the Game object
var (
	GamePlayer *Player
	GameNPCs   []*NPC
	GameCamera *Camera
	//GameBullets     *BulletController
	//GameWorld *World
	//GameExplosion   *ExplosionController

	TermGame *Game
)

//StopUpdate stops what was happening in the previous round/game
func (g *Game) StopUpdate() {
	if g.loop != nil {
		g.loop.Stop()
	}
	GamePlayer = nil
	GameCamera = nil
	//GameWorld = nil
	GameNPCs = nil
	g.playermanager = nil
}

//Point is a coordinate the character is moving to
type Point struct {
	X, Y int
}

//Shift removes the first element from the array and returns the value
func Shift(slc *[]*Point) *Point {
	s := *slc
	x, s := s[0], s[1:]
	*slc = s
	return x

}
