package league

const (
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
)
