package query

import "fmt"

// PlayerQuery can be used to query the /players or /player Yahoo Fantasy endpoints.
type PlayerQuery struct {
	query
}

// Keys adds the "player_keys" parameter with the given keys to the query.
func (p *PlayerQuery) Keys(keys []string) *PlayerQuery {
	p.keys = append(p.keys, keys...)
	return p
}

// Key sets the "player_keys" parameter to the the given key. When querying the
// /game endpoint the key will be appended to the query path (i.e. /player/<key>).
func (p *PlayerQuery) Key(key string) *PlayerQuery {
	p.keys = []string{key}
	return p
}

// Position adds the "position" parameter with the provided position to the query.
// Valid player positions can be provided as input (e.x. "QB", "PG")
func (p *PlayerQuery) Position(pos string) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("position=%s", pos))
	return p
}

// Status adds the "status" parameter with the provided status to the query.
// Valid statuses are:
// A (all available players)
// FA (free agents only)
// W (waivers only)
// T (all taken players)
// K (keepers only)
func (p *PlayerQuery) Status(status string) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("status=%s", status))
	return p
}

// Search adds the "search" parameter with the provided name to the query.
func (p *PlayerQuery) Search(name string) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("search=%s", name))
	return p
}

// Sort adds the "sort" parameter with the provided sort criteria to the query.
// Valid inputs are:
// {stat_id}
// NAME (last, first)
// OR (overall rank)
// AR (actual rank)
// PTS (fantasy points)
func (p *PlayerQuery) Sort(sort string) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("sort=%s", sort))
	return p
}

// SortType adds the "sort_type" parameter with the provided type to the query.
// Valid types are:
// season
// date (baseball, basketball, and hockey only)
// week (football only)
// lastweek (baseball, basketball, and hockey only)
// lastmonth
func (p *PlayerQuery) SortType(sortType string) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("sort_type=%s", sortType))
	return p
}

// SortSeason adds the "sort_season" parameter with the provided season to the
// query.
func (p *PlayerQuery) SortSeason(season int) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("sort_season=%d", season))
	return p
}

// SortDate adds the "sort_date" parameter with the provided date to the query.
// date should be formatted as YYYY-MM-DD.
func (p *PlayerQuery) SortDate(date string) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("sort_date=%s", date))
	return p
}

// SortWeek adds the "sort_week" parameter with the provided week to the query.
// Yahoo only supports this parameter for football. week is expected by Yahoo
// to be a positive integer.
func (p *PlayerQuery) SortWeek(week int) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("sort_week=%d", week))
	return p
}

// Start adds the "start" parameter with the provided start to the query. start
// is expected by Yahoo to be a positive integer.
func (p *PlayerQuery) Start(start int) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("start=%d", start))
	return p
}

// Count adds the "count" parameter with the provided count to the query. count
// is expected by Yahoo to be a positive integer.
func (p *PlayerQuery) Count(count int) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("count=%d", count))
	return p
}

// Stats returns a StatsQuery for the /stats subresource.
func (p *PlayerQuery) Stats(count int) *StatsQuery {
	return &StatsQuery{
		query{
			base:     p.ToString(),
			resource: "stats",
		},
	}
}
