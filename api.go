package yfantasy

import (
	"fmt"
	"net/http"

	"github.com/famendola1/yfantasy/query"
	"github.com/famendola1/yfantasy/schema"
)

// YFantasy is the client for the Yahoo Fantasy API.
type YFantasy struct {
	client *http.Client
}

// New returns a new YFantasy object.
func New(client *http.Client) *YFantasy {
	return &YFantasy{client: client}
}

// MakeLeagueKey creates a league key from the gameKey and leagueID.
func MakeLeagueKey(gameKey string, leagueID int) string {
	return fmt.Sprintf("%s.l.%d", gameKey, leagueID)
}

// League queries the Yahoo Fantasy API for a League.
func (yf *YFantasy) League(leagueKey string) (*schema.League, error) {
	fc, err := query.League().Key(leagueKey).Standings().Get(yf.client)
	if err != nil {
		return nil, err
	}
	return &fc.League, nil
}

// Standings queries the Yahoo Fantasy API for a leagues Standings.
func (yf *YFantasy) Standings(leagueKey string) (*schema.Standings, error) {
	fc, err := query.League().Key(leagueKey).Standings().Get(yf.client)
	if err != nil {
		return nil, err
	}

	return &fc.League.Standings, nil
}

// CurrentScoreboard queries the Yahoo Fantasy API for a league's current scoreboard.
func (yf *YFantasy) CurrentScoreboard(leagueKey string) (*schema.Scoreboard, error) {
	fc, err := query.League().Key(leagueKey).CurrentScoreboard().Get(yf.client)
	if err != nil {
		return nil, err
	}

	return &fc.League.Scoreboard, nil
}

// Scoreboard queries the Yahoo Fantasy API for the scoreboard of a given week.
func (yf *YFantasy) Scoreboard(leagueKey string, week int) (*schema.Scoreboard, error) {
	fc, err := query.League().Key(leagueKey).Scoreboard(week).Get(yf.client)
	if err != nil {
		return nil, err
	}

	return &fc.League.Scoreboard, nil
}

// Rosters queries the Yahoo Fantasy API for all the team rosters in a league.
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
