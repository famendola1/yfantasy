package query

import (
	"fmt"
	"strings"
)

// GameQuery can be used to query the /games or /game Yahoo Fantasy endpoints.
type GameQuery struct {
	query
}

// Game returns a GameQuery for the /game endpoint.
func Game() *GameQuery {
	return &GameQuery{query{resource: "game"}}
}

// Games returns a GameQuery for the /games endpoint.
func Games() *GameQuery {
	return &GameQuery{query{resource: "game", isCollection: true}}
}

// IsAvailable adds the "is_available" parameter to the query.
func (g *GameQuery) IsAvailable() *GameQuery {
	g.params = append(g.params, "is_available=1")
	return g
}

// Keys adds the "game_keys" parameter with the given keys to the query.
func (g *GameQuery) Keys(keys []string) *GameQuery {
	g.keys = append(g.keys, keys...)
	return g
}

// Key sets the "game_keys" parameter to the the given key. When querying the
// /game endpoint the key will be appended to the query path (i.e. /game/<key>).
func (g *GameQuery) Key(key string) *GameQuery {
	g.keys = []string{key}
	return g
}

// Types adds the game_types parameter with the given types to the query. Valid
// game types are full|pickem-team|pickem-group|pickem-team-list.
func (g *GameQuery) Types(types []string) *GameQuery {
	g.params = append(g.params, fmt.Sprintf("game_types=%s", strings.Join(types, ",")))
	return g
}

// Codes adds the "game_codes" parameter with the given codes to the query. Any
// valid game code can be provided (incl. nba, nhl, nfl, mlb).
func (g *GameQuery) Codes(codes []string) *GameQuery {
	g.params = append(g.params, fmt.Sprintf("game_codes=%s", strings.Join(codes, ",")))
	return g
}

// Seasons adds the "seasons" parameter with the given seasons to the query.
func (g *GameQuery) Seasons(seasons []int) *GameQuery {
	g.params = append(g.params, fmt.Sprintf("seasons=%s", strings.Trim(strings.Replace(fmt.Sprint(seasons), " ", ",", -1), "[]")))
	return g
}

// Leagues returns a LeagueQuery for the /leagues subresource.
func (g *GameQuery) Leagues() *LeagueQuery {
	lg := Leagues()
	lg.base = g.ToString()
	return lg
}

// Teams returns a TeamQuery for the /teams subresource.
func (g *GameQuery) Teams() *TeamQuery {
	tm := Teams()
	tm.base = g.ToString()
	return tm
}
