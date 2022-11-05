package yfantasy

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
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

type TeamStandings struct {
	Rank                    int                     `xml:"rank"`
	OutcomeTotals           OutcomeTotals           `xml:"outcome_totals"`
	DivisionalOutcomeTotals DivisionalOutcomeTotals `xml:"divisional_outcome_totals"`
}

type OutcomeTotals struct {
	Wins       string `xml:"wins"`
	Losses     string `xml:"losses"`
	Ties       string `xml:"ties"`
	Percentage string `xml:"percentage"`
}

type DivisionalOutcomeTotals struct {
	Wins   string `xml:"wins"`
	Losses string `xml:"losses"`
	Ties   string `xml:"ties"`
}

// newTeamFromXML returns a new Team object parsed from an XML string.
func (yf *YFantasy) newTeamFromXML(rawXML string) (*Team, error) {
	var tm Team
	err := xml.NewDecoder(strings.NewReader(rawXML)).Decode(&tm)
	if err != nil {
		return nil, err
	}
	tm.yf = yf
	return &tm, nil
}

// newTeam returns a new Team populated with information from Yahoo.
func (yf *YFantasy) newTeam(teamKey string) *Team {
	tm := &Team{TeamKey: teamKey, yf: yf}
	yf.fetchTeamData(tm)
	return tm
}

// FetchTeamData gets all the data for a team and populates all the fields.
func (yf *YFantasy) fetchTeamData(tm *Team) error {
	if !yf.IsValid() {
		return fmt.Errorf("unable to fetch team data, YFantasy is invalid")
	}
	rawResp, err := yf.getTeamRaw(tm.TeamKey)
	if err != nil {
		return err
	}

	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return err
	}

	node, err := xmlquery.Query(doc, "//team")
	if err != nil {
		return err
	}

	tm, err = yf.newTeamFromXML(node.OutputXML(true))
	if err != nil {
		return err
	}
	return nil
}

// Roster returns the list of players on this team.
func (t *Team) Roster() ([]*Player, error) {
	rawResp, err := t.yf.getTeamRosterRaw(t.TeamKey)
	if err != nil {
		return nil, err
	}
	return t.extractPlayersFromRoster(rawResp)
}

// extractPlayersFromRoster parses the raw XML response from the /team//roster
// endpoint for players.
func (t *Team) extractPlayersFromRoster(rawResp string) ([]*Player, error) {
	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	nodes, err := xmlquery.QueryAll(doc, "//player")
	if err != nil {
		return nil, err
	}

	players := make([]*Player, len(nodes))
	for i, node := range nodes {
		playerKey, err := xmlquery.Query(node, "/player_key")
		if err != nil {
			return nil, err
		}
		players[i] = t.yf.newPlayer(playerKey.InnerText())
	}

	return players, nil
}

// LeagueKey returns the key of the leauge this team is in.
func (t *Team) LeagueKey() string {
	return strings.Split(t.TeamKey, ".t.")[0]
}

// AddDrop adds the specified player to the team and drops the specified player
// from the team in a single transaction.
func (t *Team) AddDrop(addPlayerKey string, dropPlayerKey string) error {
	return t.yf.postAddDropTransaction(t.LeagueKey(), t.TeamKey, addPlayerKey, dropPlayerKey)
}

// Add adds the specified player to the team.
func (t *Team) Add(addPlayerKey string) error {
	return t.yf.postAddTransaction(t.LeagueKey(), t.TeamKey, addPlayerKey)
}

// Drop adds drops the specified player from the team.
func (t *Team) Drop(dropPlayerKey string) error {
	return t.yf.postDropTransaction(t.LeagueKey(), t.TeamKey, dropPlayerKey)
}

// GetTeamStats returns the team's stats for a given duration.
func (t *Team) GetTeamStats(duration StatDuration) (*TeamStats, error) {
	rawResp, err := t.yf.getTeamStatsRaw(t.TeamKey, duration)
	if err != nil {
		return nil, err
	}

	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	node, err := xmlquery.Query(doc, "//team_stats")
	if err != nil {
		return nil, err
	}

	var stats TeamStats
	err = xml.NewDecoder(strings.NewReader(node.OutputXML(true))).Decode(&stats)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}
