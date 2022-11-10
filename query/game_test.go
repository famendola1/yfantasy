package query

import "testing"

func TestGameQueryCollection(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Games(),
				"/games",
			},
			{
				Games().IsAvailable(),
				"/games;is_available=1",
			},
			{
				Games().Types([]string{"full", "pickem-team"}),
				"/games;game_types=full,pickem-team",
			},
			{
				Games().Codes([]string{"nba", "nhl", "nfl"}),
				"/games;game_codes=nba,nhl,nfl",
			},
			{
				Games().Seasons([]int{2022}),
				"/games;seasons=2022",
			},
			{
				Games().Keys([]string{"223", "224"}),
				"/games;game_keys=223,224",
			},
			{
				Games().Key("223"),
				"/games;game_keys=223",
			},
			{
				Games().Key("224"),
				"/games;game_keys=224",
			},
			{
				Games().IsAvailable().Codes([]string{"nba"}).Seasons([]int{2022, 2021}).Types([]string{"full"}).Keys([]string{"223"}),
				"/games;game_keys=223;is_available=1;game_codes=nba;seasons=2022,2021;game_types=full",
			},
		})
}

func TestGameQueryResource(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Game(),
				"/game",
			},
			{
				Game().Key("nba"),
				"/game/nba",
			},
			{
				Game().Key("nba").Leagues(),
				"/game/nba/leagues",
			},
			{
				Game().Key("nba").Teams(),
				"/game/nba/teams",
			},
		})
}
