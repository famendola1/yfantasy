package query

import "testing"

func TestPlayersQuery(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Players(),
				"/players",
			},
			{
				Players().Keys([]string{"253.p.7569", "253.p.4365"}),
				"/players;player_keys=253.p.7569,253.p.4365",
			},
			{
				Players().Key("253.p.4365"),
				"/players;player_keys=253.p.4365",
			},
			{
				Players().Key("253.p.4365").Key("253.p.7569"),
				"/players;player_keys=253.p.7569",
			},
			{
				Players().Position("PG"),
				"/players;position=PG",
			},
			{
				Players().Status("A"),
				"/players;status=A",
			},
			{
				Players().Search("Kevin Durant"),
				"/players;search=Kevin+Durant",
			},
			{
				Players().Sort("NAME"),
				"/players;sort=NAME",
			},
			{
				Players().SortType("season"),
				"/players;sort_type=season",
			},
			{
				Players().SortSeason(2022),
				"/players;sort_season=2022",
			},
			{
				Players().SortDate("2006-10-06"),
				"/players;sort_date=2006-10-06",
			},
			{
				Players().SortWeek(3),
				"/players;sort_week=3",
			},
			{
				Players().Start(25),
				"/players;start=25",
			},
			{
				Players().Count(50),
				"/players;count=50",
			},
			{
				Players().Key("253.p.4365").Stats(),
				"/players;player_keys=253.p.4365/stats",
			},
		})
}

func TestPlayerQuery(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Player(),
				"/player",
			},
			{
				Player().Key("253.p.4365"),
				"/player/253.p.4365",
			},
			{
				Player().Key("253.p.4365").Key("253.p.7569"),
				"/player/253.p.7569",
			},
			{
				Player().Key("253.p.4365").Stats(),
				"/player/253.p.4365/stats",
			},
		})
}
