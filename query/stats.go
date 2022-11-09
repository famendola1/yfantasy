package yfantasy

import "fmt"

type StatsQuery struct {
	query
}

func (s *StatsQuery) LastWeek() *StatsQuery {
	s.params = append(s.params, "type=lastweek")
	return s
}

func (s *StatsQuery) LastMonth() *StatsQuery {
	s.params = append(s.params, "type=lastmonth")
	return s
}

func (s *StatsQuery) Today() *StatsQuery {
	s.params = append(s.params, "type=date")
	return s
}

func (s *StatsQuery) Day(date string) *StatsQuery {
	s.params = append(s.params, []string{"type=date", fmt.Sprintf("date=%s", date)}...)
	return s
}

func (s *StatsQuery) CurrentSeason() *StatsQuery {
	s.params = append(s.params, "type=season")
	return s
}

func (s *StatsQuery) Season(season string) *StatsQuery {
	s.params = append(s.params, []string{"type=season", fmt.Sprintf("season=%s", season)}...)
	return s
}

func (s *StatsQuery) CurrentSeasonAverage() *StatsQuery {
	s.params = append(s.params, "type=average_season")
	return s
}

func (s *StatsQuery) SeasonAverage(season string) *StatsQuery {
	s.params = append(s.params, []string{"type=average_season", fmt.Sprintf("season=%s", season)}...)
	return s
}
