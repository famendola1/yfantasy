package yfantasy

// Stats holds multiple Stat
type Stats struct {
	Stat []Stat `xml:"stat"`
}

// Stat represents a stat category in Yahoo.
type Stat []struct {
	StatID int `xml:"stat_id"`
	Value  int `xml:"value"`
}
