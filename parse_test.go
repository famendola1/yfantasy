package yfantasy

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestParseLeague(t *testing.T) {
	want := &League{LeagueKey: "410.l.16883"}
	got := &League{}

	if err := parse(leagueTestResp, "//league", got); err != nil {
		t.Errorf("parse(%s, \"//league\", %+v) failed, want success", leagueTestResp, got)
		return
	}

	if diff := cmp.Diff(got, want, cmpopts.IgnoreFields(League{}, "yf")); diff != "" {
		t.Errorf("diff (-got, +want):\n %+s", diff)
	}
}

func TestParseAllLeagues(t *testing.T) {
	yf := &YFantasy{}
	want := []*League{
		&League{LeagueKey: "410.l.16883"},
		&League{LeagueKey: "410.l.61777"},
		&League{LeagueKey: "410.l.159928"},
	}
	got, err := parseAllLeagues(leagueTestResp, yf)
	if err != nil {
		t.Errorf("parseAllLeagues(%s, %+v) failed, want success", leagueTestResp, yf)
		return
	}

	if diff := cmp.Diff(got, want, cmpopts.IgnoreFields(League{}, "yf")); diff != "" {
		t.Errorf("diff (-got, +want):\n %s", diff)
	}
}

func TestParseTeam(t *testing.T) {
	want := &Team{
		TeamKey:               "410.l.16883.t.1",
		TeamID:                1,
		Name:                  "Bring Me A Shot",
		IsOwnedByCurrentLogin: true,
		URL:                   "https://basketball.fantasysports.yahoo.com/nba/16883/1",
		TeamLogos: TeamLogos{
			TeamLogo: []TeamLogo{
				{Size: "large",
					URL: "https://yahoofantasysports-res.cloudinary.com/image/upload/t_s192sq/fantasy-logos/6d1c4b1ab7f7d94e33ea9f4a3306381954a2880b244fe54ca8dfea504c7be242.jpg"}}},
		WaiverPriority:    9,
		NumberOfMoves:     23,
		NumberOfTrades:    0,
		RosterAdds:        RosterAdds{CoverageType: "week", CoverageValue: 4, Value: 3},
		LeagueScoringType: "headone",
		DraftPosition:     4,
		HasDraftGrade:     false,
		Managers: Managers{
			Manager: []Manager{
				{ManagerID: 1,
					Nickname:       "Fabio",
					IsCommissioner: true,
					IsCurrentLogin: true,
					Email:          "example@gmail.com",
					ImageURL:       "https://s.yimg.com/ag/images/default_user_profile_pic_64sq.jpg",
					FeloScore:      890,
					FeloTier:       "platinum"}}}}
	got := &Team{}

	if err := parse(teamFullTestResp, "//team", got); err != nil {
		t.Errorf("parse(%s, \"//team\", %+v) failed, want success", teamFullTestResp, got)
		return
	}

	if diff := cmp.Diff(got, want, cmpopts.IgnoreFields(Team{}, "yf")); diff != "" {
		t.Errorf("diff (-got, +want):\n %s", diff)
	}
}

func TestParseAllTeams(t *testing.T) {
	yf := &YFantasy{}
	want := []*Team{
		&Team{TeamKey: "223.l.431.t.10"},
		&Team{TeamKey: "223.l.431.t.5"},
		&Team{TeamKey: "223.l.431.t.8"},
		&Team{TeamKey: "223.l.431.t.12"},
	}
	got, err := parseAllTeams(standingsTestResp, yf)
	if err != nil {
		t.Errorf("parseAllTeams(%s, %+v) failed, want success", standingsTestResp, yf)
		return
	}

	if diff := cmp.Diff(got, want, cmpopts.IgnoreFields(Team{}, "yf")); diff != "" {
		t.Errorf("diff (-got, +want):\n %s", diff)
	}
}

func TestParsePlayer(t *testing.T) {
	want := &Player{
		PlayerKey:                "410.p.6065",
		PlayerID:                 6065,
		Name:                     Name{Full: "Shake Milton", First: "Shake", Last: "Milton", ASCIIFirst: "Shake", ASCIILast: "Milton"},
		EditorialPlayerKey:       "nba.p.6065",
		EditorialTeamKey:         "nba.t.20",
		EditorialTeamFullName:    "Philadelphia 76ers",
		EditorialTeamAbbr:        "PHI",
		UniformNumber:            18,
		DisplayPosition:          "PG,SG",
		Headshot:                 Headshot{URL: "https://s.yimg.com/iu/api/res/1.2/PTF3UNtaGJzwH3Ah22R0Ow--~C/YXBwaWQ9eXNwb3J0cztjaD0yMzM2O2NyPTE7Y3c9MTc5MDtkeD04NTc7ZHk9MDtmaT11bGNyb3A7aD02MDtxPTEwMDt3PTQ2/https://s.yimg.com/xe/i/us/sp/v/nba_cutout/players_l/10142021/6065.png", Size: "small"},
		ImageURL:                 "https://s.yimg.com/iu/api/res/1.2/PTF3UNtaGJzwH3Ah22R0Ow--~C/YXBwaWQ9eXNwb3J0cztjaD0yMzM2O2NyPTE7Y3c9MTc5MDtkeD04NTc7ZHk9MDtmaT11bGNyb3A7aD02MDtxPTEwMDt3PTQ2/https://s.yimg.com/xe/i/us/sp/v/nba_cutout/players_l/10142021/6065.png",
		IsUndroppable:            false,
		PositionType:             "P",
		PrimaryPosition:          "PG",
		EligiblePositions:        EligiblePositions{Position: []string{"PG", "SG", "G", "Util"}},
		HasPlayerNotes:           true,
		PlayerNotesLastTimestamp: 1636515545}
	got := &Player{}

	if err := parse(playerFullTestResp, "//player", got); err != nil {
		t.Errorf("parse(%s, \"//player\", %+v) failed, want success", playerFullTestResp, got)
		return
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("diff (-got, +want):\n %+s", diff)
	}
}

func TestParseAllPlayers(t *testing.T) {
	want := []*Player{
		&Player{PlayerKey: "253.p.7569"},
		&Player{PlayerKey: "253.p.7054"},
		&Player{PlayerKey: "253.p.7382"},
	}
	got, err := parseAllPlayers(rosterTestResp)
	if err != nil {
		t.Errorf("parseAllPlayers(%s) failed, want success", rosterTestResp)
		return
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("diff (-got, +want):\n %+s", diff)
	}
}

func TestParseTransaction(t *testing.T) {
	want := &Transaction{
		TransactionKey: "410.l.16883.tr.227",
		TransactionID:  227,
		Type:           "add/drop",
		Status:         "successful",
		Timestamp:      1636674697,
		Players: Players{
			Count: 2,
			Player: []Player{
				{PlayerKey: "410.p.6450",
					PlayerID: 6450,
					Name:     Name{Full: "Paul Reed"},
					TransactionData: TransactionData{
						Type:                "add",
						SourceType:          "freeagents",
						DestinationType:     "team",
						DestinationTeamKey:  "410.l.16883.t.8",
						DestinationTeamName: "Anti-Vax and INJ"}},
				{PlayerKey: "410.p.4488",
					PlayerID: 4488,
					Name:     Name{Full: "George Hill"},
					TransactionData: TransactionData{
						Type:            "drop",
						SourceType:      "team",
						DestinationType: "waivers",
						SourceTeamKey:   "410.l.16883.t.8",
						SourceTeamName:  "Anti-Vax and INJ"},
				}}}}
	got := &Transaction{}

	if err := parse(transactionFullTestResp, "//transaction", got); err != nil {
		t.Errorf("parse(%s, \"//transaction\", %+v) failed, want success", transactionFullTestResp, got)
		return
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("diff (-got, +want):\n %+s", diff)
	}
}

func TestParseAllTransactions(t *testing.T) {
	want := []*Transaction{
		&Transaction{TransactionKey: "257.l.193.pt.1"},
		&Transaction{TransactionKey: "257.l.193.pt.2"},
		&Transaction{TransactionKey: "257.l.193.pt.3"},
	}
	got, err := parseAllTransactions(transactionsTestResp)
	if err != nil {
		t.Errorf("parseAllTransactions(%s) failed, want success", transactionsTestResp)
		return
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("diff (-got, +want):\n %+s", diff)
	}
}

func TestParseGame(t *testing.T) {
	want := &Game{
		GameKey: "410",
		GameID:  410,
		Name:    "Basketball",
		Code:    "nba",
		Type:    "full",
		URL:     "https://football.fantasysports.yahoo.com/f1",
		Season:  "2021",
	}
	got := &Game{}

	if err := parse(gameTestResp, "//game", got); err != nil {
		t.Errorf("parse(%s, \"//game\", %+v) failed, want success", gameTestResp, got)
		return
	}

	if diff := cmp.Diff(got, want, cmpopts.IgnoreFields(Game{}, "yf")); diff != "" {
		t.Errorf("diff (-got, +want):\n %+s", diff)
	}
}

func TestParseAllGames(t *testing.T) {
	yf := &YFantasy{}
	want := []*Game{
		&Game{Code: "nba"},
		&Game{Code: "nfl"},
		&Game{Code: "nhl"},
	}
	got, err := parseAllGames(gamesTestResp, yf)
	if err != nil {
		t.Errorf("parseAllGames(%s, %+v) failed, want success", gamesTestResp, yf)
		return
	}

	if diff := cmp.Diff(got, want, cmpopts.IgnoreFields(Game{}, "yf")); diff != "" {
		t.Errorf("diff (-got, +want):\n %+s", diff)
	}
}

func ParseAllString(t *testing.T) {
	want := []string{"410.l.16883", "410.l.61777", "410.l.159928"}
	got, err := parseAllString(leagueTestResp, "//league_key")
	if err != nil {
		t.Errorf("parseAllString(%s, \"//league_key\") failed, want success", leagueTestResp)
		return
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("diff (-got, +want):\n %s", diff)
	}
}
