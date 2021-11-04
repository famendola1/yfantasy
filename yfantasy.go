// Package yfantasy is a client implementation of the Yahoo Fantasy API
package yfantasy

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

const (
	endpoint string = "https://fantasysports.yahooapis.com/fantasy/v2"
)

// YFantasy holds a client for interacting with the Yahoo Fantasy API.
type YFantasy struct {
	client *http.Client
}

// New returns a new YFantasy object.
func New(client *http.Client) *YFantasy {
	return &YFantasy{client: client}
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
	return y.get(fmt.Sprintf("team/%v/matchups;weeks=%v,%v", teamKey, startWeek,
		startWeek+numWeeks-1))
}

// GetTeamStatsRaw queries the /team//stats endpoint for team stats of the given
// duration.
// Valid durations are: season, average_season, date, last_week, last_month.
// TODO(famendola1): Add support to specify a season or date to fetch data for.
func (y *YFantasy) GetTeamStatsRaw(teamKey string, duration string) (string, error) {
	return y.get(fmt.Sprintf("team/%v/stats;type=%v", teamKey, duration))
}

// GetTeamRosterRaw queries the /team//roster endpoint for the team's current
// day roster and returns the raw response as a string.
func (y *YFantasy) GetTeamRosterRaw(teamKey string) (string, error) {
	return y.get(fmt.Sprintf("team/%v/roster", teamKey))
}

// GetTeamRosterWeekRaw queries the /team//roster endpoint for the team's roster
// for the given week and returns the raw response as a string.
func (y *YFantasy) GetTeamRosterWeekRaw(teamKey string, weekNum int) (string, error) {
	return y.get(fmt.Sprintf("team/%v/roster;week=%v", teamKey, weekNum))
}

// GetTeamRosterDayRaw queries the /team//roster endpoint for the team's roster
// for the given day and returns the raw response as a string.
// day is formatted as: yyyy-mm-dd
func (y *YFantasy) GetTeamRosterDayRaw(teamKey string, day string) (string, error) {
	return y.get(fmt.Sprintf("team/%v/roster;date=%v", teamKey, day))
}

// GetPlayerRaw queries the /league//players endpoint for player data and
// returns the raw response as a string.
// playerKey is formatted as: <game_key>.p.<player_id>
func (y *YFantasy) GetPlayerRaw(leageuKey string, playerKey string) (string, error) {
	return y.get(fmt.Sprintf("league/%v/players;player_keys=%v", leageuKey, playerKey))
}

// GetPlayersRaw queries the /league//players endpoint for player data for
// multiple players and returns the raw response as a string.
func (y *YFantasy) GetPlayersRaw(leagueKey string, playerKeys []string) (string, error) {
	return y.get(fmt.Sprintf("league/%v/players;player_keys=%v", leagueKey, strings.Join(playerKeys, ",")))
}

// TODO(famendola1): Query /league//players//stats endpoint

// TODO(famendola1): Query /league//players using filters

// GetTransactionsRaw queries the /league//transactions endpoint for league
// transactions of the given type and returns the raw response as a string.
// Valid transactionTypes are: add, drop, commish, trade
func (y *YFantasy) GetTransactionsRaw(leagueKey string, transactionType string) (string, error) {
	return y.get(fmt.Sprintf("league/%v/transactions;type=%v", leagueKey, transactionType))
}

// GetTeamTransactionsRaw queries the /league//transactions endpoint for league
// transactions of the given type for the given team and returns the raw
// response as a string.
// Valid transactionTypes are: pending_trade, waiver
func (y *YFantasy) GetTeamTransactionsRaw(leagueKey string, teamKey string, transactionType string) (string, error) {
	return y.get(fmt.Sprintf("league/%v/transactions;team_key=%v;type=%v", leagueKey, teamKey, transactionType))
}

// TODO(famendola1): Implement more filters for transactions.

// GetUserTeams queries the /users endpoint to get all the teams for the logged
// in user and returns the raw response as a string.
func (y *YFantasy) GetUserTeams() (string, error) {
	return y.get(fmt.Sprintf("users;use_login=1/games/teams"))
}

// GetUserTeamsForSport is the same as GetUserTeams except the teams
// restricted to the given sport and only active teams are returned.
func (y *YFantasy) GetUserTeamsForSport(sport string) (string, error) {
	return y.get(fmt.Sprintf("users;use_login=1/games;game_keys=%v/teams", sport))
}

// GetUserLeagues queries the /users endpoint to get all the leagues for the
// logged in user and returns the raw response as a string.
func (y *YFantasy) GetUserLeagues() (string, error) {
	return y.get(fmt.Sprintf("users;use_login=1/games/leagues"))
}

// GetUserLeaguesForSport is the same as GetUserLeagues except the leagues
// restricted to the given sport and only active leagues are returned.
func (y *YFantasy) GetUserLeaguesForSport(sport string) (string, error) {
	return y.get(fmt.Sprintf("users;use_login=1/games;game_keys=%v/leagues", sport))
}
