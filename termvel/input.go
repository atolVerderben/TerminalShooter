package termvel

import tl "github.com/JoelOtter/termloop"

//Input represents the player input
type Input struct {
	Keys map[string]tl.Key
}

//NewInput returns the default Input struct
func NewInput() *Input {
	i := &Input{
		Keys: map[string]tl.Key{},
	}

	i.Keys["Up"] = tl.KeyCtrlW
	i.Keys["Down"] = tl.KeyCtrlA
	i.Keys["Left"] = tl.KeyArrowLeft
	i.Keys["Right"] = tl.KeyArrowRight

	i.Keys["AimUp"] = tl.KeyArrowUp
	i.Keys["AimDown"] = tl.KeyArrowDown
	i.Keys["AimLeft"] = tl.KeyArrowLeft
	i.Keys["AimRight"] = tl.KeyArrowRight

	return i
}
