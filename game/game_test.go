package game

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNew(t *testing.T) {
	want := &Game{yf: nil, Sport: "nba"}
	got := New(nil, "nba")

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(*got, *want)) {
		t.Errorf("unexpected Game built: got %+v, want %+v", got, want)
	}
}

func TestExtractGameId(t *testing.T) {
	testResp :=
		`<?xml version="1.0" encoding="UTF-8"?>
    <fantasy_content xml:lang="en-US" yahoo:uri="http://fantasysports.yahooapis.com/fantasy/v2/game/nfl" xmlns:yahoo="http://www.yahooapis.com/v1/base.rng" time="30.575037002563ms" copyright="Data provided by Yahoo! and STATS, LLC" xmlns="http://fantasysports.yahooapis.com/fantasy/v2/base.rng">
      <game>
        <game_key>257</game_key>
        <game_id>257</game_id>
        <name>Football</name>
        <code>nfl</code>
        <type>full</type>
        <url>https://football.fantasysports.yahoo.com/f1</url>
        <season>2011</season>
      </game>
    </fantasy_content>`

	want := "257"
	got, _ := extractGameID(testResp)

	if got != want {
		t.Errorf("unexpected game id: got %q, want %q", got, want)
	}
}
