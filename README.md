# (Deprecated) Yahoo Fantasy API Client for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/famendola1/yfantasy.svg)](https://pkg.go.dev/github.com/famendola1/yfantasy)
![License](https://img.shields.io/badge/License-Apache-green)
[![Go Report Card](https://goreportcard.com/badge/github.com/famendola1/yfantasy)](https://goreportcard.com/report/github.com/famendola1/yfantasy)

## Notice
This package has been replaced by the [yfquery](https://pkg.go.dev/github.com/famendola1/yfquery) and [yflib](https://pkg.go.dev/github.com/famendola1/yflib) packages.

## Installation
~~~bash
go get github.com/famendola1/yfantasy
~~~

## Yahoo Endpoints
The Yahoo offical documentation for their [Fantasy Sports API](https://developer.yahoo.com/fantasysports/guide) is not comprehensive and incomplete, despite being the offical. For a more complete overview of the supported endpoints, see the [README](https://github.com/edwarddistel/yahoo-fantasy-baseball-reader#yahoo-fantasy-api-docs).

## Usage
There are two ways to interact with the Yahoo Fantasy API through this Go client:
* Using high-level functions in [api.go](api.go)
* Using the query builders in [query/](query/)

### Before You Start
Both of these usages require the use of a `*http.Client` that is configured for the Yahoo Fantasy API endpoint. You can use the [github.com/famendola1/yauth](https://pkg.go.dev/github.com/famendola1/yauth) package to configure a `*http.Client` to use.

### High Level Functions
The high level functions in [api.go](api.go) provide additional functionality on top of the query builders. These functions can be accessed through the [`YFantasy`](https://pkg.go.dev/github.com/famendola1/yfantasy#YFantasy) struct. A valid `YFantasy` struct must be initialized with [`yfantasy.New`](https://pkg.go.dev/github.com/famendola1/yfantasy#New) prior to using these functions.

Contributions to set the high level functions provided through this client are welcome 🙂.

### Query Builders
The query buiilders were designed to be able to easily generate queries for all the Yahoo Fantasy API endpoints. The builders expose functions that add pieces and parameters to the query. They also expose the following functions:

* `ToString`: Builds the string for the query that the builder represents.
* `Get`: Sends a GET request to the Yahoo Fantasy API for the endpoint that the query represents. A successful query will return a [`FantasyContent`](https://pkg.go.dev/github.com/famendola1/yfantasy/schema#FantasyContent) struct.

WARNING: The query builders do not validate that the queries they build are valid Yahoo endpoints.

#### Examples
```go
// Get all the teams in a league.
Leagues().Key("nba.l.12345").Teams().Get(client)

// Search for a player and return their stats for the past week.
Leagues().Key("nba.l.12345").Players().Search("Donovan Mitchell").Stats().LastWeek().Get(client)

// Get all the leagues a user is in.
Users().Leagues().Get(client)

// Get the rosters for all teams in a league.
League().Key("nba.l.12345").Teams().Roster().Get(client)

// Get all the add and drop transactions in a league
League().Key("nba.l.12345").Transactions().Types([]string{"add", "drop"}).Get(client)
```

## Projects using this client
A fantasy sports Discord bot is being developed with this client: https://github.com/famendola1/fantasy-discord-bot.
