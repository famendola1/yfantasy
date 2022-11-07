package yfantasy

import (
	"fmt"
	"strings"
)

// Teams holds multiple Team.
type Teams struct {
	Count string `xml:"count,attr"`
	Team  []Team `xml:"team"`
}

// Team represents a Yahoo team.
type Team struct {
	TeamKey               string             `xml:"team_key"`
	TeamID                int                `xml:"team_id"`
	Name                  string             `xml:"name"`
	IsOwnedByCurrentLogin bool               `xml:"is_owned_by_current_login"`
	URL                   string             `xml:"url"`
	TeamLogos             TeamLogos          `xml:"team_logos"`
	WaiverPriority        int                `xml:"waiver_priority"`
	NumberOfMoves         int                `xml:"number_of_moves"`
	NumberOfTrades        int                `xml:"number_of_trades"`
	RosterAdds            RosterAdds         `xml:"roster_adds"`
	LeagueScoringType     string             `xml:"league_scoring_type"`
	DraftPosition         int                `xml:"draft_position"`
	HasDraftGrade         bool               `xml:"has_draft_grade"`
	Managers              Managers           `xml:"managers"`
	TeamStats             TeamStats          `xml:"team_stats"`
	TeamPoints            TeamPoints         `xml:"team_points"`
	TeamRemainingGames    TeamRemainingGames `xml:"team_remaining_games"`
	ClinchedPlayoffs      bool               `xml:"clinched_playoffs"`
	TeamStandings         TeamStandings      `xml:"team_standings"`
	Roster                Roster             `xml:"roster"`

	yf *YFantasy
}

// TeamLogos for Team.
type TeamLogos struct {
	TeamLogo []TeamLogo `xml:"team_logo"`
}

// TeamLogo for Team.
type TeamLogo struct {
	Size string `xml:"size"`
	URL  string `xml:"url"`
}

// RosterAdds for Team.
type RosterAdds struct {
	CoverageType  string `xml:"coverage_type"`
	CoverageValue int    `xml:"coverage_value"`
	Value         int    `xml:"value"`
}

// Managers for Team.
type Managers struct {
	Manager []Manager `xml:"manager"`
}

// Manager for Team.
type Manager struct {
	ManagerID      int    `xml:"manager_id"`
	Nickname       string `xml:"nickname"`
	GUID           string `xml:"guid"`
	IsCommissioner bool   `xml:"is_commissioner"`
	IsCurrentLogin bool   `xml:"is_current_login"`
	Email          string `xml:"email"`
	ImageURL       string `xml:"image_url"`
	FeloScore      int    `xml:"felo_score"`
	FeloTier       string `xml:"felo_tier"`
}

// TeamStats for Team.
type TeamStats struct {
	CoverageType string `xml:"coverage_type"`
	Week         int    `xml:"week"`
	Stats        Stats  `xml:"stats"`
}

// TeamPoints for Team.
type TeamPoints struct {
	CoverageType string `xml:"coverage_type"`
	Week         int    `xml:"week"`
	Total        int    `xml:"total"`
}

// TeamRemainingGames for Team.
type TeamRemainingGames struct {
	CoverageType string `xml:"coverage_type"`
	Week         int    `xml:"week"`
	Total        Total  `xml:"total"`
}

// Total for TeamRemainingGames.
type Total struct {
	RemainingGames int `xml:"remaining_games"`
	LiveGames      int `xml:"live_games"`
	CompletedGames int `xml:"completed_games"`
}

// TeamStandings contains information about a Team's ranking in their league.
type TeamStandings struct {
	Rank                    int                     `xml:"rank"`
	OutcomeTotals           OutcomeTotals           `xml:"outcome_totals"`
	DivisionalOutcomeTotals DivisionalOutcomeTotals `xml:"divisional_outcome_totals"`
}

// OutcomeTotals contains information on the outcomes of a Team's matchups.
type OutcomeTotals struct {
	Wins       int     `xml:"wins"`
	Losses     int     `xml:"losses"`
	Ties       int     `xml:"ties"`
	Percentage float32 `xml:"percentage"`
}

// DivisionalOutcomeTotals contains information on the outcomes of a Team's matchups in their division.
type DivisionalOutcomeTotals struct {
	Wins   int `xml:"wins"`
	Losses int `xml:"losses"`
	Ties   int `xml:"ties"`
}

// Roster contains information on a Team's roster.
type Roster struct {
	CoverageType string  `xml:"coverage_type"`
	Date         string  `xml:"date"`
	IsEditable   bool    `xml:"is_editable"`
	Players      Players `xml:"players"`
}

// Team returns a new Team populated with information from Yahoo.
func (yf *YFantasy) Team(gameKey string, leagueID, teamID int) (*Team, error) {
	var tm Team
	rawResp, err := yf.getTeamRaw(fmt.Sprintf("%s.l.%d.t.%d", gameKey, leagueID, teamID))
	if err != nil {
		return nil, err
	}
	if err := parse(rawResp, "//team", tm); err != nil {
		return nil, err
	}
	tm.yf = yf
	return &tm, nil
}

// GetRoster queries for and returns a Team's roster.
func (t *Team) GetRoster() (*Roster, error) {
	rawResp, err := t.yf.getTeamRosterRaw(t.TeamKey)
	if err != nil {
		return nil, err
	}

	var r Roster
	if err := parse(rawResp, "//roster", &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (t *Team) leagueKey() string {
	return strings.Split(t.TeamKey, ".t.")[0]
}

// AddDrop adds the specified player to the team and drops the specified player
// from the team in a single transaction.
func (t *Team) AddDrop(addPlayerKey string, dropPlayerKey string) error {
	return t.yf.postAddDropTransaction(t.leagueKey(), t.TeamKey, addPlayerKey, dropPlayerKey)
}

// Add adds the specified player to the team.
func (t *Team) Add(addPlayerKey string) error {
	return t.yf.postAddTransaction(t.leagueKey(), t.TeamKey, addPlayerKey)
}

// Drop adds drops the specified player from the team.
func (t *Team) Drop(dropPlayerKey string) error {
	return t.yf.postDropTransaction(t.leagueKey(), t.TeamKey, dropPlayerKey)
}

// GetTeamStats returns the team's stats for a given duration.
func (t *Team) GetTeamStats(duration StatDuration) (*TeamStats, error) {
	rawResp, err := t.yf.getTeamStatsRaw(t.TeamKey, duration)
	if err != nil {
		return nil, err
	}

	var stats TeamStats
	if err := parse(rawResp, "//team_stats", &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}
