package yfantasy

import (
	"fmt"
	"strings"
)

type TeamQuery struct {
	query
}

type TeamRosterQuery struct {
	query
}

func Team() *TeamQuery {
	return &TeamQuery{query{resource: "team"}}
}

func Teams() *TeamQuery {
	return &TeamQuery{query{resource: "team", isCollection: true}}
}

func (t *TeamQuery) Keys(keys []string) *TeamQuery {
	t.keys = append(t.keys, keys...)
	return t
}

func (t *TeamQuery) Key(key string) *TeamQuery {
	t.keys = []string{key}
	return t
}

func (t *TeamQuery) Roster() *TeamQuery {
	t.outs = append(t.outs, "roster")
	return t
}

func (t *TeamQuery) AllMatchups() *TeamQuery {
	t.outs = append(t.outs, "matchups")
	return t
}

func (t *TeamQuery) Matchups(weeks []int) *TeamQuery {
	t.outs = append(t.outs, "matchups")
	t.params = append(t.params, fmt.Sprintf("weeks=%s", strings.Trim(strings.Replace(fmt.Sprint(weeks), " ", ",", -1), "[]")))
	return t
}

func (t *TeamQuery) Stats() *StatsQuery {
	return &StatsQuery{
		query{
			base:     t.ToString(),
			resource: "stats",
		},
	}
}

func (t *TeamQuery) RosterDay(date string) *TeamRosterQuery {
	return &TeamRosterQuery{
		query{
			base:     t.ToString(),
			resource: "roster",
			params:   []string{fmt.Sprintf("date=%s", date)},
		},
	}
}

func (t *TeamQuery) RosterWeek(week int) *TeamRosterQuery {
	return &TeamRosterQuery{
		query{
			base:     t.ToString(),
			resource: "roster",
			params:   []string{fmt.Sprintf("week=%d", week)},
		},
	}
}
