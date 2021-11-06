package game

const (
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
