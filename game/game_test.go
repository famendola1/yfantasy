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

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(*got, *want)) {
		t.Errorf("unexpected Game built: got %+v, want %+v", *got, *want)
	}
}

func TestExtractGameId(t *testing.T) {
	want := "257"
	got, err := extractGameID(gameTestResp)
	if err != nil {
		t.Errorf("extractGameID failed, expected success")
	}

	if got != want {
		t.Errorf("unexpected game id: got %q, want %q", got, want)
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
		t.Errorf("extractLeagues failed, expected success")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected leagues extracted: got %v, want %v", got, want)
	}
}

func TestLeagueKeys(t *testing.T) {
	want := []string{"410.l.16883", "410.l.61777", "410.l.159928"}
	got, err := extractLeagueKeys(leagueTestResp)
	if err != nil {
		t.Errorf("extractLeagueKeys failed, expected success")
	}

	if !cmp.Equal(got, want) {
		t.Errorf("unexpected league keys: got %v, want %v", got, want)
	}
}

func TestMakeLeague(t *testing.T) {
	game := New(nil, "nba")

	want := league.New(nil, "nba.l.1234")
	got := game.MakeLeague("1234")

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(*got, *want)) {
		t.Errorf("unexpected League built: got %+v, want %+v", *got, *want)
	}
}
