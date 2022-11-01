package yfantasy

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewGame(t *testing.T) {
	yf := &YFantasy{}
	want := &Game{Sport: "nba"}
	got := yf.Game("nba")

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(Game{})) {
		t.Errorf("New() = %+v, want %+v", *got, *want)
	}
}

func TestExtractGameId(t *testing.T) {
	yf := &YFantasy{}
	game := yf.Game("nba")
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
	yf := &YFantasy{}
	game := yf.Game("nba")
	want := []*League{
		yf.newLeague("410.l.16883"),
		yf.newLeague("410.l.61777"),
		yf.newLeague("410.l.159928"),
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
	yf := &YFantasy{}
	game := yf.Game("nba")
	game.gameID = "410"

	leagueID := "1234"
	want := yf.newLeague("410.l.1234")
	got, err := game.League(leagueID)
	if err != nil {
		t.Errorf("League(%q) failed, want success", leagueID)
	}

	if !cmp.Equal(got, want, cmp.AllowUnexported(League{}, YFantasy{})) {
		t.Errorf("League(%q) = %+v, want %+v", leagueID, *got, *want)
	}
}
