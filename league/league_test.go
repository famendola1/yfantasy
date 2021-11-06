package league

import (
	"reflect"
	"testing"

	"github.com/famendola1/yfantasy/player"
	"github.com/famendola1/yfantasy/team"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNew(t *testing.T) {
	want := &League{nil, "789.l.456"}
	got := New(nil, "789.l.456")

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(*got, *want)) {
		t.Errorf("unexpected league: got %+v, want %+v", *got, *want)
	}
}

func TestExtractTeams(t *testing.T) {
	lg := New(nil, "223.l.431")
	want := []*team.Team{
		team.New(nil, "223.l.431.t.10", "223.l.431"),
		team.New(nil, "223.l.431.t.5", "223.l.431"),
		team.New(nil, "223.l.431.t.8", "223.l.431"),
		team.New(nil, "223.l.431.t.12", "223.l.431"),
	}

	got, err := lg.extractTeamsFromStandings(standingsResp)
	if err != nil {
		t.Errorf("extractTeams failed, expected success")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected teams extracted: got %v, want %v", got, want)
	}
}

func TestExtractPlayers(t *testing.T) {
	lg := New(nil, "123.1.456")
	want := []*player.Player{player.New(nil, "410.p.6513")}
	got, err := lg.extractPlayersFromSearch(searchResp)
	if err != nil {
		t.Errorf("extractPlayersFromSearch failed, expected success")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected players extracted: got %v, want %v", got, want)
	}
}
