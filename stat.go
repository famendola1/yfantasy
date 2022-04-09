package yfantasy

import (
	"strconv"
	"strings"
)

// Stats holds multiple Stat
type Stats struct {
	Stat []Stat `xml:"stat"`
}

// Stat represents a stat category in Yahoo.
type Stat struct {
	StatID int    `xml:"stat_id"`
	Value  string `xml:"value"`
}

// GetValue parses the stat's value and returns it as a float64.
func (s *Stat) GetValue() float64 {
	if strings.ContainsRune(s.Value, '-') {
		return 0
	}

	if strings.ContainsRune(s.Value, '/') {
		vals := strings.Split(s.Value, "/")
		top, _ := strconv.ParseFloat(vals[0], 64)
		bot, _ := strconv.ParseFloat(vals[0], 64)
		return top / bot
	}

	val, _ := strconv.ParseFloat(s.Value, 64)
	return val
}
