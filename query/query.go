package yfantasy

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/famendola1/yfantasy/schema"
)

const (
	endpoint = "https://fantasysports.yahooapis.com/fantasy/v2"
)

type query struct {
	base         string
	resource     string
	isCollection bool
	keys         []string
	outs         []string
	params       []string
}

func (q *query) ToString() string {
	uri := q.base

	if !q.isCollection && len(q.keys) == 0 {
		return ""
	}

	if q.isCollection {
		uri += q.resource + "s"
		if len(q.keys) != 0 {
			uri += fmt.Sprintf(";%s_keys=%s", q.resource, strings.Join(q.keys, ","))
		}
	} else {
		uri += fmt.Sprintf("%s/%s", q.resource, q.keys[0])
	}

	if len(q.outs) == 1 {
		uri += fmt.Sprintf("/%s", q.outs[0])
	}
	if len(q.outs) > 1 {
		uri += fmt.Sprintf(";out=%s", strings.Join(q.outs, ","))
	}

	if len(q.params) != 0 {
		uri += (";" + strings.Join(q.params, ";"))
	}

	return uri
}

func (q *query) Get(client *http.Client) (*schema.FantasyContent, error) {
	var fc schema.FantasyContent
	if err := getAndParse(client, q.ToString(), "//fantasy_content", &fc); err != nil {
		return nil, err
	}
	return &fc, nil
}

func (q *query) Reset() {
	q.keys = []string{}
	q.outs = []string{}
	q.params = []string{}
}

// get sends a GET request to the provided URI and returns the repsone as a
// string.
func get(client *http.Client, uri string) (string, error) {
	resp, err := client.Get(fmt.Sprintf("%s/%s", endpoint, uri))
	if err != nil {
		return "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", handleError(resp)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String(), nil
}

func getAndParse(client *http.Client, uri, expr string, out any) error {
	rawResp, err := get(client, uri)
	if err != nil {
		return err
	}

	return parse(rawResp, expr, out)
}

// post sends a POST request to the provided URI and returns the response as a
// string.
func post(client *http.Client, uri, data string) error {
	resp, err := client.Post(fmt.Sprintf("%s/%s", endpoint, uri),
		"application/xml", strings.NewReader(data))
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return handleError(resp)
	}

	return nil
}

// handleError returns an error containing the error message in the response.
func handleError(resp *http.Response) error {
	doc, err := xmlquery.Parse(resp.Body)
	if err != nil {
		return err
	}

	node, err := xmlquery.Query(doc, "//description")
	if err != nil {
		return err
	}

	return fmt.Errorf("%s: %s", resp.Status, node.InnerText())
}
