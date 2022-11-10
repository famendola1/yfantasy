package query

import "testing"

func TestTeamsQuery(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Teams(),
				"/teams",
			},
			{
				Teams().Keys([]string{"nba.l.12345.t.1", "nba.l.12345.t.2"}),
				"/teams;team_keys=nba.l.12345.t.1,nba.l.12345.t.2",
			},
			{
				Teams().Key("nba.l.12345.t.1"),
				"/teams;team_keys=nba.l.12345.t.1",
			},
			{
				Teams().Key("nba.l.12345.t.1").Key("nba.l.12345.t.2"),
				"/teams;team_keys=nba.l.12345.t.2",
			},
			{
				Teams().Key("nba.l.12345.t.1").Roster(),
				"/teams;team_keys=nba.l.12345.t.1/roster",
			},
			{
				Teams().Key("nba.l.12345.t.1").AllMatchups(),
				"/teams;team_keys=nba.l.12345.t.1/matchups",
			},
			{
				Teams().Key("nba.l.12345.t.1").Matchups([]int{1, 5}),
				"/teams;team_keys=nba.l.12345.t.1/matchups;weeks=1,5",
			},
			{
				Teams().Key("nba.l.12345.t.1").Roster().AllMatchups(),
				"/teams;team_keys=nba.l.12345.t.1;out=roster,matchups",
			},
			{
				Teams().Key("nba.l.12345.t.1").Stats(),
				"/teams;team_keys=nba.l.12345.t.1/stats",
			},
			{
				Teams().Key("nba.l.12345.t.1").StatsWithDuration().LastWeek(),
				"/teams;team_keys=nba.l.12345.t.1/stats;type=lastweek",
			},
			{
				Teams().Key("nba.l.12345.t.1").RosterWeek(3),
				"/teams;team_keys=nba.l.12345.t.1/roster;week=3",
			},
			{
				Teams().Key("nba.l.12345.t.1").RosterDay("2006-10-06"),
				"/teams;team_keys=nba.l.12345.t.1/roster;date=2006-10-06",
			},
		})
}

func TestTeamQuery(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Team(),
				"/team",
			},
			{
				Team().Key("nba.l.12345.t.1"),
				"/team/nba.l.12345.t.1",
			},
			{
				Team().Key("nba.l.12345.t.1").Key("nba.l.12345.t.2"),
				"/team/nba.l.12345.t.2",
			},
			{
				Team().Key("nba.l.12345.t.1").Roster(),
				"/team/nba.l.12345.t.1/roster",
			},
			{
				Team().Key("nba.l.12345.t.1").AllMatchups(),
				"/team/nba.l.12345.t.1/matchups",
			},
			{
				Team().Key("nba.l.12345.t.1").Matchups([]int{1, 5}),
				"/team/nba.l.12345.t.1/matchups;weeks=1,5",
			},
			{
				Team().Key("nba.l.12345.t.1").Roster().AllMatchups(),
				"/team/nba.l.12345.t.1;out=roster,matchups",
			},
			{
				Team().Key("nba.l.12345.t.1").Stats(),
				"/team/nba.l.12345.t.1/stats",
			},
			{
				Team().Key("nba.l.12345.t.1").StatsWithDuration().LastWeek(),
				"/team/nba.l.12345.t.1/stats;type=lastweek",
			},
			{
				Team().Key("nba.l.12345.t.1").RosterWeek(3),
				"/team/nba.l.12345.t.1/roster;week=3",
			},
			{
				Team().Key("nba.l.12345.t.1").RosterDay("2006-10-06"),
				"/team/nba.l.12345.t.1/roster;date=2006-10-06",
			},
		})
}
