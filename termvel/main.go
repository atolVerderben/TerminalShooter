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
	ArenaSize   string
	ArenaWidth  int
	ArenaHeight int
}

//NewGame returns a new game
func NewGame() *Game {
	g := &Game{
		Game: tl.NewGame(),
	}
	return g
}

//Run the game
func (g *Game) Run() {
	rand.Seed(time.Now().Unix())
	GameOverInfo = nil
	//game := tl.NewGame()
	Information = NewEventInfo(0, 0)
	NPCInformation = NewEventInfo(50, 0)
	NPCInformation.right = true
	g.Screen().AddEntity(Information)
	g.Screen().AddEntity(NPCInformation)
	g.Screen().SetFps(60)

	//game.Screen().AddEntity(tl.NewFpsText(0, 1, tl.ColorWhite, tl.ColorBlack, .2))
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlue,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})

	if g.ArenaSize == "large" {
		GameArenaHeight = 100
		GameArenaWidth = 200

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

		GamePlayer = CreatePlayer(20, 20, tl.ColorWhite, level)
		GameNPC = CreateNPC(30, 50, tl.ColorMagenta, level)
		GameCamera = CreateCamera(-100, 0, 40, 10, level, 1)
	}

	if g.ArenaSize == "small" {
		GameArenaHeight = 50
		GameArenaWidth = 100
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
		GamePlayer = CreatePlayer(20, 20, tl.ColorWhite, level)
		GameNPC = CreateNPC(40, 45, tl.ColorMagenta, level)
		GameCamera = CreateCamera(-100, 0, 40, 10, level, 0)
	}

	level.AddEntity(GamePlayer)
	level.AddEntity(GameNPC)
	GameBullets = CreateBulletController(level)

	g.Screen().SetLevel(level)
	screenWidth, screenHeight := g.Screen().Size()
	x, y := GamePlayer.Position()
	level.SetOffset(screenWidth/2-x, screenHeight/2-y)

	GameExplosion = CreateExplosionController(level)
	GameWorld = &World{
		BaseLevel: level,
		Grid:      ReadAStarFile("testmap.txt"),
	}
	//go UpdateLoop()
	g.Start()
}

//Start the game... overrides termloop
func (g *Game) Start() {
	go UpdateLoop()
	g.Game.Start()
}

//GameLoop is the update loop of the game naturally....
var GameLoop *time.Ticker
var hey = 0

//These are all the global actors of the game.
//TODO: fit these into the Game object
var (
	GamePlayer      *Player
	GameNPC         *NPC
	GameCamera      *Camera
	GameBullets     *BulletController
	GameWorld       *World
	GameExplosion   *ExplosionController
	MainGame        *Game
	GameArenaWidth  int
	GameArenaHeight int
)

//UpdateLoop runs at close to 60 fps
func UpdateLoop() {
	GameLoop := time.NewTicker(time.Millisecond * 17) //roughly 60fps (16.67)
	for {
		if GamePlayer == nil {
			GameLoop.Stop()
			return
		}
		currTime = time.Now()
		Update(time.Since(lastTime))
		lastTime = currTime
		<-GameLoop.C
	}

}

//StopUpdate stops what was happening in the previous round/game
func StopUpdate() {
	if GameLoop != nil {
		GameLoop.Stop()
	}
	GamePlayer = nil
	GameNPC = nil
	GameCamera = nil
	GameBullets = nil
	GameWorld = nil
	GameExplosion = nil
}

//Update runs all the updates when fired
func Update(time time.Duration) {
	GamePlayer.Update()
	GameNPC.Update()
	GameBullets.Update()
	GameExplosion.Update()
}

var currTime = time.Now()
var lastTime = currTime

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
