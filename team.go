package yfantasy

import (
	"encoding/xml"
	"strings"

	"github.com/antchfx/xmlquery"
)

// Team represents a Yahoo team.
type Team struct {
	XMLName               xml.Name   `xml:"team"`
	TeamKey               string     `xml:"team_key"`
	TeamID                string     `xml:"team_id"`
	Name                  string     `xml:"name"`
	IsOwnedByCurrentLogin string     `xml:"is_owned_by_current_login"`
	URL                   string     `xml:"url"`
	TeamLogos             TeamLogos  `xml:"team_logos"`
	WaiverPriority        string     `xml:"waiver_priority"`
	NumberOfMoves         string     `xml:"number_of_moves"`
	NumberOfTrades        string     `xml:"number_of_trades"`
	RosterAdds            RosterAdds `xml:"roster_adds"`
	LeagueScoringType     string     `xml:"league_scoring_type"`
	DraftPosition         string     `xml:"draft_position"`
	HasDraftGrade         string     `xml:"has_draft_grade"`
	Managers              Managers   `xml:"managers"`

	yf *YFantasy
}

// TeamLogos for Team
type TeamLogos struct {
	TeamLogo []TeamLogo `xml:"team_logo"`
}

// TeamLogo for Team
type TeamLogo struct {
	Size string `xml:"size"`
	URL  string `xml:"url"`
}

// RosterAdds for Team
type RosterAdds struct {
	CoverageType  string `xml:"coverage_type"`
	CoverageValue string `xml:"coverage_value"`
	Value         string `xml:"value"`
}

// Managers for Team
type Managers struct {
	Manager []Manager `xml:"manager"`
}

// Manager for Team
type Manager struct {
	ManagerID      string `xml:"manager_id"`
	Nickname       string `xml:"nickname"`
	Guid           string `xml:"guid"`
	IsCommissioner string `xml:"is_commissioner"`
	IsCurrentLogin string `xml:"is_current_login"`
	Email          string `xml:"email"`
	ImageURL       string `xml:"image_url"`
	FeloScore      string `xml:"felo_score"`
	FeloTier       string `xml:"felo_tier"`
}

// NewTeamFromXML returns a new Team object parsed form an XML string.
func NewTeamFromXML(rawXML string, yf *YFantasy) (*Team, error) {
	var tm Team
	err := xml.NewDecoder(strings.NewReader(rawXML)).Decode(&tm)
	if err != nil {
		return nil, err
	}
	tm.yf = yf
	return &tm, nil
}

// NewTeam returns a new Team
func NewTeam(teamKey string, yf *YFantasy) *Team {
	return &Team{XMLName: xml.Name{Local: "team"}, TeamKey: teamKey, yf: yf}
}

// FetchTeamData gets all the data for a team and populates all the fields.
func (t *Team) FetchTeamData() error {
	rawResp, err := t.yf.GetTeamRaw(t.TeamKey)
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

	t, err = NewTeamFromXML(node.OutputXML(true), t.yf)
	if err != nil {
		return err
	}
	return nil
}

// Roster returns the list of players on this team.
func (t *Team) Roster() ([]*Player, error) {
	rawResp, err := t.yf.GetTeamRosterRaw(t.TeamKey)
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
		players[i] = NewPlayer(playerKey.InnerText(), t.yf)
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
	return t.yf.PostAddDropTransaction(t.LeagueKey(), t.TeamKey, addPlayerKey, dropPlayerKey)
}
