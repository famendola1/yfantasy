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

// IsValid returns whether or not the YFantasy instance is valid.
func (yf *YFantasy) IsValid() bool {
	return yf.client != nil
}

// get sends a GET request to the provided URI and returns the repsone as a
// string.
func (yf *YFantasy) get(uri string) (string, error) {
	resp, err := yf.client.Get(fmt.Sprintf("%s/%s", endpoint, uri))
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
func (yf *YFantasy) post(uri string, data string) error {
	resp, err := yf.client.Post(fmt.Sprintf("%s/%s", endpoint, uri),
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

// getGameRaw queries the /game endpoint for game data and returns the raw
// response body as a string.
// Valid inputs for sport are: nba, nfl, mlb, nhl
func (yf *YFantasy) getGameRaw(sport string) (string, error) {
	return yf.get(fmt.Sprintf("game/%s", sport))
}

// getAvailableGamesRaw queries the /games endpoint for the available games and
// returns the raw response body as a string.
func (yf *YFantasy) getAvailableGamesRaw() (string, error) {
	return yf.get(fmt.Sprintf("games;is_available=1"))
}

// getLeagueRaw queries the /league endpoint for league information and returns
// the raw response body as a string.
// leagueKey has the format: <game_key>.l.<league_id> (ex: nba.l.1234)
func (yf *YFantasy) getLeagueRaw(leagueKey string) (string, error) {
	return yf.get(fmt.Sprintf("league/%s", leagueKey))
}

// getLeagueSettingsRaw queries the /league//settings endpoint for the league
// settings and returns the raw response body as a string.
func (yf *YFantasy) getLeagueSettingsRaw(leagueKey string) (string, error) {
	return yf.get(fmt.Sprintf("league/%s/settings", leagueKey))
}

// getLeagueStandingsRaw queries the /league//standings endpoint for the league
// standings and returns the raw response body as a string.
func (yf *YFantasy) getLeagueStandingsRaw(leagueKey string) (string, error) {
	return yf.get(fmt.Sprintf("league/%s/standings", leagueKey))
}

// getLeagueScoreboardRaw queries the /league//scoreboard endpoint for the
// league scoreboard (i.e. matchup data) for a given week and returns the
// raw response body as a string. Passing in 0 for the week, will return the
// scorebaord of the current week.
// TODO(famendola1): Use startWeek and numWeeks instead of week.
func (yf *YFantasy) getLeagueScoreboardRaw(leagueKey string, week int) (string, error) {
	if week == 0 {
		return yf.get(fmt.Sprintf("league/%s/scoreboard", leagueKey))
	}
	return yf.get(fmt.Sprintf("league/%s/scoreboard;week=%d", leagueKey, week))

}

// getTeamRaw queries the /team endpoint for team information and returns the
// response body as a string.
// teamKey has the format: <game_key>.l.<league_id>.t.<team_id>
func (yf *YFantasy) getTeamRaw(teamKey string) (string, error) {
	return yf.get(fmt.Sprintf("team/%s", teamKey))
}

// getTeamMatchupsRaw queries the /team//matchups endpoint for matchup data for
// a team for weeks startWeek to startWeek+numWeeks-1.
func (yf *YFantasy) getTeamMatchupsRaw(teamKey string, startWeek int, numWeeks int) (string, error) {
	weeks := make([]string, numWeeks)
	for i := 0; i < numWeeks; i++ {
		weeks[i] = strconv.Itoa(startWeek + i)
	}
	return yf.get(fmt.Sprintf("team/%s/matchups;weeks=%s", teamKey, strings.Join(weeks, ",")))
}

// getTeamStatsRaw queries the /team//stats endpoint for team stats of the given
// duration.
func (yf *YFantasy) getTeamStatsRaw(teamKey string, duration StatDuration) (string, error) {
	if duration.DurationType == "lastweek" || duration.DurationType == "lastmonth" {
		return yf.get(fmt.Sprintf("team/%s/stats;type=%s", teamKey, duration.DurationType))
	}

	if duration.DurationType == "date" {
		date := duration.Date
		if date == "" {
			date = time.Now().Format("2006-01-02")
		}
		return yf.get(fmt.Sprintf("team/%s/stats;type=%s;date=%s", teamKey, duration.DurationType, date))
	}

	if duration.DurationType == "season" || duration.DurationType == "average_season" {
		if duration.Season == "" {
			return yf.get(fmt.Sprintf("team/%s/stats;type=%s", teamKey, duration.DurationType))
		}
		return yf.get(fmt.Sprintf("team/%s/stats;type=%s;season=%s", teamKey, duration.DurationType, duration.Season))
	}

	return "", fmt.Errorf("requested duration invalid or not supported")
}

// getTeamRosterRaw queries the /team//roster endpoint for the team's current
// day roster and returns the raw response as a string.
func (yf *YFantasy) getTeamRosterRaw(teamKey string) (string, error) {
	return yf.get(fmt.Sprintf("team/%s/roster", teamKey))
}

// getTeamRosterWeekRaw queries the /team//roster endpoint for the team's roster
// for the given week and returns the raw response as a string.
func (yf *YFantasy) getTeamRosterWeekRaw(teamKey string, weekNum int) (string, error) {
	return yf.get(fmt.Sprintf("team/%s/roster;week=%d", teamKey, weekNum))
}

// getTeamRosterDayRaw queries the /team//roster endpoint for the team's roster
// for the given day and returns the raw response as a string.
// day is formatted as: yyyy-mm-dd
func (yf *YFantasy) getTeamRosterDayRaw(teamKey string, day string) (string, error) {
	return yf.get(fmt.Sprintf("team/%s/roster;date=%s", teamKey, day))
}

// getPlayerRaw queries the /league//players endpoint for player data and
// returns the raw response as a string.
// playerKey is formatted as: <game_key>.p.<player_id>
func (yf *YFantasy) getPlayerRaw(leageuKey string, playerKey string) (string, error) {
	return yf.get(
		fmt.Sprintf("league/%s/players;player_keys=%s", leageuKey, playerKey))
}

// getPlayersRaw queries the /league//players endpoint for player data for
// multiple players and returns the raw response as a string.
func (yf *YFantasy) getPlayersRaw(leagueKey string, playerKeys []string) (string, error) {
	return yf.get(fmt.Sprintf("league/%s/players;player_keys=%s", leagueKey, strings.Join(playerKeys, ",")))
}

// getPlayersBySearchRaw queries the /league//players endpoint with the "search"
// filter set to the provided search string.
func (yf *YFantasy) getPlayersBySearchRaw(leagueKey string, searchStr string) (string, error) {
	return yf.get(fmt.Sprintf("league/%s/players;search=%s", leagueKey, url.QueryEscape(searchStr)))
}

// getPlayersStatsRaw queries the /league//players//stats endpoint for stats of
// all the requested players of the given duration. Only 25 players can be
// requested at once.
func (yf *YFantasy) getPlayersStatsRaw(leagueKey string, playerKeys []string, duration StatDuration) (string, error) {
	players := strings.Join(playerKeys, ",")
	if duration.DurationType == "lastweek" || duration.DurationType == "lastmonth" {
		return yf.get(fmt.Sprintf("league/%s/players;player_keys=%s/stats;type=%s", leagueKey, players, duration.DurationType))
	}

	if duration.DurationType == "date" {
		date := duration.Date
		if date == "" {
			date = time.Now().Format("2006-01-02")
		}
		return yf.get(fmt.Sprintf("league/%s/players;player_keys=%s/stats;type=%s;date=%s", leagueKey, players, duration.DurationType, date))
	}

	if duration.DurationType == "season" || duration.DurationType == "average_season" {
		if duration.Season == "" {
			return yf.get(fmt.Sprintf("league/%s/players;player_keys=%s/stats;type=%s", leagueKey, players, duration.DurationType))
		}
		return yf.get(fmt.Sprintf("league/%s/players;player_keys=%s/stats;type=%s;season=%s", leagueKey, players, duration.DurationType, duration.Season))
	}

	return "", fmt.Errorf("requested duration invalid or not supported")
}

// TODO(famendola1): Query /league//players using filters

// getTransactionsRaw queries the /league//transactions endpoint for league
// transactions of the given types and returns the raw response as a string.
// Valid transactionTypes are: add, drop, commish, trade
func (yf *YFantasy) getTransactionsRaw(leagueKey string, transactionTypes []string) (string, error) {
	return yf.get(fmt.Sprintf("league/%s/transactions;types=%s", leagueKey, strings.Join(transactionTypes, ",")))
}

// getTeamTransactionsRaw queries the /league//transactions endpoint for league
// transactions of the given type for the given team and returns the raw
// response as a string.
// Valid transactionTypes are: pending_trade, waiver
func (yf *YFantasy) getTeamTransactionsRaw(leagueKey string, teamKey string, transactionType string) (string, error) {
	return yf.get(fmt.Sprintf("league/%s/transactions;team_key=%s;type=%s",
		leagueKey, teamKey, transactionType))
}

// TODO(famendola1): Implement more filters for transactions.

// getUserTeams queries the /users endpoint to get all the teams for the logged
// in user and returns the raw response as a string.
func (yf *YFantasy) getUserTeams() (string, error) {
	return yf.get(fmt.Sprintf("users;use_login=1/games/teams"))
}

// getUserTeamsForSport is the same as getUserTeams except the teams
// restricted to the given sport and only active teams are returned.
func (yf *YFantasy) getUserTeamsForSport(sport string) (string, error) {
	return yf.get(fmt.Sprintf("users;use_login=1/games;game_keys=%s/teams", sport))
}

// getUserLeagues queries the /users endpoint to get all the leagues for the
// logged in user and returns the raw response as a string.
func (yf *YFantasy) getUserLeagues() (string, error) {
	return yf.get(fmt.Sprintf("users;use_login=1/games/leagues"))
}

// getUserLeaguesForSport is the same as getUserLeagues except the leagues
// restricted to the given sport and only active leagues are returned.
func (yf *YFantasy) getUserLeaguesForSport(sport string) (string, error) {
	return yf.get(fmt.Sprintf("users;use_login=1/games;game_keys=%s/leagues", sport))
}

// getUserTeamInLeagueRaw queries the /users endpoint for the user's team in the
// given league.
func (yf *YFantasy) getUserTeamInLeagueRaw(gameKey string, leagueKey string) (string, error) {
	return yf.get(fmt.Sprintf("users;use_login=1/games;game_keys=%s/leagues;league_keys=%s/teams", gameKey, leagueKey))
}

// postAddDropTransaction sends a POST request to the /league//transactions
// endpoint to add and drop the selected player for the given team. The raw
// response is returned as a string.
func (yf *YFantasy) postAddDropTransaction(leagueKey string, teamKey string, addPlayerKey string, dropPlayerKey string) error {
	data := fmt.Sprintf(addDropTransaction, addPlayerKey, teamKey, dropPlayerKey, teamKey)
	return yf.post(fmt.Sprintf("league/%s/transactions", leagueKey), data)
}

// postAddTransaction sends a POST request to the /league//transactions
// endpoint to add the selected player for the given team. The raw response is
// returned as a string.
func (yf *YFantasy) postAddTransaction(leagueKey string, teamKey string, addPlayerKey string) error {
	data := fmt.Sprintf(addTransaction, addPlayerKey, teamKey)
	return yf.post(fmt.Sprintf("league/%s/transactions", leagueKey), data)
}

// postDropTransaction sends a POST request to the /league//transactions
// endpoint to drop the selected player for the given team. The raw response is
// returned as a string.
func (yf *YFantasy) postDropTransaction(leagueKey string, teamKey string, dropPlayerKey string) error {
	data := fmt.Sprintf(dropTransaction, teamKey, dropPlayerKey)
	return yf.post(fmt.Sprintf("league/%s/transactions", leagueKey), data)
}
