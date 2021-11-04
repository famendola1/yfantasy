// Package game contains functionality for interacting with a Yahoo game.
package game

import (
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/famendola1/yfantasy"
	"github.com/famendola1/yfantasy/league"
)

// Game represents a Yahoo game
type Game struct {
	yf    *yfantasy.YFantasy
	Sport string
}

// New returns a new Game object.
func New(yf *yfantasy.YFantasy, sport string) *Game {
	return &Game{yf: yf, Sport: sport}
}

// GameID queries the Yahoo fantasy API for the ID of the game.
func (g *Game) GameID() (string, error) {
	rawResp, err := g.yf.GetGameRaw(g.Sport)
	if err != nil {
		return "", nil
	}

	return extractGameID(rawResp)
}

// extractGameID parsed the raw XML response for a the game id.
func extractGameID(rawResp string) (string, error) {
	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return "", err
	}

	node, err := xmlquery.Query(doc, "/fantasy_content/game/game_id")
	if err != nil {
		return "", err
	}

	return node.InnerText(), err
}

func (g *Game) Leagues() ([]*league.League, error) {
	rawResp, err := g.yf.GetUserLeaguesForSport(g.Sport)
	if err != nil {
		return nil, err
	}

	return g.extractLeagues(rawResp)
}

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
