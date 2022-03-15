package yfantasy

import (
	"encoding/xml"
	"strings"
)

// Player represents a Yahoo player.
type Player struct {
	XMLName                  xml.Name          `xml:"player"`
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
	PlayerNotesLastTimestamp string            `xml:"player_notes_last_timestamp"`
	TransactionData          TransactionData   `xml:"transaction_data"`
	PlayerStats              PlayerStats       `xml:"player_stats"`

	yf *YFantasy
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
	XMLName      xml.Name `xml:"player_stats"`
	CoverageType string   `xml:"coverage_type"`
	Season       string   `xml:"season"`
	Date         string   `xml:"date"`
	Stats        Stats    `xml:"stats"`
}

// Players is a list of players.
type Players struct {
	Count  string    `xml:"count,attr"`
	Player []*Player `xml:"player"`
}

// NewPlayer returns a new player.
func NewPlayer(playerKey string, yf *YFantasy) *Player {
	return &Player{PlayerKey: playerKey, yf: yf}
}

// NewPlayerFromXML returns a new Player object parsed from an XML string.
func NewPlayerFromXML(rawXML string, yf *YFantasy) (*Player, error) {
	var p Player
	err := xml.NewDecoder(strings.NewReader(rawXML)).Decode(&p)
	if err != nil {
		return nil, err
	}
	p.yf = yf
	return &p, nil
}
