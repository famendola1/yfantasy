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

// GetGameRaw queries the Yahoo Fantasy API /game endpoint and returns the raw
// response as an XML string.
// Valid inputs for sport are: nba, nfl, mlb, nhl
func (y *YFantasy) GetGameRaw(sport string) (string, error) {
	resp, err := y.client.Get(fmt.Sprintf("%v/game/%v", endpoint, sport))
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String(), nil
}
