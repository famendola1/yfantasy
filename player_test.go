package yfantasy

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewPlayer(t *testing.T) {
	want := &Player{PlayerKey: "123.p.789"}
	got := NewPlayer("123.p.789", nil)

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(Player{})) {
		t.Errorf("New() = %+v, want %+v", *got, *want)
	}
}

func TestNewPLayerFromXML(t *testing.T) {
	want := &Player{
		XMLName:                  xml.Name{Local: "player"},
		PlayerKey:                "410.p.6065",
		PlayerID:                 "6065",
		Name:                     Name{Full: "Shake Milton", First: "Shake", Last: "Milton", ASCIIFirst: "Shake", ASCIILast: "Milton"},
		EditorialPlayerKey:       "nba.p.6065",
		EditorialTeamKey:         "nba.t.20",
		EditorialTeamFullName:    "Philadelphia 76ers",
		EditorialTeamAbbr:        "PHI",
		UniformNumber:            "18",
		DisplayPosition:          "PG,SG",
		Headshot:                 Headshot{URL: "https://s.yimg.com/iu/api/res/1.2/PTF3UNtaGJzwH3Ah22R0Ow--~C/YXBwaWQ9eXNwb3J0cztjaD0yMzM2O2NyPTE7Y3c9MTc5MDtkeD04NTc7ZHk9MDtmaT11bGNyb3A7aD02MDtxPTEwMDt3PTQ2/https://s.yimg.com/xe/i/us/sp/v/nba_cutout/players_l/10142021/6065.png", Size: "small"},
		ImageURL:                 "https://s.yimg.com/iu/api/res/1.2/PTF3UNtaGJzwH3Ah22R0Ow--~C/YXBwaWQ9eXNwb3J0cztjaD0yMzM2O2NyPTE7Y3c9MTc5MDtkeD04NTc7ZHk9MDtmaT11bGNyb3A7aD02MDtxPTEwMDt3PTQ2/https://s.yimg.com/xe/i/us/sp/v/nba_cutout/players_l/10142021/6065.png",
		IsUndroppable:            "0",
		PositionType:             "P",
		PrimaryPosition:          "PG",
		EligiblePositions:        EligiblePositions{Position: []string{"PG", "SG", "G", "Util"}},
		HasPlayerNotes:           "1",
		PlayerNotesLastTimestamp: "1636515545",
	}
	got, err := NewPlayerFromXML(playerFullTestResp, nil)
	if err != nil {
		t.Errorf("NewPlayerFromXML(%q) failed, want success.", playerFullTestResp)
		return
	}

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(Player{})) {
		t.Errorf("NewPlayerFromXML(%q) = %+v, want %+v", playerFullTestResp, *got, *want)
	}
}
