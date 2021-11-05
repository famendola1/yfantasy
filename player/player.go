// Package player constains functionality for interacting with a Yahoo player.
package player

import (
	"github.com/famendola1/yfantasy"
)

// Player represents a Yahoo player.
type Player struct {
	yf        *yfantasy.YFantasy
	PlayerKey string
}

// New returns a new player.
func New(yf *yfantasy.YFantasy, playerKey string) *Player {
	return &Player{yf: yf, PlayerKey: playerKey}
}
