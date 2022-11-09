package yfantasy

import (
	"fmt"
)

type LeagueQuery struct {
	query
}

type LeaguePlayerQuery struct {
	query
}

func League() *LeagueQuery {
	return &LeagueQuery{query{resource: "league"}}
}

func Leagues() *LeagueQuery {
	return &LeagueQuery{query{resource: "league", isCollection: true}}
}

func (l *LeagueQuery) Key(key string) *LeagueQuery {
	l.keys = []string{key}
	return l
}

func (l *LeagueQuery) Keys(keys []string) *LeagueQuery {
	l.keys = append(l.keys, keys...)
	return l
}

func (l *LeagueQuery) Settings() *LeagueQuery {
	l.outs = append(l.outs, "settings")
	return l
}

func (l *LeagueQuery) Roster() *LeagueQuery {
	l.outs = append(l.outs, "roster")
	return l
}

func (l *LeagueQuery) CurrentScoreboard() *LeagueQuery {
	l.outs = append(l.outs, "scoreboard")
	return l
}

func (l *LeagueQuery) Scoreboard(week int) *LeagueQuery {
	l.outs = append(l.outs, "scoreboard")
	l.params = append(l.params, fmt.Sprintf("week=%d", week))
	return l
}

func (l *LeagueQuery) Standings() *LeagueQuery {
	l.outs = append(l.outs, "standings")
	return l
}

func (l *LeagueQuery) Teams() *TeamQuery {
	tm := Teams()
	tm.base = l.ToString()
	return tm
}

func (l *LeagueQuery) Player() *LeaguePlayerQuery {
	return &LeaguePlayerQuery{
		query{
			base:     l.ToString(),
			resource: "player",
		},
	}
}

func (l *LeagueQuery) Players() *LeaguePlayerQuery {
	return &LeaguePlayerQuery{
		query{
			base:         l.ToString(),
			resource:     "player",
			isCollection: true,
		},
	}
}

func (l *LeaguePlayerQuery) Keys(keys []string) *LeaguePlayerQuery {
	l.keys = append(l.keys, keys...)
	return l
}

func (l *LeaguePlayerQuery) Key(key string) *LeaguePlayerQuery {
	l.keys = []string{key}
	return l
}

func (l *LeaguePlayerQuery) Position(pos string) *LeaguePlayerQuery {
	l.params = append(l.params, fmt.Sprintf("position=%s", pos))
	return l
}

func (l *LeaguePlayerQuery) Status(status string) *LeaguePlayerQuery {
	l.params = append(l.params, fmt.Sprintf("status=%s", status))
	return l
}

func (l *LeaguePlayerQuery) Search(name string) *LeaguePlayerQuery {
	l.params = append(l.params, fmt.Sprintf("search=%s", name))
	return l
}

func (l *LeaguePlayerQuery) Sort(sort string) *LeaguePlayerQuery {
	l.params = append(l.params, fmt.Sprintf("sort=%s", sort))
	return l
}

func (l *LeaguePlayerQuery) SortType(sortType string) *LeaguePlayerQuery {
	l.params = append(l.params, fmt.Sprintf("sort_type=%s", sortType))
	return l
}

func (l *LeaguePlayerQuery) SortSeason(season string) *LeaguePlayerQuery {
	l.params = append(l.params, fmt.Sprintf("sort_season=%s", season))
	return l
}

func (l *LeaguePlayerQuery) SortDate(date string) *LeaguePlayerQuery {
	l.params = append(l.params, fmt.Sprintf("sort_date=%s", date))
	return l
}

func (l *LeaguePlayerQuery) SortWeek(week int) *LeaguePlayerQuery {
	l.params = append(l.params, fmt.Sprintf("sort_week=%d", week))
	return l
}

func (l *LeaguePlayerQuery) Start(start int) *LeaguePlayerQuery {
	l.params = append(l.params, fmt.Sprintf("start=%d", start))
	return l
}

func (l *LeaguePlayerQuery) Count(count int) *LeaguePlayerQuery {
	l.params = append(l.params, fmt.Sprintf("count=%d", count))
	return l
}

func (l *LeaguePlayerQuery) Stats(count int) *StatsQuery {
	return &StatsQuery{
		query{
			base:     l.ToString(),
			resource: "stats",
		},
	}
}
