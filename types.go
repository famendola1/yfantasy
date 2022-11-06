package yfantasy

// Player represents a Yahoo player.
type Player struct {
	PlayerKey                string            `xml:"player_key"`
	PlayerID                 int               `xml:"player_id"`
	Name                     Name              `xml:"name"`
	EditorialPlayerKey       string            `xml:"editorial_player_key"`
	EditorialTeamKey         string            `xml:"editorial_team_key"`
	EditorialTeamFullName    string            `xml:"editorial_team_full_name"`
	EditorialTeamAbbr        string            `xml:"editorial_team_abbr"`
	UniformNumber            int               `xml:"uniform_number"`
	DisplayPosition          string            `xml:"display_position"`
	Headshot                 Headshot          `xml:"headshot"`
	ImageURL                 string            `xml:"image_url"`
	IsUndroppable            bool              `xml:"is_undroppable"`
	PositionType             string            `xml:"position_type"`
	PrimaryPosition          string            `xml:"primary_position"`
	EligiblePositions        EligiblePositions `xml:"eligible_positions"`
	HasPlayerNotes           bool              `xml:"has_player_notes"`
	PlayerNotesLastTimestamp uint64            `xml:"player_notes_last_timestamp"`
	TransactionData          TransactionData   `xml:"transaction_data"`
	PlayerStats              PlayerStats       `xml:"player_stats"`
}

// Name for Player
type Name struct {
	Full       string `xml:"full"`
	First      string `xml:"first"`
	Last       string `xml:"last"`
	ASCIIFirst string `xml:"ascii_first"`
	ASCIILast  string `xml:"ascii_last"`
}

// Headshot for Player
type Headshot struct {
	URL  string `xml:"url"`
	Size string `xml:"size"`
}

// EligiblePositions for Player
type EligiblePositions struct {
	Position []string `xml:"position"`
}

// TransactionData for Player
type TransactionData struct {
	Type                string `xml:"type"`
	SourceType          string `xml:"source_type"`
	DestinationType     string `xml:"destination_type"`
	DestinationTeamKey  string `xml:"destination_team_key"`
	DestinationTeamName string `xml:"destination_team_name"`
	SourceTeamKey       string `xml:"source_team_key"`
	SourceTeamName      string `xml:"source_team_name"`
}

// PlayerStats for a Player.
type PlayerStats struct {
	CoverageType string `xml:"coverage_type"`
	Season       string `xml:"season"`
	Date         string `xml:"date"`
	Stats        Stats  `xml:"stats"`
}

// Players is a list of players.
type Players struct {
	Count  int      `xml:"count,attr"`
	Player []Player `xml:"player"`
}

// Transaction represents a Yahoo fantasy transaction.
type Transaction struct {
	TransactionKey string  `xml:"transaction_key"`
	TransactionID  int     `xml:"transaction_id"`
	Type           string  `xml:"type"`
	Status         string  `xml:"status"`
	Timestamp      uint64  `xml:"timestamp"`
	Players        Players `xml:"players"`
}

// Matchups hold multiple Matchup.
type Matchups struct {
	Matchup []Matchup `xml:"matchup"`
}

// Matchup represents a Yahoo matchup.
type Matchup struct {
	Week          int         `xml:"week"`
	WeekStart     string      `xml:"week_start"`
	WeekEnd       string      `xml:"week_end"`
	Status        string      `xml:"status"`
	IsPlayoffs    bool        `xml:"is_playoffs"`
	IsConsolation bool        `xml:"is_consolation"`
	IsTied        bool        `xml:"is_tied"`
	WinnerTeamKey string      `xml:"winner_team_key"`
	StatWinners   StatWinners `xml:"stat_winners"`
	Teams         Teams       `xml:"teams"`
}

// StatWinners for Matchup
type StatWinners struct {
	StatWinner []StatWinner `xml:"stat_winner"`
}

// StatWinner represents the winner of a stat category.
type StatWinner struct {
	StatID        int    `xml:"stat_id"`
	WinnerTeamKey string `xml:"winner_team_key"`
	IsTied        bool   `xml:"is_tied"`
}

// Stats holds multiple Stat
type Stats struct {
	Stat []Stat `xml:"stat"`
}

// Stat represents a stat category in Yahoo.
type Stat struct {
	StatID int    `xml:"stat_id"`
	Value  string `xml:"value"`
}
