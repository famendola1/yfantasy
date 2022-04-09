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

// NewMatchupsFromXML creates a new Matchups object parsed from an XML string.
func NewMatchupsFromXML(rawXML string, yf *YFantasy) (*Matchups, error) {
	var m Matchups
	if err := xml.NewDecoder(strings.NewReader(rawXML)).Decode(&m); err != nil {
		return nil, err
	}

	for i := range m.Matchup {
		for j := range m.Matchup[i].Teams.Team {
			m.Matchup[i].Teams.Team[j].yf = yf
		}
	}
	return &m, nil
}

// NewMatchupFromXML creates a new Matchup object parsed from an XML string.
func NewMatchupFromXML(rawXML string, yf *YFantasy) (*Matchup, error) {
	var m Matchup
	if err := xml.NewDecoder(strings.NewReader(rawXML)).Decode(&m); err != nil {
		return nil, err
	}

	for i := range m.Teams.Team {
		m.Teams.Team[i].yf = yf
	}
	return &m, nil
}
