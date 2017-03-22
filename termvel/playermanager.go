package termvel

import tl "github.com/JoelOtter/termloop"

//this is just a placeholder interface so I can group NPC and Player together
type participant interface {
	Update()
}

//PlayerManager keeps track of all our players
type PlayerManager struct {
	players []*Player
	ai      []*NPC
	level   *tl.BaseLevel
}

//NewPlayerManager returns a PlayerManager
func NewPlayerManager(level *tl.BaseLevel) *PlayerManager {
	p := &PlayerManager{
		level: level,
	}
	return p
}

//AddPlayer sorts the added participant into the appropriate slice
func (p *PlayerManager) AddPlayer(player participant) {
	switch guy := player.(type) {
	case *Player:
		p.players = append(p.players, guy)
		break
	case *NPC:
		p.ai = append(p.ai, guy)
		break
	}
}

//Update all our players
//TODO: Right now this is also updating the gameover screen... but that should be moved later
func (p *PlayerManager) Update() {
	for i := range p.players {
		p.players[i].Update()
	}
	dead := 0
	for i := range p.ai {
		p.ai[i].Update()
		if p.ai[i].isDead {
			dead++
		}
	}
	if dead >= len(p.ai) {
		if GameOverInfo == nil {
			GameOverInfo = NewEventInfo(0, 10)
			GameOverInfo.center = true
			p.level.AddEntity(GameOverInfo)
		}
		GameOverInfo.text.SetText("You Win! Foes are Vanquished! Press Ctrl+C to Exit")
	}
	if GameOverInfo != nil {
		GameOverInfo.OffsetX, GameOverInfo.OffsetY = p.level.Offset()
	}
}
