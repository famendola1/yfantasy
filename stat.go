package yfantasy

// Stats holds multiple Stat
type Stats struct {
	Stat []Stat `xml:"stat"`
}

// Stat represents a stat category in Yahoo.
type Stat []struct {
	StatID string `xml:"stat_id"`
	Value  string `xml:"value"`
}
