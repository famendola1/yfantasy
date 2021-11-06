package yfantasy

const (
	addDropTransaction = `
  <fantasy_content>
    <transaction>
      <type>add/drop</type>
      <players>
        <player>
          <player_key>%v</player_key>
          <transaction_data>
            <type>add</type>
            <destination_team_key>%v</destination_team_key>
          </transaction_data>
        </player>
        <player>
          <player_key>%v</player_key>
          <transaction_data>
            <type>drop</type>
            <source_team_key>%v</source_team_key>
          </transaction_data>
        </player>
      </players>
    </transaction>
  </fantasy_content>
  `
)
