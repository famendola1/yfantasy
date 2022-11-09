package yfantasy

import (
	"fmt"
	"strings"
)

type GameQuery struct {
	query
}

func Game() *GameQuery {
	return &GameQuery{query{resource: "game"}}
}

func Games() *GameQuery {
	return &GameQuery{query{resource: "game", isCollection: true}}
}

func (g *GameQuery) IsAvailable() *GameQuery {
	g.params = append(g.params, "is_available=1")
	return g
}

func (g *GameQuery) Keys(keys []string) *GameQuery {
	g.keys = append(g.keys, keys...)
	return g
}

func (g *GameQuery) Key(key string) *GameQuery {
	g.keys = []string{key}
	return g
}

func (g *GameQuery) Types(types []string) *GameQuery {
	g.params = append(g.params, fmt.Sprintf("game_types=%s", strings.Join(types, ",")))
	return g
}

func (g *GameQuery) Codes(codes []string) *GameQuery {
	g.params = append(g.params, fmt.Sprintf("game_codes=%s", strings.Join(codes, ",")))
	return g
}

func (g *GameQuery) Seasons(seasons []string) *GameQuery {
	g.params = append(g.params, fmt.Sprintf("seasons=%s", strings.Join(seasons, ",")))
	return g
}

func (g *GameQuery) Leagues() *LeagueQuery {
	lg := Leagues()
	lg.base = g.ToString()
	return lg
}

func (g *GameQuery) Teams() *TeamQuery {
	tm := Teams()
	tm.base = g.ToString()
	return tm
}
