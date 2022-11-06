package yfantasy

// Transaction represents a Yahoo fantasy transaction.
type Transaction struct {
	TransactionKey string  `xml:"transaction_key"`
	TransactionID  int     `xml:"transaction_id"`
	Type           string  `xml:"type"`
	Status         string  `xml:"status"`
	Timestamp      string  `xml:"timestamp"`
	Players        Players `xml:"players"`

	yf *YFantasy
}
