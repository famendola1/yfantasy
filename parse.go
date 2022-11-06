package yfantasy

import (
	"encoding/xml"
	"strings"

	"github.com/antchfx/xmlquery"
)

func parse(xmlDoc, expr string, out any) error {
	doc, err := xmlquery.Parse(strings.NewReader(xmlDoc))
	if err != nil {
		return err
	}

	node, err := xmlquery.Query(doc, expr)
	if err != nil {
		return err
	}

	return xml.NewDecoder(strings.NewReader(node.OutputXML(true))).Decode(out)
}

func parseAll(xmlDoc, expr string) ([]*xmlquery.Node, error) {
	doc, err := xmlquery.Parse(strings.NewReader(xmlDoc))
	if err != nil {
		return nil, err
	}

	return xmlquery.QueryAll(doc, expr)
}

func parseAllLeagues(xmlDoc string, yf *YFantasy) ([]*League, error) {
	nodes, err := parseAll(xmlDoc, "//league")
	if err != nil {
		return nil, err
	}

	leagues := make([]*League, len(nodes))
	for i, node := range nodes {
		var lg League
		err := xml.NewDecoder(strings.NewReader(node.OutputXML(true))).Decode(&lg)
		if err != nil {
			return nil, err
		}
		lg.yf = yf
		leagues[i] = &lg
	}
	return leagues, nil
}

func parseAllTeams(xmlDoc string, yf *YFantasy) ([]*Team, error) {
	nodes, err := parseAll(xmlDoc, "//team")
	if err != nil {
		return nil, err
	}

	teams := make([]*Team, len(nodes))
	for i, node := range nodes {
		var tm Team
		err := xml.NewDecoder(strings.NewReader(node.OutputXML(true))).Decode(&tm)
		if err != nil {
			return nil, err
		}
		tm.yf = yf
		teams[i] = &tm
	}
	return teams, nil
}

func parseAllPlayers(xmlDoc string) ([]*Player, error) {
	nodes, err := parseAll(xmlDoc, "//player")
	if err != nil {
		return nil, err
	}

	players := make([]*Player, len(nodes))
	for i, node := range nodes {
		var p Player
		err := xml.NewDecoder(strings.NewReader(node.OutputXML(true))).Decode(&p)
		if err != nil {
			return nil, err
		}
		players[i] = &p
	}
	return players, nil
}

func parseAllTransactions(xmlDoc string) ([]*Transaction, error) {
	nodes, err := parseAll(xmlDoc, "//transaction")
	if err != nil {
		return nil, err
	}

	txns := make([]*Transaction, len(nodes))
	for i, node := range nodes {
		var txn Transaction
		err := xml.NewDecoder(strings.NewReader(node.OutputXML(true))).Decode(&txn)
		if err != nil {
			return nil, err
		}
		txns[i] = &txn
	}
	return txns, nil
}

func parseAllString(xmlDoc, expr string) ([]string, error) {
	nodes, err := parseAll(xmlDoc, expr)
	if err != nil {
		return nil, err
	}

	strs := make([]string, len(nodes))
	for i, node := range nodes {
		strs[i] = node.InnerText()
	}
	return strs, nil
}
