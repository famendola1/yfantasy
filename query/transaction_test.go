package query

import "testing"

func TestTransactionsQuery(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Transactions(),
				"/transactions",
			},
			{
				Transactions().Keys([]string{"223.l.431.tr.26", "257.l.193.w.c.2_6390"}),
				"/transactions;transaction_keys=223.l.431.tr.26,257.l.193.w.c.2_6390",
			},
			{
				Transactions().Key("223.l.431.tr.26"),
				"/transactions;transaction_keys=223.l.431.tr.26",
			},
			{
				Transactions().Key("223.l.431.tr.26").Key("257.l.193.w.c.2_6390"),
				"/transactions;transaction_keys=257.l.193.w.c.2_6390",
			},
			{
				Transactions().Types([]string{"add", "drop"}),
				"/transactions;types=add,drop",
			},
			{
				Transactions().Count(25),
				"/transactions;count=25",
			},
			{
				Transactions().TeamKey("nba.l.12345.t.1"),
				"/transactions;team_key=nba.l.12345.t.1",
			},
		})
}

func TestTransactionQuery(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Transaction(),
				"/transaction",
			},
			{
				Transaction().Key("223.l.431.tr.26"),
				"/transaction/223.l.431.tr.26",
			},
		})
}
