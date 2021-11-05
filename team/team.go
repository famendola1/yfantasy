package team

import (
	"github.com/famendola1/yfantasy"
)

type Team struct {
	yf      *yfantasy.YFantasy
	TeamKey string
}

func New(yf *yfantasy.YFantasy, teamKey string) *Team {
	return &Team{yf: yf, TeamKey: teamKey}
}
