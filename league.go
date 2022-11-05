package yfantasy

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
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
	StartWeek             string `xml:"start_week"`
	StartDate             string `xml:"start_date"`
	EndWeek               string `xml:"end_week"`
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

// NewLeagueFromXML returns a new League object parsed from an XML string.
func (yf *YFantasy) newLeagueFromXML(rawXML string) (*League, error) {
	var lg League
	err := xml.NewDecoder(strings.NewReader(rawXML)).Decode(&lg)
	if err != nil {
		return nil, err
	}
	lg.yf = yf
	return &lg, nil
}

// NewLeague creates a League containing all the league data from Yahoo.
func (yf *YFantasy) newLeague(lgKey string) *League {
	lg := &League{LeagueKey: lgKey, yf: yf}
	yf.fetchLeagueData(lg)
	return lg
}

// fetchLeagueData gets all the data for a league and populates all the fields.
func (yf *YFantasy) fetchLeagueData(lg *League) error {
	if !yf.IsValid() {
		return fmt.Errorf("unable to fetch league data, YFantasy is invalid")
	}
	rawResp, err := yf.getLeagueRaw(lg.LeagueKey)
	if err != nil {
		return err
	}

	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return err
	}

	node, err := xmlquery.Query(doc, "//league")
	if err != nil {
		return err
	}

	lg, err = yf.newLeagueFromXML(node.OutputXML(true))
	if err != nil {
		return err
	}
	return nil
}

// GameKey returns the game key for the league.
func (l *League) GameKey() string {
	return strings.Split(l.LeagueKey, ".l.")[0]
}

// Teams returns a list of the teams in the league
func (l *League) Teams() ([]*Team, error) {
	rawResp, err := l.yf.getLeagueStandingsRaw(l.LeagueKey)
	if err != nil {
		return nil, err
	}

	return l.extractTeams(rawResp)
}

// MyTeam returns the team that the user has in this league.
func (l *League) MyTeam() (*Team, error) {
	rawResp, err := l.yf.getUserTeamInLeagueRaw(l.GameKey(), l.LeagueKey)
	if err != nil {
		return nil, err
	}

	teams, err := l.extractTeams(rawResp)
	if err != nil {
		return nil, err
	}

	if len(teams) == 0 {
		return nil, fmt.Errorf("user has no teams in this league")
	}
	return teams[0], nil
}

// extractTeams parses the raw XML response from the
// /league//standings endpoint for teams.
func (l *League) extractTeams(rawResp string) ([]*Team, error) {
	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	nodes, err := xmlquery.QueryAll(doc, "//team")
	if err != nil {
		return nil, err
	}

	teams := make([]*Team, len(nodes))
	for i, node := range nodes {
		teams[i], err = l.yf.newTeamFromXML(node.OutputXML(true))
		if err != nil {
			return nil, err
		}
	}

	return teams, nil
}

// SearchPlayers searches for players using the provided name.
// playerName can be the player's full name or a partial name.
func (l *League) SearchPlayers(playerName string) ([]*Player, error) {
	rawResp, err := l.yf.getPlayersBySearchRaw(l.LeagueKey, playerName)
	if err != nil {
		return nil, err
	}

	return l.extractPlayersFromSearch(rawResp)
}

// extractPlayersFromSearch extracts players from the search results.
func (l *League) extractPlayersFromSearch(rawResp string) ([]*Player, error) {
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
		players[i], err = l.yf.newPlayerFromXML(node.OutputXML(true))
		if err != nil {
			return nil, err
		}
	}
	return players, nil
}

// Transactions returns all the league's transaction for the given types.
func (l *League) Transactions(transactionTypes []string) ([]*Transaction, error) {
	rawResp, err := l.yf.getTransactionsRaw(l.LeagueKey, transactionTypes)
	if err != nil {
		return nil, err
	}

	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	nodes, err := xmlquery.QueryAll(doc, "//transaction")
	if err != nil {
		return nil, err
	}

	transactions := make([]*Transaction, len(nodes))
	for i, node := range nodes {
		transactions[i], err = l.yf.newTransactionFromXML(node.OutputXML(true))
		if err != nil {
			return nil, err
		}
	}

	return transactions, nil
}

// GetPlayersStats fetches the stats for the requested players for the requested
// duration.
func (l *League) GetPlayersStats(playerKeys []string, duration StatDuration) ([]*Player, error) {
	rawResp, err := l.yf.getPlayersStatsRaw(l.LeagueKey, playerKeys, duration)
	if err != nil {
		return nil, err
	}

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
		players[i], err = l.yf.newPlayerFromXML(node.OutputXML(true))
		if err != nil {
			return nil, err
		}
	}
	return players, nil
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

	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	node, err := xmlquery.Query(doc, "//matchups")
	if err != nil {
		return nil, err
	}

	matchups, err := l.yf.newMatchupsFromXML(node.OutputXML(true))
	if err != nil {
		return nil, err
	}

	return matchups, nil
}

// GetStandings fetches the standings for a league.
func (l *League) GetStandings() (*Standings, error) {
	rawResp, err := l.yf.getLeagueStandingsRaw(l.LeagueKey)
	if err != nil {
		return nil, err
	}

	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return nil, err
	}

	node, err := xmlquery.Query(doc, "//standings")
	if err != nil {
		return nil, err
	}

	var s Standings
	if err := xml.NewDecoder(strings.NewReader(node.OutputXML(true))).Decode(&s); err != nil {
		return nil, err
	}

	return &s, nil
}
