package yfantasy

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

// Game represents a Yahoo game
type Game struct {
	yf     *YFantasy
	Sport  string
	gameID string
}

// NewGame returns a new Game object.
func NewGame(sport string, yf *YFantasy) *Game {
	return &Game{yf: yf, Sport: sport}
}

// GetGameID queries the Yahoo fantasy API for the ID of the game and sets
// GameID in the Game.
func (g *Game) GetGameID() (string, error) {
	if g.gameID != "" {
		return g.gameID, nil
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

	g.gameID = node.InnerText()
	return g.gameID, err
}

// Leagues returns all the active leagues the user is in for the game.
func (g *Game) Leagues() ([]*League, error) {
	rawResp, err := g.yf.GetUserLeaguesForSport(g.Sport)
	if err != nil {
		return nil, err
	}

	return g.extractLeagues(rawResp)
}

// GetLeagueByName returns a league with the given name. If no league is found
// an error is returned.
func (g *Game) GetLeagueByName(lgName string) (*League, error) {
	lgs, err := g.Leagues()
	if err != nil {
		return nil, err
	}

	for _, lg := range lgs {
		if lg.Name == lgName {
			return lg, nil
		}
	}

	return nil, fmt.Errorf("league with name: %q not found", lgName)
}

// extractLeagues parses the raw XML response for leagues.
func (g *Game) extractLeagues(rawResp string) ([]*League, error) {
	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	nodes, err := xmlquery.QueryAll(doc, "//league")
	if err != nil {
		return nil, err
	}

	leagues := make([]*League, len(nodes))
	for i, node := range nodes {
		leagues[i], err = NewLeagueFromXML(node.OutputXML(true), g.yf)
		if err != nil {
			return nil, err
		}
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
func (g *Game) MakeLeague(leagueID string) (*League, error) {
	gameID, err := g.GetGameID()
	if err != nil {
		return nil, err
	}
	return NewLeague(fmt.Sprintf("%v.l.%v", gameID, leagueID), g.yf), nil
}
