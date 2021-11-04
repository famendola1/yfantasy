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
	testResp := `<?xml version="1.0" encoding="UTF-8"?>
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
	got, err := extractGameID(testResp)
	if err != nil {
		t.Errorf("extractGameID failed, expected success")
	}

	if got != want {
		t.Errorf("unexpected game id: got %q, want %q", got, want)
	}
}

func TestExtractLeagues(t *testing.T) {
	game := New(nil, "nba")
	testResp := `<?xml version="1.0" encoding="UTF-8"?>
  <fantasy_content xml:lang="en-US" yahoo:uri="http://fantasysports.yahooapis.com/fantasy/v2/users;use_login=1/games;game_keys=nba/leagues" time="37.668943405151ms" copyright="Data provided by Yahoo! and STATS, LLC" refresh_rate="60" xmlns:yahoo="http://www.yahooapis.com/v1/base.rng" xmlns="http://fantasysports.yahooapis.com/fantasy/v2/base.rng">
   <users count="1">
    <user>
     <guid>EKFDPDVSJIGZD64VAL6WYCIH2I</guid>
     <games count="1">
      <game>
       <game_key>410</game_key>
       <game_id>410</game_id>
       <name>Basketball</name>
       <code>nba</code>
       <type>full</type>
       <url>https://basketball.fantasysports.yahoo.com/nba</url>
       <season>2021</season>
       <is_registration_over>0</is_registration_over>
       <is_game_over>0</is_game_over>
       <is_offseason>0</is_offseason>
       <leagues count="3">
        <league>
         <league_key>410.l.16883</league_key>
        </league>
        <league>
         <league_key>410.l.61777</league_key>
        </league>
        <league>
         <league_key>410.l.159928</league_key>
        </league>
       </leagues>
      </game>
     </games>
    </user>
   </users>
  </fantasy_content>`

	want := []*league.League{
		league.New(nil, "410.l.16883"),
		league.New(nil, "410.l.61777"),
		league.New(nil, "410.l.159928"),
	}
	got, err := game.extractLeagues(testResp)
	if err != nil {
		t.Errorf("extractLeagues failed, expected success")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected leagues extracted: got %v, want %v", got, want)
	}
}
