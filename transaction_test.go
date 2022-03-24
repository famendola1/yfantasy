package yfantasy

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewTransactionFromXML(t *testing.T) {
	want := &Transaction{
		XMLName:        xml.Name{Local: "transaction"},
		TransactionKey: "410.l.16883.tr.227",
		TransactionID:  227,
		Type:           "add/drop",
		Status:         "successful",
		Timestamp:      "1636674697",
		Players: Players{
			Count: "2",
			Player: []*Player{
				{
					XMLName:           xml.Name{Local: "player"},
					PlayerKey:         "410.p.6450",
					PlayerID:          6450,
					Name:              Name{Full: "Paul Reed", First: "Paul", Last: "Reed", ASCIIFirst: "Paul", ASCIILast: "Reed"},
					EditorialTeamAbbr: "PHI",
					DisplayPosition:   "SF",
					PositionType:      "P",
					TransactionData:   TransactionData{Type: "add", SourceType: "freeagents", DestinationType: "team", DestinationTeamKey: "410.l.16883.t.8", DestinationTeamName: "Anti-Vax and INJ"},
				},
				{
					XMLName:           xml.Name{Local: "player"},
					PlayerKey:         "410.p.4488",
					PlayerID:          4488,
					Name:              Name{Full: "George Hill", First: "George", Last: "Hill", ASCIIFirst: "George", ASCIILast: "Hill"},
					EditorialTeamAbbr: "MIL",
					DisplayPosition:   "PG,SG",
					PositionType:      "P",
					TransactionData:   TransactionData{Type: "drop", SourceType: "team", DestinationType: "waivers", SourceTeamKey: "410.l.16883.t.8", SourceTeamName: "Anti-Vax and INJ"}},
			},
		},
	}
	got, err := NewTransactionFromXML(transactionFullTestResp, nil)
	if err != nil {
		fmt.Println(err)
		t.Errorf("NewTransactionFromXML(%q) failed, want success.", transactionFullTestResp)
		return
	}

	if !cmp.Equal(got, want, cmpopts.IgnoreUnexported(Transaction{}, Player{})) {
		t.Errorf("NewTransactionFromXML(%q) = %+v, want %+v", transactionFullTestResp, *got, *want)
	}
}
