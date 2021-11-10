package yfantasy

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

// League represents a Yahoo league.
type League struct {
	yf        *YFantasy
	LeagueKey string
}

// NewLeague returns a new League object.
func NewLeague(yf *YFantasy, leagueKey string) *League {
	return &League{yf: yf, LeagueKey: leagueKey}
}

// LeagueID returns the ID of the league.
func (l *League) LeagueID() string {
	return strings.Split(l.LeagueKey, ".l.")[1]
}

// GameKey returns the game key for the league.
func (l *League) GameKey() string {
	return strings.Split(l.LeagueKey, ".l.")[0]
}

// Teams returns a list of the teams in the league
func (l *League) Teams() ([]*Team, error) {
	rawResp, err := l.yf.GetLeagueStandingsRaw(l.LeagueKey)
	if err != nil {
		return nil, err
	}

	return l.extractTeams(rawResp)
}

// UserTeam returns the team that the user has in this league.
func (l *League) UserTeam() (*Team, error) {
	rawResp, err := l.yf.GetUserTeamInLeagueRaw(l.GameKey(), l.LeagueKey)
	if err != nil {
		return nil, err
	}

	teams, err := l.extractTeams(rawResp)
	if err != nil {
		return nil, err
	}

	if len(teams) == 0 {
		return nil, fmt.Errorf("user has no teams in this league")
	}
	return teams[0], nil
}

// extractTeams parses the raw XML response from the
// /league//standings endpoint for teams.
func (l *League) extractTeams(rawResp string) ([]*Team, error) {
	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	nodes, err := xmlquery.QueryAll(doc, "//team")
	if err != nil {
		return nil, err
	}

	teams := make([]*Team, len(nodes))
	for i, node := range nodes {
		teamKey, err := xmlquery.Query(node, "/team_key")
		if err != nil {
			return nil, err
		}
		teams[i] = NewTeam(l.yf, teamKey.InnerText())
	}

	return teams, nil
}

// SearchPlayers searches for players using the provided name.
// playerName can be the player's full name or a partial name.
func (l *League) SearchPlayers(playerName string) ([]*Player, error) {
	rawResp, err := l.yf.GetPlayersBySearchRaw(l.LeagueKey, playerName)
	if err != nil {
		return nil, err
	}

	return l.extractPlayersFromSearch(rawResp)
}

// extractPlayersFromSearch extracts players from the search results.
func (l *League) extractPlayersFromSearch(rawResp string) ([]*Player, error) {
	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	nodes, err := xmlquery.QueryAll(doc, "//player")
	if err != nil {
		return nil, err
	}

	players := make([]*Player, len(nodes))
	for i, node := range nodes {
		playerKey, err := xmlquery.Query(node, "/player_key")
		if err != nil {
			return nil, err
		}
		players[i] = NewPlayer(l.yf, playerKey.InnerText())
	}

	return players, nil
}
