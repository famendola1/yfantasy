// Package league contains functionality for interacting with a Yahoo team.
package team

import (
	"github.com/famendola1/yfantasy"
)

// Team represents a Yahoo team.
type Team struct {
	yf      *yfantasy.YFantasy
	TeamKey string
}

// New returns a new Team
func New(yf *yfantasy.YFantasy, teamKey string) *Team {
	return &Team{yf: yf, TeamKey: teamKey}
}
