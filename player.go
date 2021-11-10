package yfantasy

// Player represents a Yahoo player.
type Player struct {
	yf        *YFantasy
	PlayerKey string
}

// NewPlayer returns a new player.
func NewPlayer(yf *YFantasy, playerKey string) *Player {
	return &Player{yf: yf, PlayerKey: playerKey}
}
