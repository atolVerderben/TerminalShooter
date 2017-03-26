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
	states             map[string]GameState
	player             *Player
	playermanager      *PlayerManager
	input              *Input
	currTime, lastTime time.Time
	Arena              *Arena
}

func (g *Game) gameLoop() {
	g.loop = time.NewTicker(time.Millisecond * 17) //roughly 60fps (16.67)
	for {
		if g.player == nil {
			g.loop.Stop()
			return
		}
		switch g.Msg {
		case MsgStartMainSmall:
			g.currentState = g.Arena
			g.Arena.createSmallArena()
			g.Arena.SetMessage(MsgNone)
			break

		case MsgStartMainLarge:
			g.currentState = g.Arena
			g.Arena.createLargeArena()
			g.Arena.SetMessage(MsgNone)
			break
		case MsgMainMenu:
			g.currentState = g.states["MainMenu"]
			g.Screen().SetLevel(g.currentState.ReturnLevel())
			g.currentState.SetMessage(MsgNone)
			break
		}
		g.input.gameState = g.currentState
		g.currTime = time.Now()
		g.Msg = g.currentState.Update(g)
		//Update(time.Since(lastTime)) //TODO Saving this as a reminder to go back and use delta time maybe(?)
		g.lastTime = g.currTime
		<-g.loop.C
	}

}

//NewGame returns a new game
func NewGame() *Game {
	g := &Game{
		Game:   tl.NewGame(),
		Arena:  CreateArena(),
		states: map[string]GameState{},
	}
	g.states["MainMenu"] = CreateMainMenu()
	g.currentState = g.states["MainMenu"] //g.Arena
	return g
}

//Run the game
func (g *Game) Run() {
	rand.Seed(time.Now().Unix())
	GameOverInfo = nil
	//TODO: Fix this nonsense here too
	TermGame = g
	g.player = CreatePlayer(20, 20, tl.ColorWhite, nil)

	Information = NewEventInfo(0, 0)
	NPCInformation = NewEventInfo(50, 0)
	NPCInformation.right = true
	g.Screen().AddEntity(Information)
	g.Screen().AddEntity(NPCInformation)
	g.Screen().SetFps(60)
	switch cs := g.currentState.(type) {
	case *MainMenu:
		cs.ShowMainMenu(g)
	}
	g.input = NewInput()
	g.Start()
}

//Start the game... overrides termloop
func (g *Game) Start() {
	go g.gameLoop()
	g.Game.Start()
}

//TermGame is the reference to the main game
var TermGame *Game

//StopUpdate stops what was happening in the previous round/game
func (g *Game) StopUpdate() {
	if g.loop != nil {
		g.loop.Stop()
	}
	g.player = nil
	//GameCamera = nil
	//GameWorld = nil
	//GameNPCs = nil
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
