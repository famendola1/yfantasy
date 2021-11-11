package yfantasy

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewTeamFromXML(t *testing.T) {
	want := &Team{
		XMLName:               xml.Name{Local: "team"},
		TeamKey:               "410.l.16883.t.1",
		TeamID:                "1",
		Name:                  "Bring Me A Shot",
		IsOwnedByCurrentLogin: "1",
		URL:                   "https://basketball.fantasysports.yahoo.com/nba/16883/1",
		TeamLogos: TeamLogos{
			TeamLogo: []TeamLogo{
				{
					Size: "large",
					URL:  "https://yahoofantasysports-res.cloudinary.com/image/upload/t_s192sq/fantasy-logos/6d1c4b1ab7f7d94e33ea9f4a3306381954a2880b244fe54ca8dfea504c7be242.jpg",
				},
			},
		},
		WaiverPriority:    "9",
		NumberOfMoves:     "23",
		NumberOfTrades:    "0",
		RosterAdds:        RosterAdds{CoverageType: "week", CoverageValue: "4", Value: "3"},
		LeagueScoringType: "headone",
		DraftPosition:     "4",
		HasDraftGrade:     "0",
		Managers: Managers{
			Manager: []Manager{
				{
					ManagerID:      "1",
					Nickname:       "Fabio",
					IsCommissioner: "1",
					IsCurrentLogin: "1",
					Email:          "example@gmail.com",
					ImageURL:       "https://s.yimg.com/ag/images/default_user_profile_pic_64sq.jpg",
					FeloScore:      "890",
					FeloTier:       "platinum",
				},
			},
		},
	}
	got, err := NewTeamFromXML(teamFullTestResp, nil)
	if err != nil {
		t.Errorf("NewTeamFromXML(%q) failed, want success.", leagueFullTestResp)
		return
	}

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(Team{})) {
		t.Errorf("NewTeamFromXML(%q) = %+v, want %+v", leagueFullTestResp, *got, *want)
	}
}

func TestNewTeam(t *testing.T) {
	want := &Team{XMLName: xml.Name{Local: "team"}, TeamKey: "123.l.456.t.789"}
	got := NewTeam("123.l.456.t.789", nil)

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(Team{})) {
		t.Errorf("New() %+v, want %+v", *got, *want)
	}
}

func TestLeagueKey(t *testing.T) {
	want := "123.l.456"
	team := NewTeam("123.l.456.t.789", nil)
	got := team.LeagueKey()

	if got != want {
		t.Errorf("LeagueKey() = %q, want %q", got, want)
	}
}

func TestExtractPlayersForTeam(t *testing.T) {
	team := NewTeam("123.1.456.t.789", nil)
	want := []*Player{
		NewPlayer(nil, "253.p.7569"),
		NewPlayer(nil, "253.p.7054"),
		NewPlayer(nil, "253.p.7382"),
	}
	got, err := team.extractPlayersFromRoster(rosterResp)
	if err != nil {
		t.Errorf("extractPlayersFromRoster(%q) failed, want success", rosterResp)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractPlayersFromRoster(%q) = %+v, want %+v", rosterResp, got, want)
	}
}
