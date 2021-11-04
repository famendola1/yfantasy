// Package league contains functionality for interacting with a Yahoo league.
package league

import (
	"github.com/famendola1/yfantasy"
)

// League represents a Yahoo league.
type League struct {
	yf        *yfantasy.YFantasy
	LeagueKey string
}

// New returns a new League object.
func New(yf *yfantasy.YFantasy, leagueKey string) *League {
	return &League{yf: yf, LeagueKey: leagueKey}
}
