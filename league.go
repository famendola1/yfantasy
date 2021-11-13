package yfantasy

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

// League represents a Yahoo league.
type League struct {
	XMLName               xml.Name `xml:"league"`
	LeagueKey             string   `xml:"league_key"`
	LeagueID              string   `xml:"league_id"`
	Name                  string   `xml:"name"`
	URL                   string   `xml:"url"`
	LogoURL               string   `xml:"logo_url"`
	DraftStatus           string   `xml:"draft_status"`
	NumTeams              string   `xml:"num_teams"`
	EditKey               string   `xml:"edit_key"`
	WeeklyDeadline        string   `xml:"weekly_deadline"`
	LeagueUpdateTimestamp string   `xml:"league_update_timestamp"`
	ScoringType           string   `xml:"scoring_type"`
	LeagueType            string   `xml:"league_type"`
	Renew                 string   `xml:"renew"`
	ShortInvitationURL    string   `xml:"short_invitation_url"`
	AllowAddToDlExtraPos  string   `xml:"allow_add_to_dl_extra_pos"`
	IsProLeague           string   `xml:"is_pro_league"`
	IsCashLeague          string   `xml:"is_cash_league"`
	CurrentWeek           string   `xml:"current_week"`
	StartWeek             string   `xml:"start_week"`
	StartDate             string   `xml:"start_date"`
	EndWeek               string   `xml:"end_week"`
	EndDate               string   `xml:"end_date"`
	GameCode              string   `xml:"game_code"`
	Season                string   `xml:"season"`

	yf *YFantasy
}

// NewLeagueFromXML returns a new League object parsed from an XML string.
func NewLeagueFromXML(rawXML string, yf *YFantasy) (*League, error) {
	var lg League
	err := xml.NewDecoder(strings.NewReader(rawXML)).Decode(&lg)
	if err != nil {
		return nil, err
	}
	lg.yf = yf
	return &lg, nil
}

// NewLeague creates a League with just the LeagueKey field set. To get all the
// league data, clients must call FetchLeagueData.
func NewLeague(lgKey string, yf *YFantasy) *League {
	return &League{XMLName: xml.Name{Local: "league"}, LeagueKey: lgKey, yf: yf}
}

// FetchLeagueData gets all the data for a league and populates all the fields.
func (l *League) FetchLeagueData() error {
	if l.yf == nil {
		return fmt.Errorf("unable to fetch league data, YFantasy is nil")
	}
	rawResp, err := l.yf.GetLeagueRaw(l.LeagueKey)
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

	l, err = NewLeagueFromXML(node.OutputXML(true), l.yf)
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
	rawResp, err := l.yf.GetLeagueStandingsRaw(l.LeagueKey)
	if err != nil {
		return nil, err
	}

	return l.extractTeams(rawResp)
}

// UserTeam returns the team that the user has in this league.
func (l *League) UserTeam() (*Team, error) {
	rawResp, err := l.yf.GetUserTeamInLeagueRaw(l.GameKey(), l.LeagueKey)
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
		teams[i], err = NewTeamFromXML(node.OutputXML(true), l.yf)
		if err != nil {
			return nil, err
		}
	}

	return teams, nil
}

// SearchPlayers searches for players using the provided name.
// playerName can be the player's full name or a partial name.
func (l *League) SearchPlayers(playerName string) ([]*Player, error) {
	rawResp, err := l.yf.GetPlayersBySearchRaw(l.LeagueKey, playerName)
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
		players[i], err = NewPlayerFromXML(node.OutputXML(true), l.yf)
		if err != nil {
			return nil, err
		}
	}
	return players, nil
}

// FetchPlayerData gets all the data for a player and populates/overrides all
// the fields in the player object.
func (l *League) FetchPlayerData(player *Player) error {
	rawResp, err := l.yf.GetPlayerRaw(l.LeagueKey, player.PlayerKey)
	if err != nil {
		return err
	}

	doc, err := xmlquery.Parse(strings.NewReader(rawResp))
	if err != nil {
		return err
	}

	node, err := xmlquery.Query(doc, "//player")
	if err != nil {
		return err
	}

	player, err = NewPlayerFromXML(node.OutputXML(true), l.yf)
	if err != nil {
		return err
	}
	return nil
}

// Transactions returns all the league's transaction for the given types.
func (l *League) Transactions(transactionTypes []string) ([]*Transaction, error) {
	rawResp, err := l.yf.GetTransactionsRaw(l.LeagueKey, transactionTypes)
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
		transactions[i], err = NewTransactionFromXML(node.OutputXML(true), l.yf)
		if err != nil {
			return nil, err
		}
	}

	return transactions, nil
}

// GetPlayersStats fetches the stats for the requested players for the requested
// duration.
func (l *League) GetPlayersStats(playerKeys []string, duration StatDuration) ([]*Player, error) {
	rawResp, err := l.yf.GetPlayersStatsRaw(l.LeagueKey, playerKeys, duration)
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
		players[i], err = NewPlayerFromXML(node.OutputXML(true), l.yf)
		if err != nil {
			return nil, err
		}
	}
	return players, nil
}
