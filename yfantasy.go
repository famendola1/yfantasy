// Package yfantasy is a client implementation of the Yahoo Fantasy API
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
func (y *YFantasy) get(uri string) (string, error) {
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
	return y.get(fmt.Sprintf("game/%v", sport))
}

// GetLeagueRaw queries the /league endpoint for league information and returns
// the raw response body as a string.
// leagueKey has the format: <game_key>.l.<league_id> (ex: nba.l.1234)
func (y *YFantasy) GetLeagueRaw(leagueKey string) (string, error) {
	return y.get(fmt.Sprintf("league/%v", leagueKey))
}

// GetLeagueSettingsRaw queries the /league//settings endpoint for the league
// settings and returns the raw response body as a string.
func (y *YFantasy) GetLeagueSettingsRaw(leagueKey string) (string, error) {
	return y.get(fmt.Sprintf("league/%v/settings", leagueKey))
}

// GetLeagueStandingsRaw queries the /league//standings endpoint for the league
// standings and returns the raw response body as a string.
func (y *YFantasy) GetLeagueStandingsRaw(leagueKey string) (string, error) {
	return y.get(fmt.Sprintf("league/%v/standings", leagueKey))
}

// GetLeagueScoreboardRaw queries the /league//scoreboard endpoint for the
// league scoreboard (i.e. matchup data) and returns the raw response body as a
// string.
func (y *YFantasy) GetLeagueScoreboardRaw(leagueKey string) (string, error) {
	return y.get(fmt.Sprintf("league/%v/scoreboard", leagueKey))
}

// GetTeamRaw queries the /team endpoint for team information and returns the
// response body as a string.
// teamKey has the format: <game_key>.l.<league_id>.t.<team_id>
func (y *YFantasy) GetTeamRaw(teamKey string) (string, error) {
	return y.get(fmt.Sprintf("team/%v", teamKey))
}

// GetTeamMatchupsRaw queries the /team//matchups endpoint for matchup data for
// a team for weeks startWeek to startWeek+numWeeks-1.
func (y *YFantasy) GetTeamMatchupsRaw(teamKey string, startWeek int, numWeeks int) (string, error) {
	if numWeeks == 1 {
		return y.get(fmt.Sprintf("team/%v/matchups;weeks=%v", teamKey, startWeek))
	}
	return y.get(fmt.Sprintf("team/%v/matchups;weeks=%v,%v", teamKey, startWeek, startWeek+numWeeks-1))
}

// GetTeamStatsRaw queries the /team//stats endpoint for team stats of the given duration.
// Valid durations are: season, average_season, date, last_week, last_month.
// TODO(famendola1): Add support to specify a season or date to fetch data for.
func (y *YFantasy) GetTeamStatsRaw(teamKey string, duration string) (string, error) {
	return y.get(fmt.Sprintf("team/%v/stats;type=%v", teamKey, duration))
}
