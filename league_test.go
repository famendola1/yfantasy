package yfantasy

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewLeague(t *testing.T) {
	want := &League{nil, "789.l.456"}
	got := NewLeague(nil, "789.l.456")

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(League{})) {
		t.Errorf("New() = %+v, want %+v", *got, *want)
	}
}

func TestExtractTeams(t *testing.T) {
	lg := NewLeague(nil, "223.l.431")
	want := []*Team{
		NewTeam(nil, "223.l.431.t.10"),
		NewTeam(nil, "223.l.431.t.5"),
		NewTeam(nil, "223.l.431.t.8"),
		NewTeam(nil, "223.l.431.t.12"),
	}

	got, err := lg.extractTeams(standingsResp)
	if err != nil {
		t.Errorf("extractTeams failed, expected success")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractTeamsFromStandings(%q) = %v, want %v", standingsResp, got, want)
	}
}

func TestExtractPlayers(t *testing.T) {
	lg := NewLeague(nil, "123.1.456")
	want := []*Player{NewPlayer(nil, "410.p.6513")}
	got, err := lg.extractPlayersFromSearch(searchResp)
	if err != nil {
		t.Errorf("extractPlayersFromSearch(%q) failed, want success", searchResp)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractPlayersFromSearch(%q) = %+v, want %+v", searchResp, got, want)
	}
}
