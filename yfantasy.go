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

// processAPIResponse converts the response body into a string.
func processAPIResponse(resp *http.Response) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String()
}

// GetGameRaw queries the /game endpoint for game data and returns the raw
// response as an XML string.
// Valid inputs for sport are: nba, nfl, mlb, nhl
func (y *YFantasy) GetGameRaw(sport string) (string, error) {
	resp, err := y.client.Get(fmt.Sprintf("%v/game/%v", endpoint, sport))
	if err != nil {
		return "", err
	}
	return processAPIResponse(resp), nil
}

// GetLeagueRaw queries the /league endpoint for league information and returns
// the raw response as an XML string.
// leagueKey has the format: <game_key>.l.<league_id> (ex: nba.l.1234)
func (y *YFantasy) GetLeagueRaw(leagueKey string) (string, error) {
	resp, err := y.client.Get(fmt.Sprintf("%v/league/%v", endpoint, leagueKey))
	if err != nil {
		return "", err
	}
	return processAPIResponse(resp), nil
}
