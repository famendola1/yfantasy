package query

import (
	"fmt"
)

// LeagueQuery can be used to query the /leagues or /league Yahoo Fantasy endpoints.
type LeagueQuery struct {
	query
}

// League returns a LeagueQuery for the /league endpoint.
func League() *LeagueQuery {
	return &LeagueQuery{query{resource: "league"}}
}

// Leagues returns a LeagueQuery for the /leagues endpoint.
func Leagues() *LeagueQuery {
	return &LeagueQuery{query{resource: "league", isCollection: true}}
}

// Key sets the league_keys parameter to the the given key. When querying the
// /league endpoint the key will be appended to the query path (i.e. /league/<key>).
func (l *LeagueQuery) Key(key string) *LeagueQuery {
	l.keys = []string{key}
	return l
}

// Keys adds the league_keys parameter with the given keys to the query.
func (l *LeagueQuery) Keys(keys []string) *LeagueQuery {
	l.keys = append(l.keys, keys...)
	return l
}

// Settings adds the "settings" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. league/settings).
func (l *LeagueQuery) Settings() *LeagueQuery {
	l.outs = append(l.outs, "settings")
	return l
}

// CurrentScoreboard adds the "scoreboard" subresource to the request to query
// the scoreboard for the current week. If combined with other subresources,
// they are all combined into the "out" parameter, otherwise it is added to the
// request path (i.e. league/scoreboard).
func (l *LeagueQuery) CurrentScoreboard() *LeagueQuery {
	l.outs = append(l.outs, "scoreboard")
	return l
}

// Scoreboard adds the "scoreboard" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. league/scoreboard). week is expected
// by Yahoo to be a positive integer.
func (l *LeagueQuery) Scoreboard(week int) *LeagueQuery {
	l.outs = append(l.outs, "scoreboard")
	l.params = append(l.params, fmt.Sprintf("week=%d", week))
	return l
}

// Standings adds the "standings" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. league/standings).
func (l *LeagueQuery) Standings() *LeagueQuery {
	l.outs = append(l.outs, "standings")
	return l
}

// Teams returns a TeamQuery for the /teams subresource.
func (l *LeagueQuery) Teams() *TeamQuery {
	tm := Teams()
	tm.base = l.ToString()
	return tm
}

// Player returns a PlayerQuery for the /player subresource.
func (l *LeagueQuery) Player() *PlayerQuery {
	return &PlayerQuery{
		query{
			base:     l.ToString(),
			resource: "player",
		},
	}
}

// Players returns a PlayerQuery for the /players subresource.
func (l *LeagueQuery) Players() *PlayerQuery {
	return &PlayerQuery{
		query{
			base:         l.ToString(),
			resource:     "player",
			isCollection: true,
		},
	}
}
