package yfantasy

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewPlayer(t *testing.T) {
	want := &Player{nil, "123.p.789"}
	got := NewPlayer(nil, "123.p.789")

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(Player{})) {
		t.Errorf("New() = %+v, want %+v", *got, *want)
	}
}
