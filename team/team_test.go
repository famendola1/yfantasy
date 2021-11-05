package team

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNew(t *testing.T) {
	want := &Team{nil, "123.l.456.t.789"}
	got := New(nil, "123.l.456.t.789")

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(*got, *want)) {
		t.Errorf("unexpected team: got %+v, want %+v", *got, *want)
	}
}
