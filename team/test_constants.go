package team

const (
	rosterResp string = `
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
)
