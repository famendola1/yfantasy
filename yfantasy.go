// Package yfantasy is a client implementation of the Yahoo Fantasy API
package yfantasy

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
)

const (
	endpoint = "https://fantasysports.yahooapis.com/fantasy/v2"
)

// YFantasy holds a client for interacting with the Yahoo Fantasy API.
type YFantasy struct {
	client *http.Client
}

// StatDuration represents a duration of team or player stats. For example, the
// average season stats for Player A, or last week's stats for Team B.
type StatDuration struct {
	// Valid types are season, average_season, date, last_week, last_month.
	DurationType string
	// Date formattted as YYYY-MM-DD.
	Date string
	// Year the season started in as YYYY.
	Season string
}

// New returns a new YFantasy object.
func New(client *http.Client) *YFantasy {
	return &YFantasy{client: client}
}

// get sends a GET request to the provided URI and returns the repsone as a
// string.
func (y *YFantasy) get(uri string) (string, error) {
	resp, err := y.client.Get(fmt.Sprintf("%s/%s", endpoint, uri))
	if err != nil {
		return "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", handleError(resp)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String(), nil
}

// post sends a POST request to the provided URI and returns the response as a
// string.
func (y *YFantasy) post(uri string, data string) error {
	resp, err := y.client.Post(fmt.Sprintf("%s/%s", endpoint, uri),
		"application/xml", strings.NewReader(data))
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return handleError(resp)
	}

	return nil
}

// handleError returns an error containing the error message in the response.
func handleError(resp *http.Response) error {
	doc, err := xmlquery.Parse(resp.Body)
	if err != nil {
		return err
	}

	node, err := xmlquery.Query(doc, "//description")
	if err != nil {
		return err
	}

	return fmt.Errorf("%s: %s", resp.Status, node.InnerText())
}

// GetGameRaw queries the /game endpoint for game data and returns the raw
// response body as a string.
// Valid inputs for sport are: nba, nfl, mlb, nhl
func (y *YFantasy) GetGameRaw(sport string) (string, error) {
	return y.get(fmt.Sprintf("game/%s", sport))
}

// GetLeagueRaw queries the /league endpoint for league information and returns
// the raw response body as a string.
// leagueKey has the format: <game_key>.l.<league_id> (ex: nba.l.1234)
func (y *YFantasy) GetLeagueRaw(leagueKey string) (string, error) {
	return y.get(fmt.Sprintf("league/%s", leagueKey))
}

// GetLeagueSettingsRaw queries the /league//settings endpoint for the league
// settings and returns the raw response body as a string.
func (y *YFantasy) GetLeagueSettingsRaw(leagueKey string) (string, error) {
	return y.get(fmt.Sprintf("league/%s/settings", leagueKey))
}

// GetLeagueStandingsRaw queries the /league//standings endpoint for the league
// standings and returns the raw response body as a string.
func (y *YFantasy) GetLeagueStandingsRaw(leagueKey string) (string, error) {
	return y.get(fmt.Sprintf("league/%s/standings", leagueKey))
}

// GetLeagueScoreboardRaw queries the /league//scoreboard endpoint for the
// league scoreboard (i.e. matchup data) for a given week and returns the
// raw response body as a string. Passing in 0 for the week, will return the
// scorebaord of the current week.
// TODO(famendola1): Use startWeek and numWeeks instead of week.
func (y *YFantasy) GetLeagueScoreboardRaw(leagueKey string, week int) (string, error) {
	if week == 0 {
		return y.get(fmt.Sprintf("league/%s/scoreboard", leagueKey))
	}
	return y.get(fmt.Sprintf("league/%s/scoreboard;week=%d", leagueKey, week))

}

// GetTeamRaw queries the /team endpoint for team information and returns the
// response body as a string.
// teamKey has the format: <game_key>.l.<league_id>.t.<team_id>
func (y *YFantasy) GetTeamRaw(teamKey string) (string, error) {
	return y.get(fmt.Sprintf("team/%s", teamKey))
}

// GetTeamMatchupsRaw queries the /team//matchups endpoint for matchup data for
// a team for weeks startWeek to startWeek+numWeeks-1.
func (y *YFantasy) GetTeamMatchupsRaw(teamKey string, startWeek int, numWeeks int) (string, error) {
	weeks := make([]string, numWeeks)
	for i := 0; i < numWeeks; i++ {
		weeks[i] = strconv.Itoa(startWeek + i)
	}
	return y.get(fmt.Sprintf("team/%s/matchups;weeks=%s", teamKey, strings.Join(weeks, ",")))
}

// GetTeamStatsRaw queries the /team//stats endpoint for team stats of the given
// duration.
func (y *YFantasy) GetTeamStatsRaw(teamKey string, duration StatDuration) (string, error) {
	if duration.DurationType == "lastweek" || duration.DurationType == "lastmonth" {
		return y.get(fmt.Sprintf("team/%s/stats;type=%s", teamKey, duration.DurationType))
	}

	if duration.DurationType == "date" {
		date := duration.Date
		if date == "" {
			date = time.Now().Format("2006-01-02")
		}
		return y.get(fmt.Sprintf("team/%s/stats;type=%s;date=%s", teamKey, duration.DurationType, date))
	}

	if duration.DurationType == "season" || duration.DurationType == "average_season" {
		if duration.Season == "" {
			return y.get(fmt.Sprintf("team/%s/stats;type=%s", teamKey, duration.DurationType))
		}
		return y.get(fmt.Sprintf("team/%s/stats;type=%s;season=%s", teamKey, duration.DurationType, duration.Season))
	}

	return "", fmt.Errorf("requested duration invalid or not supported")
}

// GetTeamRosterRaw queries the /team//roster endpoint for the team's current
// day roster and returns the raw response as a string.
func (y *YFantasy) GetTeamRosterRaw(teamKey string) (string, error) {
	return y.get(fmt.Sprintf("team/%s/roster", teamKey))
}

// GetTeamRosterWeekRaw queries the /team//roster endpoint for the team's roster
// for the given week and returns the raw response as a string.
func (y *YFantasy) GetTeamRosterWeekRaw(teamKey string, weekNum int) (string, error) {
	return y.get(fmt.Sprintf("team/%s/roster;week=%d", teamKey, weekNum))
}

// GetTeamRosterDayRaw queries the /team//roster endpoint for the team's roster
// for the given day and returns the raw response as a string.
// day is formatted as: yyyy-mm-dd
func (y *YFantasy) GetTeamRosterDayRaw(teamKey string, day string) (string, error) {
	return y.get(fmt.Sprintf("team/%s/roster;date=%s", teamKey, day))
}

// GetPlayerRaw queries the /league//players endpoint for player data and
// returns the raw response as a string.
// playerKey is formatted as: <game_key>.p.<player_id>
func (y *YFantasy) GetPlayerRaw(leageuKey string, playerKey string) (string, error) {
	return y.get(
		fmt.Sprintf("league/%s/players;player_keys=%s", leageuKey, playerKey))
}

// GetPlayersRaw queries the /league//players endpoint for player data for
// multiple players and returns the raw response as a string.
func (y *YFantasy) GetPlayersRaw(leagueKey string, playerKeys []string) (string, error) {
	return y.get(fmt.Sprintf("league/%s/players;player_keys=%s", leagueKey, strings.Join(playerKeys, ",")))
}

// GetPlayersBySearchRaw queries the /league//players endpoint with the "search"
// filter set to the provided search string.
func (y *YFantasy) GetPlayersBySearchRaw(leagueKey string, searchStr string) (string, error) {
	return y.get(fmt.Sprintf("league/%s/players;search=%s", leagueKey, url.QueryEscape(searchStr)))
}

// GetPlayersStatsRaw queries the /league//players//stats endpoint for stats of
// all the requested players of the given duration. Only 25 players can be
// requested at once.
func (y *YFantasy) GetPlayersStatsRaw(leagueKey string, playerKeys []string, duration StatDuration) (string, error) {
	players := strings.Join(playerKeys, ",")
	if duration.DurationType == "lastweek" || duration.DurationType == "lastmonth" {
		return y.get(fmt.Sprintf("league/%s/players;player_keys=%s/stats;type=%s", leagueKey, players, duration.DurationType))
	}

	if duration.DurationType == "date" {
		date := duration.Date
		if date == "" {
			date = time.Now().Format("2006-01-02")
		}
		return y.get(fmt.Sprintf("league/%s/players;player_keys=%s/stats;type=%s;date=%s", leagueKey, players, duration.DurationType, date))
	}

	if duration.DurationType == "season" || duration.DurationType == "average_season" {
		if duration.Season == "" {
			return y.get(fmt.Sprintf("league/%s/players;player_keys=%s/stats;type=%s", leagueKey, players, duration.DurationType))
		}
		return y.get(fmt.Sprintf("league/%s/players;player_keys=%s/stats;type=%s;season=%s", leagueKey, players, duration.DurationType, duration.Season))
	}

	return "", fmt.Errorf("requested duration invalid or not supported")
}

// TODO(famendola1): Query /league//players using filters

// GetTransactionsRaw queries the /league//transactions endpoint for league
// transactions of the given types and returns the raw response as a string.
// Valid transactionTypes are: add, drop, commish, trade
func (y *YFantasy) GetTransactionsRaw(leagueKey string, transactionTypes []string) (string, error) {
	return y.get(fmt.Sprintf("league/%s/transactions;types=%s", leagueKey, strings.Join(transactionTypes, ",")))
}

// GetTeamTransactionsRaw queries the /league//transactions endpoint for league
// transactions of the given type for the given team and returns the raw
// response as a string.
// Valid transactionTypes are: pending_trade, waiver
func (y *YFantasy) GetTeamTransactionsRaw(leagueKey string, teamKey string, transactionType string) (string, error) {
	return y.get(fmt.Sprintf("league/%s/transactions;team_key=%s;type=%s",
		leagueKey, teamKey, transactionType))
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
	return y.get(fmt.Sprintf("users;use_login=1/games;game_keys=%s/teams", sport))
}

// GetUserLeagues queries the /users endpoint to get all the leagues for the
// logged in user and returns the raw response as a string.
func (y *YFantasy) GetUserLeagues() (string, error) {
	return y.get(fmt.Sprintf("users;use_login=1/games/leagues"))
}

// GetUserLeaguesForSport is the same as GetUserLeagues except the leagues
// restricted to the given sport and only active leagues are returned.
func (y *YFantasy) GetUserLeaguesForSport(sport string) (string, error) {
	return y.get(fmt.Sprintf("users;use_login=1/games;game_keys=%s/leagues", sport))
}

// GetUserTeamInLeagueRaw queries the /users endpoint for the user's team in the
// given league.
func (y *YFantasy) GetUserTeamInLeagueRaw(gameKey string, leagueKey string) (string, error) {
	return y.get(fmt.Sprintf("users;use_login=1/games;game_keys=%s/leagues;league_keys=%s/teams", gameKey, leagueKey))
}

// PostAddDropTransaction sends a POST request to the /league//transactions
// endpoint to add and drop the selected player for the given team. The raw
// response is returned as a string.
func (y *YFantasy) PostAddDropTransaction(leagueKey string, teamKey string, addPlayerKey string, dropPlayerKey string) error {
	data := fmt.Sprintf(addDropTransaction, addPlayerKey, teamKey, dropPlayerKey, teamKey)
	return y.post(fmt.Sprintf("league/%s/transactions", leagueKey), data)
}

// PostAddTransaction sends a POST request to the /league//transactions
// endpoint to add the selected player for the given team. The raw response is
// returned as a string.
func (y *YFantasy) PostAddTransaction(leagueKey string, teamKey string, addPlayerKey string) error {
	data := fmt.Sprintf(addTransaction, addPlayerKey, teamKey)
	return y.post(fmt.Sprintf("league/%s/transactions", leagueKey), data)
}

// PostDropTransaction sends a POST request to the /league//transactions
// endpoint to drop the selected player for the given team. The raw response is
// returned as a string.
func (y *YFantasy) PostDropTransaction(leagueKey string, teamKey string, dropPlayerKey string) error {
	data := fmt.Sprintf(dropTransaction, teamKey, dropPlayerKey)
	return y.post(fmt.Sprintf("league/%s/transactions", leagueKey), data)
}
