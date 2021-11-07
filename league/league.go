// Package league contains functionality for interacting with a Yahoo league.
package league

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/famendola1/yfantasy"
	"github.com/famendola1/yfantasy/player"
	"github.com/famendola1/yfantasy/team"
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

// LeagueID returns the ID of the league.
func (l *League) LeagueID() string {
	return strings.Split(l.LeagueKey, ".l.")[1]
}

// GameKey returns the game key for the league.
func (l *League) GameKey() string {
	return strings.Split(l.LeagueKey, ".l.")[0]
}

// Teams returns a list of the teams in the league
func (l *League) Teams() ([]*team.Team, error) {
	rawResp, err := l.yf.GetLeagueStandingsRaw(l.LeagueKey)
	if err != nil {
		return nil, err
	}

	return l.extractTeams(rawResp)
}

// UserTeam returns the team that the user has in this league.
func (l *League) UserTeam() (*team.Team, error) {
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
func (l *League) extractTeams(rawResp string) ([]*team.Team, error) {
	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	nodes, err := xmlquery.QueryAll(doc, "//team")
	if err != nil {
		return nil, err
	}

	teams := make([]*team.Team, len(nodes))
	for i, node := range nodes {
		teamKey, err := xmlquery.Query(node, "/team_key")
		if err != nil {
			return nil, err
		}
		teams[i] = team.New(l.yf, teamKey.InnerText())
	}

	return teams, nil
}

// SearchPlayers searches for players using the provided name.
// playerName can be the player's full name or a partial name.
func (l *League) SearchPlayers(playerName string) ([]*player.Player, error) {
	rawResp, err := l.yf.GetPlayersBySearchRaw(l.LeagueKey, playerName)
	if err != nil {
		return nil, err
	}

	return l.extractPlayersFromSearch(rawResp)
}

// extractPlayersFromSearch extracts players from the search results.
func (l *League) extractPlayersFromSearch(rawResp string) ([]*player.Player, error) {
	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	nodes, err := xmlquery.QueryAll(doc, "//player")
	if err != nil {
		return nil, err
	}

	players := make([]*player.Player, len(nodes))
	for i, node := range nodes {
		playerKey, err := xmlquery.Query(node, "/player_key")
		if err != nil {
			return nil, err
		}
		players[i] = player.New(l.yf, playerKey.InnerText())
	}

	return players, nil
}
