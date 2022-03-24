package yfantasy

const (
	addDropTransaction = `
  <fantasy_content>
    <transaction>
      <type>add/drop</type>
      <players>
        <player>
          <player_key>%s</player_key>
          <transaction_data>
            <type>add</type>
            <destination_team_key>%s</destination_team_key>
          </transaction_data>
        </player>
        <player>
          <player_key>%s</player_key>
          <transaction_data>
            <type>drop</type>
            <source_team_key>%s</source_team_key>
          </transaction_data>
        </player>
      </players>
    </transaction>
  </fantasy_content>
  `

	addTransaction = `
  <fantasy_content>
    <transaction>
      <type>add</type>
      <player>
        <player_key>%s</player_key>
        <transaction_data>
          <type>add</type>
          <destination_team_key>%s</destination_team_key>
        </transaction_data>
      </player>
    </transaction>
</fantasy_content>`

	dropTransaction = `
  <fantasy_content>
    <transaction>
      <type>drop</type>
      <player>
        <player_key>%s</player_key>
        <transaction_data>
          <type>drop</type>
          <source_team_key>%s</source_team_key>
        </transaction_data>
      </player>
    </transaction>
  </fantasy_content>`
)
