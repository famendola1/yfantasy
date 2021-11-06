// Package game contains functionality for interacting with a Yahoo game.
package game

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/famendola1/yfantasy"
	"github.com/famendola1/yfantasy/league"
)

// Game represents a Yahoo game
type Game struct {
	yf     *yfantasy.YFantasy
	Sport  string
	GameID string
}

// New returns a new Game object.
func New(yf *yfantasy.YFantasy, sport string) *Game {
	return &Game{yf: yf, Sport: sport}
}

// GetGameID queries the Yahoo fantasy API for the ID of the game and sets
// GameID in the Game.
func (g *Game) GetGameID() (string, error) {
	if g.GameID != "" {
		return g.GameID, nil
	}

	rawResp, err := g.yf.GetGameRaw(g.Sport)
	if err != nil {
		return "", nil
	}

	return g.extractGameID(rawResp)
}

// extractGameID parses the raw XML response for a the game id.
func (g *Game) extractGameID(rawResp string) (string, error) {
	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return "", err
	}

	node, err := xmlquery.Query(doc, "/fantasy_content/game/game_id")
	if err != nil {
		return "", err
	}

	g.GameID = node.InnerText()
	return g.GameID, err
}

// Leagues returns all the active leagues the user is in for the game.
func (g *Game) Leagues() ([]*league.League, error) {
	rawResp, err := g.yf.GetUserLeaguesForSport(g.Sport)
	if err != nil {
		return nil, err
	}

	return g.extractLeagues(rawResp)
}

// extractLeagues parses the raw XML response for leagues.
func (g *Game) extractLeagues(rawResp string) ([]*league.League, error) {
	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	nodes, err := xmlquery.QueryAll(doc, "//league")
	if err != nil {
		return nil, err
	}

	leagues := make([]*league.League, len(nodes))
	for i, node := range nodes {
		leagueKey, err := xmlquery.Query(node, "/league_key")
		if err != nil {
			return nil, err
		}
		leagues[i] = league.New(g.yf, leagueKey.InnerText())
	}
	return leagues, nil
}

// LeagueKeys returns all the active league keys of the leagues the user is in
// for the game.
func (g *Game) LeagueKeys() ([]string, error) {
	rawResp, err := g.yf.GetUserLeaguesForSport(g.Sport)
	if err != nil {
		return nil, err
	}

	return extractLeagueKeys(rawResp)
}

// extractLeagueKeys parses the raw XML response for all the league keys.
func extractLeagueKeys(rawResp string) ([]string, error) {
	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	nodes, err := xmlquery.QueryAll(doc, "//league_key")
	if err != nil {
		return nil, err
	}

	leagueKeys := make([]string, len(nodes))
	for i, node := range nodes {
		leagueKeys[i] = node.InnerText()
	}
	return leagueKeys, nil
}

// MakeLeague creates a League object from the given league id.
func (g *Game) MakeLeague(leagueID string) (*league.League, error) {
	gameID, err := g.GetGameID()
	if err != nil {
		return nil, err
	}
	return league.New(g.yf, fmt.Sprintf("%v.l.%v", gameID, leagueID)), nil
}
