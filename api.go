package yfantasy

import (
	"fmt"
	"net/http"

	query "github.com/famendola1/yfantasy/query"
	schema "github.com/famendola1/yfantasy/schema"
)

type YFantasy struct {
	client *http.Client
}

func MakeLeagueKey(gameKey string, leagueID int) string {
	return fmt.Sprintf("%s.l.%d", gameKey, leagueID)
}

func (yf *YFantasy) League(leagueKey string) (*schema.League, error) {
	fc, err := query.League().Key(leagueKey).Standings().Get(yf.client)
	if err != nil {
		return nil, err
	}
	return &fc.League, nil
}

func (yf *YFantasy) Standings(leagueKey string) (*schema.Standings, error) {
	fc, err := query.League().Key(leagueKey).Standings().Get(yf.client)
	if err != nil {
		return nil, err
	}

	return &fc.League.Standings, nil
}

func (yf *YFantasy) CurrentScoreboard(leagueKey string) (*schema.Scoreboard, error) {
	fc, err := query.League().Key(leagueKey).CurrentScoreboard().Get(yf.client)
	if err != nil {
		return nil, err
	}

	return &fc.League.Scoreboard, nil
}

func (yf *YFantasy) Scoreboard(leagueKey string, week int) (*schema.Scoreboard, error) {
	fc, err := query.League().Key(leagueKey).Scoreboard(week).Get(yf.client)
	if err != nil {
		return nil, err
	}

	return &fc.League.Scoreboard, nil
}

func (yf *YFantasy) Rosters(leagueKey string) ([]*schema.Roster, error) {
	fc, err := query.League().Key(leagueKey).Teams().Roster().Get(yf.client)
	if err != nil {
		return nil, err
	}

	rosters := []*schema.Roster{}
	for _, tm := range fc.League.Teams.Team {
		rosters = append(rosters, &tm.Roster)
	}
	return rosters, nil
}
