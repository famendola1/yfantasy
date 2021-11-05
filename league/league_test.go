package league

import (
	"reflect"
	"testing"

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
	lg := New(nil, "12345")
	want := []*team.Team{
		team.New(nil, "223.l.431.t.10"),
		team.New(nil, "223.l.431.t.5"),
		team.New(nil, "223.l.431.t.8"),
		team.New(nil, "223.l.431.t.12"),
	}

	got, err := lg.extractTeamsFromStandings(standingsResp)
	if err != nil {
		t.Errorf("extractTeams failed, expected success")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected teams extracted: got %v, want %v", got, want)
	}

}
