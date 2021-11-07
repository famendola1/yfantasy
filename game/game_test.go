package game

import (
	"reflect"
	"testing"

	"github.com/famendola1/yfantasy/league"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNew(t *testing.T) {
	want := &Game{yf: nil, Sport: "nba"}
	got := New(nil, "nba")

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(Game{})) {
		t.Errorf("New() = %+v, want %+v", *got, *want)
	}
}

func TestExtractGameId(t *testing.T) {
	game := New(nil, "nba")
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
	game := New(nil, "nba")
	want := []*league.League{
		league.New(nil, "410.l.16883"),
		league.New(nil, "410.l.61777"),
		league.New(nil, "410.l.159928"),
	}
	got, err := game.extractLeagues(leagueTestResp)
	if err != nil {
		t.Errorf("extractLeagues(%q) failed, want success", leagueTestResp)
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
	game := New(nil, "nba")
	game.gameID = "410"

	leagueID := "1234"
	want := league.New(nil, "410.l.1234")
	got, err := game.MakeLeague(leagueID)
	if err != nil {
		t.Errorf("MakeLeague(%q) failed, want success", leagueID)
	}

	if !cmp.Equal(got, want, cmp.AllowUnexported(Game{})) {
		t.Errorf("MakeLeague(%q) = %+v, want %+v", leagueID, *got, *want)
	}
}
