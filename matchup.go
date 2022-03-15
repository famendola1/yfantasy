package yfantasy

// TODO(famendola1): Add tests for this file.

import (
	"encoding/xml"
	"strings"
)

// Matchups hold multiple Matchup.
type Matchups struct {
	Matchup []Matchup `xml:"matchup"`
}

// Matchup represents a Yahoo matchup.
type Matchup struct {
	XMLName       xml.Name    `xml:"matchup"`
	Text          string      `xml:",chardata"`
	Week          string      `xml:"week"`
	WeekStart     string      `xml:"week_start"`
	WeekEnd       string      `xml:"week_end"`
	Status        string      `xml:"status"`
	IsPlayoffs    string      `xml:"is_playoffs"`
	IsConsolation string      `xml:"is_consolation"`
	StatWinners   StatWinners `xml:"stat_winners"`
	Teams         Teams       `xml:"teams"`
}

// StatWinners for Matchup
type StatWinners struct {
	StatWinner []StatWinner `xml:"stat_winner"`
}

// StatWinner represents the winner of a stat category.
type StatWinner struct {
	Text          string `xml:",chardata"`
	StatID        string `xml:"stat_id"`
	WinnerTeamKey string `xml:"winner_team_key"`
	IsTied        string `xml:"is_tied"`
}

// NewMatchupsFromXML creates a new Matchups object parsed from an XML string.
func NewMatchupsFromXML(rawXML string) (*Matchups, error) {
	var m Matchups
	if err := xml.NewDecoder(strings.NewReader(rawXML)).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

// NewMatchupFromXML creates a new Matchup object parsed from an XML string.
func NewMatchupFromXML(rawXML string) (*Matchup, error) {
	var m Matchup
	if err := xml.NewDecoder(strings.NewReader(rawXML)).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}
