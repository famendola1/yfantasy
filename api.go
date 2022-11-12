package yfantasy

import (
	"fmt"
	"net/http"

	"github.com/famendola1/yfantasy/query"
	"github.com/famendola1/yfantasy/schema"
)

// Enum of types when requesting for stats.
const (
	StatsTypeUnknown = iota
	StatsTypeSeason
	StatsTypeAverageSeason
	StatsTypeDate
	StatsTypeLastWeek
	StatsTypeLastMonth
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
func (yf *YFantasy) Rosters(leagueKey string) (*schema.Teams, error) {
	fc, err := query.League().Key(leagueKey).Teams().Roster().Get(yf.client)
	if err != nil {
		return nil, err
	}

	return &fc.Teams, nil
}

// Team searches the given league for a team with the provided team name.
// If the team is not found an error is returned.
func (yf *YFantasy) Team(leagueKey, teamName string) (*schema.Team, error) {
	fc, err := query.League().Key(leagueKey).Teams().Get(yf.client)
	if err != nil {
		return nil, err
	}

	for _, tm := range fc.League.Teams.Team {
		if tm.Name == teamName {
			return &tm, nil
		}
	}

	return nil, fmt.Errorf("team %q not found", teamName)
}

// TeamRoster searches the given league for a team with the provided team name
// and return's its roster. If the team is not found an error is returned.
func (yf *YFantasy) TeamRoster(leagueKey, teamName string) (*schema.Team, error) {
	fc, err := query.League().Key(leagueKey).Teams().Roster().Get(yf.client)
	if err != nil {
		return nil, err
	}

	for _, tm := range fc.League.Teams.Team {
		if tm.Name == teamName {
			return &tm, nil
		}
	}

	return nil, fmt.Errorf("team %q not found", teamName)
}

// TeamStats searches the given league for a team with the provided team name
// and return's its stats. If the team is not found an error is returned.
func (yf *YFantasy) TeamStats(leagueKey, teamName string, statsType int) (*schema.Team, error) {
	q := query.League().Key(leagueKey).Teams().Stats()
	switch statsType {
	case StatsTypeUnknown:
		return nil, fmt.Errorf("unknown stats type requested")
	case StatsTypeSeason:
		q = q.CurrentSeason()
		break
	case StatsTypeAverageSeason:
		q = q.CurrentSeasonAverage()
		break
	case StatsTypeDate:
		q = q.Today()
		break
	case StatsTypeLastWeek:
		q = q.LastWeek()
		break
	case StatsTypeLastMonth:
		q = q.LastMonth()
		break
	}

	fc, err := q.Get(yf.client)
	if err != nil {
		return nil, err
	}

	for _, tm := range fc.League.Teams.Team {
		if tm.Name == teamName {
			return &tm, nil
		}
	}

	return nil, fmt.Errorf("team %q not found", teamName)
}

// Player searches the given league for a player with the provided player name.
// If the player is not found, an error is returned. name should contain at
// least 3 letters.
func (yf *YFantasy) Player(leagueKey, name string) (*schema.Player, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("name (%q) must contain at least 3 letters", name)
	}

	fc, err := query.League().Key(leagueKey).Players().Search(name).Get(yf.client)
	if err != nil {
		return nil, err
	}

	for _, p := range fc.League.Players.Player {
		if p.Name.Full == name {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("player %q not found", name)
}

// SearchPlayers searches the given league for a players with the provided player
// name. name should contain at least 3 letters.
func (yf *YFantasy) SearchPlayers(leagueKey, name string) ([]*schema.Player, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("name (%q) must contain at least 3 letters", name)
	}

	fc, err := query.League().Key(leagueKey).Players().Search(name).Get(yf.client)
	if err != nil {
		return nil, err
	}

	var players []*schema.Player
	for _, p := range fc.League.Players.Player {
		players = append(players, &p)
	}

	return players, nil
}

// PlayerStats searches the given league for a player with the provided player name.
// and returns their average stats for the current season. If the player is not
// found, an error is returned. name should contain at least 3 letters.
func (yf *YFantasy) PlayerStats(leagueKey, name string, statsType int) (*schema.Player, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("name (%q) must contain at least 3 letters", name)
	}

	q := query.League().Key(leagueKey).Players().Search(name).Stats()
	switch statsType {
	case StatsTypeUnknown:
		return nil, fmt.Errorf("unknown stats type requested")
	case StatsTypeSeason:
		q = q.CurrentSeason()
		break
	case StatsTypeAverageSeason:
		q = q.CurrentSeasonAverage()
		break
	case StatsTypeDate:
		q = q.Today()
		break
	case StatsTypeLastWeek:
		q = q.LastWeek()
		break
	case StatsTypeLastMonth:
		q = q.LastMonth()
		break
	}

	fc, err := q.Get(yf.client)
	if err != nil {
		return nil, err
	}

	for _, p := range fc.League.Players.Player {
		if p.Name.Full == name {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("player %q not found", name)
}

// PlayerAdvancedStats searches the given league for a player with the provided
// player name and returns their advanced stats. If the player is not found, an
// error is returned. name should contain at least 3 letters.
func (yf *YFantasy) PlayerAdvancedStats(leagueKey, name string) (*schema.Player, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("name (%q) must contain at least 3 letters", name)
	}
	fc, err := query.League().Key(leagueKey).Players().Search(name).Stats().Get(yf.client)
	if err != nil {
		return nil, err
	}

	for _, p := range fc.League.Players.Player {
		if p.Name.Full == name {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("player %q not found", name)
}
