package yfantasy

import (
	"fmt"
)

// League represents a Yahoo league.
type League struct {
	LeagueKey             string `xml:"league_key"`
	LeagueID              int    `xml:"league_id"`
	Name                  string `xml:"name"`
	URL                   string `xml:"url"`
	LogoURL               string `xml:"logo_url"`
	DraftStatus           string `xml:"draft_status"`
	NumTeams              int    `xml:"num_teams"`
	EditKey               string `xml:"edit_key"`
	WeeklyDeadline        string `xml:"weekly_deadline"`
	LeagueUpdateTimestamp string `xml:"league_update_timestamp"`
	ScoringType           string `xml:"scoring_type"`
	LeagueType            string `xml:"league_type"`
	Renew                 string `xml:"renew"`
	ShortInvitationURL    string `xml:"short_invitation_url"`
	AllowAddToDlExtraPos  string `xml:"allow_add_to_dl_extra_pos"`
	IsProLeague           bool   `xml:"is_pro_league"`
	IsCashLeague          bool   `xml:"is_cash_league"`
	CurrentWeek           int    `xml:"current_week"`
	StartWeek             int    `xml:"start_week"`
	StartDate             string `xml:"start_date"`
	EndWeek               int    `xml:"end_week"`
	EndDate               string `xml:"end_date"`
	GameCode              string `xml:"game_code"`
	Season                string `xml:"season"`
	IsFinished            bool   `xml:"is_finished"`

	yf *YFantasy
}

// Standings contains the standings for a league.
type Standings struct {
	Teams Teams `xml:"teams"`
}

// League creates a League containing all the league data from Yahoo.
// gameKey can either be the key or code (e.g. nba, nfl, nhl, etc) for a game.
func (yf *YFantasy) League(gameKey string, leagueID int) (*League, error) {
	rawResp, err := yf.getLeagueRaw(fmt.Sprintf("%s.l.%d", gameKey, leagueID))
	if err != nil {
		return nil, err
	}

	var lg League
	if err := parse(rawResp, "//league", &lg); err != nil {
		return nil, err
	}
	lg.yf = yf
	return &lg, nil
}

// UserLeagues returns all the leagues the current user is registered in for a
// given game. gameKey can either be the key or code (e.g. nba, nfl, nhl, etc)
// for a game.
func (yf *YFantasy) UserLeagues(gameKey string) ([]*League, error) {
	rawResp, err := yf.getUserLeaguesForSport(gameKey)
	if err != nil {
		return nil, err
	}
	return parseAllLeagues(rawResp, yf)
}

// FindLeagueByName returns a league with the given name for the given game. If
// no league is found an error is returned. gameKey can either be the key or
// code (e.g. nba, nfl, nhl, etc) for a game.
func (yf *YFantasy) FindLeagueByName(gameKey, lgName string) (*League, error) {
	lgs, err := yf.UserLeagues(gameKey)
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

// Teams returns a list of the teams in the league.
func (l *League) Teams() ([]*Team, error) {
	rawResp, err := l.yf.getLeagueStandingsRaw(l.LeagueKey)
	if err != nil {
		return nil, err
	}

	return parseAllTeams(rawResp, l.yf)
}

// UserTeam returns the team that the user has in this league.
func (l *League) UserTeam() (*Team, error) {
	rawResp, err := l.yf.getUserTeamInLeagueRaw(l.GameCode, l.LeagueKey)
	if err != nil {
		return nil, err
	}

	teams, err := parseAllTeams(rawResp, l.yf)
	if err != nil {
		return nil, err
	}

	if len(teams) == 0 {
		return nil, fmt.Errorf("user has no teams in this league")
	}
	return teams[0], nil
}

// SearchPlayers searches for players using the provided name.
// playerName can be the player's full name or a partial name.
func (l *League) SearchPlayers(playerName string) ([]*Player, error) {
	rawResp, err := l.yf.getPlayersBySearchRaw(l.LeagueKey, playerName)
	if err != nil {
		return nil, err
	}

	return parseAllPlayers(rawResp)
}

// Transactions returns all the league's transaction for the given types.
func (l *League) Transactions(transactionTypes []string) ([]*Transaction, error) {
	rawResp, err := l.yf.getTransactionsRaw(l.LeagueKey, transactionTypes)
	if err != nil {
		return nil, err
	}
	return parseAllTransactions(rawResp)
}

// GetPlayersStats fetches the stats for the requested players for the requested
// duration.
func (l *League) GetPlayersStats(playerKeys []string, duration StatDuration) ([]*Player, error) {
	rawResp, err := l.yf.getPlayersStatsRaw(l.LeagueKey, playerKeys, duration)
	if err != nil {
		return nil, err
	}
	return parseAllPlayers(rawResp)
}

// GetScoreboard fetches all the matchups in a league for the given week.
func (l *League) GetScoreboard(week int) (*Matchups, error) {
	if week < 1 {
		return nil, fmt.Errorf("invalid week number")
	}

	rawResp, err := l.yf.getLeagueScoreboardRaw(l.LeagueKey, week)
	if err != nil {
		return nil, err
	}

	var m Matchups
	if err := parse(rawResp, "//matchups", &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// GetStandings fetches the standings for a league.
func (l *League) GetStandings() (*Standings, error) {
	rawResp, err := l.yf.getLeagueStandingsRaw(l.LeagueKey)
	if err != nil {
		return nil, err
	}

	var s Standings
	if err := parse(rawResp, "//standings", &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// FindTeamByName searches the league for a Team with the given name.
func (l *League) FindTeamByName(teamName string) (*Team, error) {
	rawResp, err := l.yf.getLeagueTeamsRaw(l.LeagueKey)
	if err != nil {
		return nil, err
	}

	teams, err := parseAllTeams(rawResp, l.yf)
	for _, team := range teams {
		if team.Name == teamName {
			return team, nil
		}
	}
	return nil, fmt.Errorf("team with name: %q not found", teamName)
}

// FindTeamRosterByName searches the league for a Team with the given name and
// returns the roster.
func (l *League) FindTeamRosterByName(teamName string) (*Roster, error) {
	rawResp, err := l.yf.getLeagueTeamsRaw(l.LeagueKey)
	if err != nil {
		return nil, err
	}

	teams, err := parseAllTeams(rawResp, l.yf)
	for _, team := range teams {
		if team.Name == teamName {
			return &team.Roster, nil
		}
	}
	return nil, fmt.Errorf("team with name: %q not found", teamName)
}
