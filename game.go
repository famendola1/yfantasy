package yfantasy

import (
	"fmt"
)

// Game represents a Yahoo game
type Game struct {
	GameKey            string `xml:"game_key"`
	GameID             int    `xml:"game_id"`
	Name               string `xml:"name"`
	Code               string `xml:"code"`
	Type               string `xml:"type"`
	URL                string `xml:"url"`
	Season             string `xml:"season"`
	IsRegistrationOver bool   `xml:"is_registration_over"`
	IsGameOver         bool   `xml:"is_game_over"`
	IsOffseason        bool   `xml:"is_offseason"`

	yf *YFantasy
}

// Game returns a new Game object.
func (yf *YFantasy) Game(sport string) (*Game, error) {
	rawResp, err := yf.getGameRaw(sport)
	if err != nil {
		return nil, err
	}

	var g Game
	if err := parse(rawResp, "//game", &g); err != nil {
		return nil, err
	}
	g.yf = yf
	return &g, nil
}

// Leagues returns all the active leagues the user is in for the game.
func (g *Game) Leagues() ([]*League, error) {
	rawResp, err := g.yf.getUserLeaguesForSport(g.Code)
	if err != nil {
		return nil, err
	}
	return parseAllLeagues(rawResp, g.yf)
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

// LeagueKeys returns all the active league keys of the leagues the user is in
// for the game.
func (g *Game) LeagueKeys() ([]string, error) {
	rawResp, err := g.yf.getUserLeaguesForSport(g.Code)
	if err != nil {
		return nil, err
	}

	return parseAllString(rawResp, "//league_key")
}

// League creates a League object from the given league id.
func (g *Game) League(leagueKey string) (*League, error) {
	return g.yf.newLeague(fmt.Sprintf("%d.l.%s", g.GameID, leagueKey)), nil
}
