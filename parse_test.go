package yfantasy

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestParseLeague(t *testing.T) {
	want := &League{LeagueKey: "410.l.16883"}
	got := &League{}

	if err := parse(leagueTestResp, "//league", got); err != nil {
		t.Errorf("parse(%s, \"//league\", %+v) failed, want success", leagueTestResp, got)
		return
	}

	if !cmp.Equal(got, want, cmp.AllowUnexported(League{}, YFantasy{})) {
		t.Errorf("got %+v, want %+v", *got, *want)
	}
}

func TestParseAllLeagues(t *testing.T) {
	yf := &YFantasy{}
	want := []*League{
		&League{LeagueKey: "410.l.16883"},
		&League{LeagueKey: "410.l.61777"},
		&League{LeagueKey: "410.l.159928"},
	}
	got, err := parseAllLeagues(leagueTestResp, yf)
	if err != nil {
		t.Errorf("parseAllLeagues(%s, %+v) failed, want success", leagueTestResp, yf)
		return
	}

	if diff := cmp.Diff(got, want, cmpopts.IgnoreFields(League{}, "yf")); diff != "" {
		t.Errorf("diff (-got, +want):\n %s", diff)
	}
}

func TestParseTeam(t *testing.T) {

}

func TestParseAllTeams(t *testing.T) {

}

func ParseAllString(t *testing.T) {
	want := []string{"410.l.16883", "410.l.61777", "410.l.159928"}
	got, err := parseAllString(leagueTestResp, "//league_key")
	if err != nil {
		t.Errorf("parseAllString(%s, \"//league_key\") failed, want success", leagueTestResp)
		return
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("diff (-got, +want):\n %s", diff)
	}
}
