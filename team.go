package yfantasy

import (
	"strings"

	"github.com/antchfx/xmlquery"
)

// Team represents a Yahoo team.
type Team struct {
	yf      *YFantasy
	TeamKey string
}

// NewTeam returns a new Team
func NewTeam(yf *YFantasy, teamKey string) *Team {
	return &Team{yf: yf, TeamKey: teamKey}
}

// Roster returns the list of players on this team.
func (t *Team) Roster() ([]*Player, error) {
	rawResp, err := t.yf.GetTeamRosterRaw(t.TeamKey)
	if err != nil {
		return nil, err
	}
	return t.extractPlayersFromRoster(rawResp)
}

// extractPlayersFromRoster parses the raw XML response from the /team//roster
// endpoint for players.
func (t *Team) extractPlayersFromRoster(rawResp string) ([]*Player, error) {
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
		players[i] = NewPlayer(t.yf, playerKey.InnerText())
	}

	return players, nil
}

// TeamID returns the ID of the team.
func (t *Team) TeamID() string {
	return strings.Split(t.TeamKey, ".t.")[1]
}

// LeagueKey returns the key of the leauge this team is in.
func (t *Team) LeagueKey() string {
	return strings.Split(t.TeamKey, ".t.")[0]
}

// AddDrop adds the specified player to the team and drops the specified player
// from the team in a single transaction.
func (t *Team) AddDrop(addPlayerKey string, dropPlayerKey string) error {
	return t.yf.PostAddDropTransaction(t.LeagueKey(), t.TeamKey, addPlayerKey, dropPlayerKey)
}
