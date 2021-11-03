// Packake yfantasy is a client implementation of the Yahoo Fantasy API
package yfantasy

import (
	"bytes"
	"fmt"
	"net/http"
)

var (
	endpoint string = "https://fantasysports.yahooapis.com/fantasy/v2"
)

// YFantasy holds a client for interacting with the Yahoo Fantasy API.
type YFantasy struct {
	client *http.Client
}

// sendGet sends a GET request to the provided URI and returns the repsone as a
// string.
func (y *YFantasy) sendGet(uri string) (string, error) {
	resp, err := y.client.Get(fmt.Sprintf("%v/%v", endpoint, uri))
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String(), nil
}

// GetGameRaw queries the /game endpoint for game data and returns the raw
// response body as a string.
// Valid inputs for sport are: nba, nfl, mlb, nhl
func (y *YFantasy) GetGameRaw(sport string) (string, error) {
	return y.sendGet(fmt.Sprintf("game/%v", sport))
}

// GetLeagueRaw queries the /league endpoint for league information and returns
// the raw response body as a string.
// leagueKey has the format: <game_key>.l.<league_id> (ex: nba.l.1234)
func (y *YFantasy) GetLeagueRaw(leagueKey string) (string, error) {
	return y.sendGet(fmt.Sprintf("league/%v", leagueKey))
}

// GetLeagueSettingsRaw queries the /league//settings endpoint for the league
// settings and returns the raw response body as a string.
func (y *YFantasy) GetLeagueSettingsRaw(leagueKey string) (string, error) {
	return y.sendGet(fmt.Sprintf("league/%v/settings", leagueKey))
}

// GetLeagueStandingsRaw queries the /league//standings endpoint for the league
// standings and returns the raw response body as a string.
func (y *YFantasy) GetLeagueStandingsRaw(leagueKey string) (string, error) {
	return y.sendGet(fmt.Sprintf("league/%v/standings", leagueKey))
}

// GetLeagueScoreboardRaw queries the /league//scoreboard endpoint for the
// league scoreboard (i.e. matchup data) and returns the raw response body as a
// string.
func (y *YFantasy) GetLeagueScoreboardRaw(leagueKey string) (string, error) {
	return y.sendGet(fmt.Sprintf("league/%v/scoreboard", leagueKey))
}
