package yfantasy

import (
	"encoding/xml"
	"strings"
)

// Transaction represents a Yahoo fantasy transaction.
type Transaction struct {
	XMLName        xml.Name `xml:"transaction"`
	TransactionKey string   `xml:"transaction_key"`
	TransactionID  int      `xml:"transaction_id"`
	Type           string   `xml:"type"`
	Status         string   `xml:"status"`
	Timestamp      string   `xml:"timestamp"`
	Players        Players  `xml:"players"`

	yf *YFantasy
}

// newTransactionFromXML returns a new Transaction object parsed from an
// XML string.
func (yf *YFantasy) newTransactionFromXML(rawXML string) (*Transaction, error) {
	var tr Transaction
	err := xml.NewDecoder(strings.NewReader(rawXML)).Decode(&tr)
	if err != nil {
		return nil, err
	}
	tr.yf = yf
	return &tr, nil
}
