package team

import (
	"reflect"
	"testing"

	"github.com/famendola1/yfantasy/player"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNew(t *testing.T) {
	want := &Team{nil, "123.l.456.t.789", "123.1.456"}
	got := New(nil, "123.l.456.t.789", "123.1.456")

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(*got, *want)) {
		t.Errorf("New() %+v, want %+v", *got, *want)
	}
}

func TestExtractPlayers(t *testing.T) {
	team := New(nil, "123.1.456.t.789", "123.1.456")
	want := []*player.Player{
		player.New(nil, "253.p.7569"),
		player.New(nil, "253.p.7054"),
		player.New(nil, "253.p.7382"),
	}
	got, err := team.extractPlayersFromRoster(rosterResp)
	if err != nil {
		t.Errorf("extractPlayersFromRoster(%q) failed, want success", rosterResp)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractPlayersFromRoster(%q) = %+v, want %+v", rosterResp, got, want)
	}
}
