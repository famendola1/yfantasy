package yfantasy

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewLeague(t *testing.T) {
	want := &League{XMLName: xml.Name{Local: "league"}, LeagueKey: "789.l.456"}
	got := NewLeague("789.l.456", nil)

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(League{})) {
		t.Errorf("New() = %+v, want %+v", *got, *want)
	}
}

func TestNewLeagueFromXML(t *testing.T) {
	want := &League{
		XMLName:               xml.Name{Local: "league"},
		LeagueKey:             "410.l.16883",
		LeagueID:              "16883",
		Name:                  "NBA Fantasy 2K22",
		URL:                   "https://basketball.fantasysports.yahoo.com/nba/16883",
		LogoURL:               "https://yahoofantasysports-res.cloudinary.com/image/upload/t_s192sq/fantasy-logos/0743c1232a845a00b408b165a683c19ad0ee273236a45766137c9698234246bf.jpg",
		DraftStatus:           "postdraft",
		NumTeams:              "12",
		EditKey:               "2021-11-09",
		WeeklyDeadline:        "intraday",
		LeagueUpdateTimestamp: "1636441834",
		ScoringType:           "headone",
		LeagueType:            "private",
		Renew:                 "402_22765",
		ShortInvitationURL:    "https://basketball.fantasysports.yahoo.com/nba/16883/invitation?key=1f14749a8282d491&ikey=bc6fb9e93fd791bb",
		AllowAddToDlExtraPos:  "1",
		IsProLeague:           "0",
		IsCashLeague:          "0",
		CurrentWeek:           "4",
		StartWeek:             "1",
		StartDate:             "2021-10-19",
		EndWeek:               "24",
		EndDate:               "2022-04-10",
		GameCode:              "nba",
		Season:                "2021"}
	got, err := NewLeagueFromXML(leagueFullTestResp, nil)
	if err != nil {
		t.Errorf("NewLeagueFromXML(%q) failed, want success.", leagueFullTestResp)
		return
	}

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(League{})) {
		t.Errorf("NewLeagueFromXML(%q) = %+v, want %+v", leagueFullTestResp, *got, *want)
	}
}

func TestExtractTeams(t *testing.T) {
	lg := NewLeague("223.l.431", nil)
	want := []*Team{
		NewTeam("223.l.431.t.10", nil),
		NewTeam("223.l.431.t.5", nil),
		NewTeam("223.l.431.t.8", nil),
		NewTeam("223.l.431.t.12", nil),
	}

	got, err := lg.extractTeams(standingsResp)
	if err != nil {
		t.Errorf("extractTeams failed, expected success")
		return
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractTeamsFromStandings(%q) = %v, want %v", standingsResp, got, want)
	}
}

func TestExtractPlayers(t *testing.T) {
	lg := NewLeague("123.1.456", nil)
	want := []*Player{NewPlayer(nil, "410.p.6513")}
	got, err := lg.extractPlayersFromSearch(searchResp)
	if err != nil {
		t.Errorf("extractPlayersFromSearch(%q) failed, want success", searchResp)
		return
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractPlayersFromSearch(%q) = %+v, want %+v", searchResp, got, want)
	}
}
