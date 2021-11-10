package yfantasy

const (
	rosterResp = `
  <?xml version="1.0" encoding="UTF-8"?>
  <fantasy_content xmlns:yahoo="http://www.yahooapis.com/v1/base.rng" xmlns="http://fantasysports.yahooapis.com/fantasy/v2/base.rng" xml:lang="en-US" yahoo:uri="http://fantasysports.yahooapis.com/fantasy/v2/team/253.l.102614.t.10/roster/players" time="110.02206802368ms" copyright="Data provided by Yahoo! and STATS, LLC">
    <team>
      <team_key>253.l.102614.t.10</team_key>
      <team_id>10</team_id>
      <name>Matt Dzaman</name>
      <url>https://baseball.fantasysports.yahoo.com/b1/102614/10</url>
      <team_logos>
        <team_logo>
          <size>medium</size>
          <url>https://l.yimg.com/a/i/us/sp/fn/mlb/gr/icon_12_2.gif</url>
        </team_logo>
      </team_logos>
      <managers>
        <manager>
          <manager_id>10</manager_id>
          <nickname>Sean Montgomery</nickname>
          <guid>VZVEVUCLSJAHSM73FMJ4BYFIKU</guid>
          <is_current_login>1</is_current_login>
        </manager>
      </managers>
      <roster>
        <coverage_type>date</coverage_type>
        <date>2011-07-22</date>
        <players count="3">
          <player>
            <player_key>253.p.7569</player_key>
          </player>
          <player>
            <player_key>253.p.7054</player_key>
          </player>
          <player>
            <player_key>253.p.7382</player_key>
          </player>
        </players>
      </roster>
    </team>
  </fantasy_content>`

  standingsResp = `
  <?xml version="1.0" encoding="UTF-8"?>
  <fantasy_content xmlns:yahoo="http://www.yahooapis.com/v1/base.rng" xmlns="http://fantasysports.yahooapis.com/fantasy/v2/base.rng" xml:lang="en-US" yahoo:uri="http://fantasysports.yahooapis.com/fantasy/v2/league/223.l.431/standings" time="201.46489143372ms" copyright="Data provided by Yahoo! and STATS, LLC">
    <league>
      <league_key>223.l.431</league_key>
      <league_id>431</league_id>
      <name>Y! Friends and Family League</name>
      <url>https://football.fantasysports.yahoo.com/archive/pnfl/2009/431</url>
      <draft_status>postdraft</draft_status>
      <num_teams>14</num_teams>
      <edit_key>17</edit_key>
      <weekly_deadline/>
      <league_update_timestamp>1262595518</league_update_timestamp>
      <scoring_type>head</scoring_type>
      <current_week>16</current_week>
      <start_week>1</start_week>
      <end_week>16</end_week>
      <is_finished>1</is_finished>
      <standings>
        <teams count="4">
          <team>
            <team_key>223.l.431.t.10</team_key>
          </team>
          <team>
            <team_key>223.l.431.t.5</team_key>
          </team>
          <team>
            <team_key>223.l.431.t.8</team_key>
          </team>
          <team>
            <team_key>223.l.431.t.12</team_key>
          </team>
        </teams>
      </standings>
    </league>
  </fantasy_content>`

	searchResp = `
  <?xml version="1.0" encoding="UTF-8"?>
  <fantasy_content xml:lang="en-US" yahoo:uri="http://fantasysports.yahooapis.com/fantasy/v2/league/410.l.16883/players;search=Jalen Green" time="25.722980499268ms" copyright="Data provided by Yahoo! and STATS, LLC" refresh_rate="60" xmlns:yahoo="http://www.yahooapis.com/v1/base.rng" xmlns="http://fantasysports.yahooapis.com/fantasy/v2/base.rng">
   <league>
    <league_key>410.l.16883</league_key>
    <league_id>16883</league_id>
    <name>NBA Fantasy 2K22</name>
    <season>2021</season>
    <players count="1">
     <player>
      <player_key>410.p.6513</player_key>
     </player>
    </players>
   </league>
  </fantasy_content>
  `

  gameTestResp = `<?xml version="1.0" encoding="UTF-8"?>
     <fantasy_content xml:lang="en-US" yahoo:uri="http://fantasysports.yahooapis.com/fantasy/v2/game/nba" xmlns:yahoo="http://www.yahooapis.com/v1/base.rng" time="30.575037002563ms" copyright="Data provided by Yahoo! and STATS, LLC" xmlns="http://fantasysports.yahooapis.com/fantasy/v2/base.rng">
      <game>
        <game_key>410</game_key>
        <game_id>410</game_id>
        <name>Basketball</name>
        <code>nba</code>
        <type>full</type>
        <url>https://football.fantasysports.yahoo.com/f1</url>
        <season>2021</season>
      </game>
    </fantasy_content>`

	leagueTestResp = `<?xml version="1.0" encoding="UTF-8"?>
  <fantasy_content xml:lang="en-US" yahoo:uri="http://fantasysports.yahooapis.com/fantasy/v2/users;use_login=1/games;game_keys=nba/leagues" time="37.668943405151ms" copyright="Data provided by Yahoo! and STATS, LLC" refresh_rate="60" xmlns:yahoo="http://www.yahooapis.com/v1/base.rng" xmlns="http://fantasysports.yahooapis.com/fantasy/v2/base.rng">
   <users count="1">
    <user>
     <guid>EKFDPDVSJIGZD64VAL6WYCIH2I</guid>
     <games count="1">
      <game>
       <game_key>410</game_key>
       <game_id>410</game_id>
       <name>Basketball</name>
       <code>nba</code>
       <type>full</type>
       <url>https://basketball.fantasysports.yahoo.com/nba</url>
       <season>2021</season>
       <is_registration_over>0</is_registration_over>
       <is_game_over>0</is_game_over>
       <is_offseason>0</is_offseason>
       <leagues count="3">
        <league>
         <league_key>410.l.16883</league_key>
        </league>
        <league>
         <league_key>410.l.61777</league_key>
        </league>
        <league>
         <league_key>410.l.159928</league_key>
        </league>
       </leagues>
      </game>
     </games>
    </user>
   </users>
  </fantasy_content>`
)
