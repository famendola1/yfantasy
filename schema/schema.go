package yfantasy

type FantasyContent struct {
	Lang        string  `xml:"lang,attr"`
	URI         string  `xml:"uri,attr"`
	Time        string  `xml:"time,attr"`
	Copyright   string  `xml:"copyright,attr"`
	RefreshRate string  `xml:"refresh_rate,attr"`
	Yahoo       string  `xml:"yahoo,attr"`
	Xmlns       string  `xml:"xmlns,attr"`
	Games       Games   `xml:"games"`
	Game        Game    `xml:"game"`
	Leagues     Leagues `xml:"leagues"`
	League      League  `xml:"league"`
	Teams       Teams   `xml:"teams"`
	Team        Team    `xml:"team"`
	Users       Users   `xml:"users"`
}

type Users struct {
	Count int  `xml:"count,attr"`
	User  User `xml:"user"`
}

type User struct {
	GUID    string  `xml:"guid"`
	Games   Games   `xml:"games"`
	Leagues Leagues `xml:"leagues"`
	Teams   Teams   `xml:"teams"`
}

// Player represents a Yahoo player.
type Player struct {
	PlayerKey                string            `xml:"player_key"`
	PlayerID                 int               `xml:"player_id"`
	Name                     Name              `xml:"name"`
	EditorialPlayerKey       string            `xml:"editorial_player_key"`
	EditorialTeamKey         string            `xml:"editorial_team_key"`
	EditorialTeamFullName    string            `xml:"editorial_team_full_name"`
	EditorialTeamAbbr        string            `xml:"editorial_team_abbr"`
	UniformNumber            int               `xml:"uniform_number"`
	DisplayPosition          string            `xml:"display_position"`
	Headshot                 Headshot          `xml:"headshot"`
	ImageURL                 string            `xml:"image_url"`
	IsUndroppable            bool              `xml:"is_undroppable"`
	PositionType             string            `xml:"position_type"`
	PrimaryPosition          string            `xml:"primary_position"`
	EligiblePositions        EligiblePositions `xml:"eligible_positions"`
	HasPlayerNotes           bool              `xml:"has_player_notes"`
	PlayerNotesLastTimestamp uint64            `xml:"player_notes_last_timestamp"`
	TransactionData          TransactionData   `xml:"transaction_data"`
	PlayerStats              PlayerStats       `xml:"player_stats"`
	Status                   string            `xml:"status"`
	StatusFull               string            `xml:"status_full"`
	InjuryNote               string            `xml:"injury_note"`
	HasRecentPlayerNotes     string            `xml:"has_recent_player_notes"`
	SelectedPosition         SelectedPosition  `xml:"selected_position"`
	IsKeeper                 IsKeeper          `xml:"is_keeper"`
}

// Name for Player
type Name struct {
	Full       string `xml:"full"`
	First      string `xml:"first"`
	Last       string `xml:"last"`
	ASCIIFirst string `xml:"ascii_first"`
	ASCIILast  string `xml:"ascii_last"`
}

// Headshot for Player
type Headshot struct {
	URL  string `xml:"url"`
	Size string `xml:"size"`
}

// EligiblePositions for Player
type EligiblePositions struct {
	Position []string `xml:"position"`
}

// TransactionData for Player
type TransactionData struct {
	Type                string `xml:"type"`
	SourceType          string `xml:"source_type"`
	DestinationType     string `xml:"destination_type"`
	DestinationTeamKey  string `xml:"destination_team_key"`
	DestinationTeamName string `xml:"destination_team_name"`
	SourceTeamKey       string `xml:"source_team_key"`
	SourceTeamName      string `xml:"source_team_name"`
}

// PlayerStats for a Player.
type PlayerStats struct {
	CoverageType string `xml:"coverage_type"`
	Season       string `xml:"season"`
	Date         string `xml:"date"`
	Stats        Stats  `xml:"stats"`
}

// SelectedPosition contains information on a Player's selected position.
type SelectedPosition struct {
	CoverageType string `xml:"coverage_type"`
	Date         string `xml:"date"`
	Position     string `xml:"position"`
	IsFlex       bool   `xml:"is_flex"`
}

// IsKeeper contains keeper information for a Player.
type IsKeeper struct {
	Status string `xml:"status"`
	Cost   string `xml:"cost"`
	Kept   string `xml:"kept"`
}

// Players is a list of players.
type Players struct {
	Count  int      `xml:"count,attr"`
	Player []Player `xml:"player"`
}

// Transaction represents a Yahoo fantasy transaction.
type Transaction struct {
	TransactionKey string  `xml:"transaction_key"`
	TransactionID  int     `xml:"transaction_id"`
	Type           string  `xml:"type"`
	Status         string  `xml:"status"`
	Timestamp      uint64  `xml:"timestamp"`
	Players        Players `xml:"players"`
}

// Matchups hold multiple Matchup.
type Matchups struct {
	Matchup []Matchup `xml:"matchup"`
}

// Matchup represents a Yahoo matchup.
type Matchup struct {
	Week          int         `xml:"week"`
	WeekStart     string      `xml:"week_start"`
	WeekEnd       string      `xml:"week_end"`
	Status        string      `xml:"status"`
	IsPlayoffs    bool        `xml:"is_playoffs"`
	IsConsolation bool        `xml:"is_consolation"`
	IsTied        bool        `xml:"is_tied"`
	WinnerTeamKey string      `xml:"winner_team_key"`
	StatWinners   StatWinners `xml:"stat_winners"`
	Teams         Teams       `xml:"teams"`
}

// StatWinners for Matchup
type StatWinners struct {
	StatWinner []StatWinner `xml:"stat_winner"`
}

// StatWinner represents the winner of a stat category.
type StatWinner struct {
	StatID        int    `xml:"stat_id"`
	WinnerTeamKey string `xml:"winner_team_key"`
	IsTied        bool   `xml:"is_tied"`
}

// Stats holds multiple Stat
type Stats struct {
	Stat []Stat `xml:"stat"`
}

// Stat represents a stat category in Yahoo.
type Stat struct {
	StatID int    `xml:"stat_id"`
	Value  string `xml:"value"`
}

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
}

// Games holds a collection of Games.
type Games struct {
	Count int    `xml:"count,attr"`
	Game  []Game `xml:"game"`
}

// League represents a Yahoo league.
type League struct {
	LeagueKey             string     `xml:"league_key"`
	LeagueID              int        `xml:"league_id"`
	Name                  string     `xml:"name"`
	URL                   string     `xml:"url"`
	LogoURL               string     `xml:"logo_url"`
	DraftStatus           string     `xml:"draft_status"`
	NumTeams              int        `xml:"num_teams"`
	EditKey               string     `xml:"edit_key"`
	WeeklyDeadline        string     `xml:"weekly_deadline"`
	LeagueUpdateTimestamp string     `xml:"league_update_timestamp"`
	ScoringType           string     `xml:"scoring_type"`
	LeagueType            string     `xml:"league_type"`
	Renew                 string     `xml:"renew"`
	ShortInvitationURL    string     `xml:"short_invitation_url"`
	AllowAddToDlExtraPos  string     `xml:"allow_add_to_dl_extra_pos"`
	IsProLeague           bool       `xml:"is_pro_league"`
	IsCashLeague          bool       `xml:"is_cash_league"`
	CurrentWeek           int        `xml:"current_week"`
	StartWeek             int        `xml:"start_week"`
	StartDate             string     `xml:"start_date"`
	EndWeek               int        `xml:"end_week"`
	EndDate               string     `xml:"end_date"`
	GameCode              string     `xml:"game_code"`
	Season                string     `xml:"season"`
	IsFinished            bool       `xml:"is_finished"`
	Standings             Standings  `xml:"standings"`
	Teams                 Teams      `xml:"teams"`
	Scoreboard            Scoreboard `xml:"scoreboard"`
}

// Leagues contains multiple Leagues.
type Leagues struct {
	Count  int      `xml:"count,attr"`
	League []League `xml:"league"`
}

type Scoreboard struct {
	Count    int      `xml:"count,attr"`
	Matchups Matchups `xml:"matchups"`
}

// Standings contains the standings for a league.
type Standings struct {
	Teams Teams `xml:"teams"`
}

// Teams holds multiple Team.
type Teams struct {
	Count int    `xml:"count,attr"`
	Team  []Team `xml:"team"`
}

// Team represents a Yahoo team.
type Team struct {
	TeamKey               string             `xml:"team_key"`
	TeamID                int                `xml:"team_id"`
	Name                  string             `xml:"name"`
	IsOwnedByCurrentLogin bool               `xml:"is_owned_by_current_login"`
	URL                   string             `xml:"url"`
	TeamLogos             TeamLogos          `xml:"team_logos"`
	WaiverPriority        int                `xml:"waiver_priority"`
	NumberOfMoves         int                `xml:"number_of_moves"`
	NumberOfTrades        int                `xml:"number_of_trades"`
	RosterAdds            RosterAdds         `xml:"roster_adds"`
	LeagueScoringType     string             `xml:"league_scoring_type"`
	DraftPosition         int                `xml:"draft_position"`
	HasDraftGrade         bool               `xml:"has_draft_grade"`
	Managers              Managers           `xml:"managers"`
	TeamStats             TeamStats          `xml:"team_stats"`
	TeamPoints            TeamPoints         `xml:"team_points"`
	TeamRemainingGames    TeamRemainingGames `xml:"team_remaining_games"`
	ClinchedPlayoffs      bool               `xml:"clinched_playoffs"`
	TeamStandings         TeamStandings      `xml:"team_standings"`
	Roster                Roster             `xml:"roster"`
	Matchups              Matchups           `xml:"matchups"`
}

// TeamLogos for Team.
type TeamLogos struct {
	TeamLogo []TeamLogo `xml:"team_logo"`
}

// TeamLogo for Team.
type TeamLogo struct {
	Size string `xml:"size"`
	URL  string `xml:"url"`
}

// RosterAdds for Team.
type RosterAdds struct {
	CoverageType  string `xml:"coverage_type"`
	CoverageValue int    `xml:"coverage_value"`
	Value         int    `xml:"value"`
}

// Managers for Team.
type Managers struct {
	Count   int       `xml:"count,attr"`
	Manager []Manager `xml:"manager"`
}

// Manager for Team.
type Manager struct {
	ManagerID      int    `xml:"manager_id"`
	Nickname       string `xml:"nickname"`
	GUID           string `xml:"guid"`
	IsCommissioner bool   `xml:"is_commissioner"`
	IsCurrentLogin bool   `xml:"is_current_login"`
	Email          string `xml:"email"`
	ImageURL       string `xml:"image_url"`
	FeloScore      int    `xml:"felo_score"`
	FeloTier       string `xml:"felo_tier"`
}

// TeamStats for Team.
type TeamStats struct {
	CoverageType string `xml:"coverage_type"`
	Week         int    `xml:"week"`
	Stats        Stats  `xml:"stats"`
}

// TeamPoints for Team.
type TeamPoints struct {
	CoverageType string `xml:"coverage_type"`
	Week         int    `xml:"week"`
	Total        int    `xml:"total"`
}

// TeamRemainingGames for Team.
type TeamRemainingGames struct {
	CoverageType string `xml:"coverage_type"`
	Week         int    `xml:"week"`
	Total        Total  `xml:"total"`
}

// Total for TeamRemainingGames.
type Total struct {
	RemainingGames int `xml:"remaining_games"`
	LiveGames      int `xml:"live_games"`
	CompletedGames int `xml:"completed_games"`
}

// TeamStandings contains information about a Team's ranking in their league.
type TeamStandings struct {
	Rank                    int                     `xml:"rank"`
	OutcomeTotals           OutcomeTotals           `xml:"outcome_totals"`
	DivisionalOutcomeTotals DivisionalOutcomeTotals `xml:"divisional_outcome_totals"`
}

// OutcomeTotals contains information on the outcomes of a Team's matchups.
type OutcomeTotals struct {
	Wins       int     `xml:"wins"`
	Losses     int     `xml:"losses"`
	Ties       int     `xml:"ties"`
	Percentage float32 `xml:"percentage"`
}

// DivisionalOutcomeTotals contains information on the outcomes of a Team's matchups in their division.
type DivisionalOutcomeTotals struct {
	Wins   int `xml:"wins"`
	Losses int `xml:"losses"`
	Ties   int `xml:"ties"`
}

// Roster contains information on a Team's roster.
type Roster struct {
	CoverageType string  `xml:"coverage_type"`
	Date         string  `xml:"date"`
	IsEditable   bool    `xml:"is_editable"`
	Players      Players `xml:"players"`
}
