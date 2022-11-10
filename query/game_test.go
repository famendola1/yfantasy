package query

import "testing"

func TestGameQueryCollection(t *testing.T) {
	gamesQuery := Games()

	want := "games"
	got := gamesQuery.ToString()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	want = "games;is_available=1"
	got = gamesQuery.IsAvailable().ToString()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	gamesQuery.Reset()
	want = "games;game_types=full,pickem-team"
	got = gamesQuery.Types([]string{"full", "pickem-team"}).ToString()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	gamesQuery.Reset()
	want = "games;game_codes=nba,nhl,nfl"
	got = gamesQuery.Codes([]string{"nba", "nhl", "nfl"}).ToString()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	gamesQuery.Reset()
	want = "games;seasons=2022"
	got = gamesQuery.Seasons([]int{2022}).ToString()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	gamesQuery.Reset()
	want = "games;game_keys=223"
	got = gamesQuery.Keys([]string{"223"}).ToString()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	gamesQuery.Reset()
	want = "games;game_keys=223;is_available=1;game_codes=nba;seasons=2022,2021;game_types=full"
	got = gamesQuery.IsAvailable().Codes([]string{"nba"}).Seasons([]int{2022, 2021}).Types([]string{"full"}).Keys([]string{"223"}).ToString()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGameQueryResource(t *testing.T) {
	gameQuery := Game()

	want := ""
	got := gameQuery.ToString()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	want = "game/nba"
	got = gameQuery.Keys([]string{"nba"}).ToString()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
