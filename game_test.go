package yfantasy

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewGame(t *testing.T) {
	want := &Game{yf: nil, Sport: "nba"}
	got := NewGame("nba", nil)

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(Game{})) {
		t.Errorf("New() = %+v, want %+v", *got, *want)
	}
}

func TestExtractGameId(t *testing.T) {
	game := NewGame("nba", nil)
	want := "410"
	got, err := game.extractGameID(gameTestResp)
	if err != nil {
		t.Errorf("extractGameID(%q) failed, want success", gameTestResp)
	}

	if got != want {
		t.Errorf("extractGameID(%q) = %q, want %q", gameTestResp, got, want)
	}
}

func TestExtractLeagues(t *testing.T) {
	game := NewGame("nba", nil)
	want := []*League{
		NewLeague("410.l.16883", nil),
		NewLeague("410.l.61777", nil),
		NewLeague("410.l.159928", nil),
	}

	got, err := game.extractLeagues(leagueTestResp)
	if err != nil {
		fmt.Println(err)
		t.Errorf("extractLeagues(%q) failed, want success", leagueTestResp)
		return
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractLeagues(%q) = %+v, want %+v", leagueTestResp, got, want)
	}
}

func TestLeagueKeys(t *testing.T) {
	want := []string{"410.l.16883", "410.l.61777", "410.l.159928"}
	got, err := extractLeagueKeys(leagueTestResp)
	if err != nil {
		t.Errorf("extractLeagueKeys(%q) failed, want success", leagueTestResp)
	}

	if !cmp.Equal(got, want) {
		t.Errorf("extractLeagueKeys(%q) = %v, want %v", leagueTestResp, got, want)
	}
}

func TestMakeLeague(t *testing.T) {
	game := NewGame("nba", nil)
	game.gameID = "410"

	leagueID := "1234"
	want := NewLeague("410.l.1234", nil)
	got, err := game.MakeLeague(leagueID)
	if err != nil {
		t.Errorf("MakeLeague(%q) failed, want success", leagueID)
	}

	if !cmp.Equal(got, want, cmp.AllowUnexported(League{})) {
		t.Errorf("MakeLeague(%q) = %+v, want %+v", leagueID, *got, *want)
	}
}
