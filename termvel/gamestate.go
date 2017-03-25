package termvel

import tl "github.com/JoelOtter/termloop"

//GameState represents a different stage of the game
type GameState interface {
	Update(*Game) GameMessage
	ReturnLevel() *tl.BaseLevel
}

//GameMessage signals when to switch between states
type GameMessage int

//Available GameMessages
const (
	MsgNone GameMessage = iota
	MsgStartMainSmall
	MsgStartMainLarge
	MsgEndGame
)
