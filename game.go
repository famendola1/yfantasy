package yfantasy

// Game represents a Yahoo game
type Game struct {
	GameKey            string `xml:"game_key"`
	GameID             int    `xml:"game_id"`
	Name               string `xml:"name"`
	Code               string `xml:"code"`
	Type               string `xml:"type"`
	URL                string `xml:"url"`
	Season             string `xml:"season"`
	IsRegistrationOver bool   `xml:"is_registration_over"`
	IsGameOver         bool   `xml:"is_game_over"`
	IsOffseason        bool   `xml:"is_offseason"`

	yf *YFantasy
}

// Game returns a new Game object.
func (yf *YFantasy) Game(sport string) (*Game, error) {
	rawResp, err := yf.getGameRaw(sport)
	if err != nil {
		return nil, err
	}

	var g Game
	if err := parse(rawResp, "//game", &g); err != nil {
		return nil, err
	}
	g.yf = yf
	return &g, nil
}

// AvailableGames returns a list of all the currently available Yahoo Games.
func (yf *YFantasy) AvailableGames() ([]*Game, error) {
	rawResp, err := yf.getAvailableGamesRaw()
	if err != nil {
		return nil, err
	}

	gms, err := parseAllGames(rawResp, yf)
	if err != nil {
		return nil, err
	}
	return gms, nil
}

// Leagues returns all the active leagues the user is in for the game.
func (g *Game) Leagues() ([]*League, error) {
	return g.yf.UserLeagues(g.GameKey)
}

// League creates a League object from the given league id.
func (g *Game) League(leagueID int) (*League, error) {
	return g.yf.League(g.GameKey, leagueID)
}
