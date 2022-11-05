package yfantasy

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewLeague(t *testing.T) {
	yf := &YFantasy{}
	want := &League{LeagueKey: "789.l.456"}
	got := yf.newLeague("789.l.456")

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(League{})) {
		t.Errorf("New() = %+v, want %+v", *got, *want)
	}
}

func TestNewLeagueFromXML(t *testing.T) {
	yf := &YFantasy{}
	want := &League{
		LeagueKey:             "410.l.16883",
		LeagueID:              16883,
		Name:                  "NBA Fantasy 2K22",
		URL:                   "https://basketball.fantasysports.yahoo.com/nba/16883",
		LogoURL:               "https://yahoofantasysports-res.cloudinary.com/image/upload/t_s192sq/fantasy-logos/0743c1232a845a00b408b165a683c19ad0ee273236a45766137c9698234246bf.jpg",
		DraftStatus:           "postdraft",
		NumTeams:              12,
		EditKey:               "2021-11-09",
		WeeklyDeadline:        "intraday",
		LeagueUpdateTimestamp: "1636441834",
		ScoringType:           "headone",
		LeagueType:            "private",
		Renew:                 "402_22765",
		ShortInvitationURL:    "https://basketball.fantasysports.yahoo.com/nba/16883/invitation?key=1f14749a8282d491&ikey=bc6fb9e93fd791bb",
		AllowAddToDlExtraPos:  "1",
		IsProLeague:           false,
		IsCashLeague:          false,
		CurrentWeek:           4,
		StartWeek:             "1",
		StartDate:             "2021-10-19",
		EndWeek:               "24",
		EndDate:               "2022-04-10",
		GameCode:              "nba",
		Season:                "2021"}
	got, err := yf.newLeagueFromXML(leagueFullTestResp)
	if err != nil {
		t.Errorf("NewLeagueFromXML(%q) failed, want success.", leagueFullTestResp)
		return
	}

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(League{})) {
		t.Errorf("NewLeagueFromXML(%q) = %+v, want %+v", leagueFullTestResp, *got, *want)
	}
}

func TestExtractTeams(t *testing.T) {
	yf := &YFantasy{}
	lg := yf.newLeague("223.l.431")
	want := []*Team{
		yf.newTeam("223.l.431.t.10"),
		yf.newTeam("223.l.431.t.5"),
		yf.newTeam("223.l.431.t.8"),
		yf.newTeam("223.l.431.t.12"),
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
	yf := &YFantasy{}
	lg := yf.newLeague("123.1.456")
	want := []*Player{{PlayerKey: "410.p.6513", yf: yf}}
	got, err := lg.extractPlayersFromSearch(searchResp)
	if err != nil {
		t.Errorf("extractPlayersFromSearch(%q) failed, want success", searchResp)
		return
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractPlayersFromSearch(%q) = %+v, want %+v", searchResp, got, want)
	}
}
