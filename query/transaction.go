package query

import (
	"fmt"
	"strings"
)

// TransactionQuery can be used to query the /transactions or /transaction Yahoo
// Fantasy endpoints.
type TransactionQuery struct {
	query
}

// Transaction returns a TransactionQuery for the /transaction endpoint.
func Transaction() *TransactionQuery {
	return &TransactionQuery{query{resource: "transaction"}}
}

// Transactions returns a TransactionQuery for the /transactions endpoint.
func Transactions() *TransactionQuery {
	return &TransactionQuery{query{resource: "transaction", isCollection: true}}
}

// Keys adds the "transaction_keys" parameter with the given keys to the query.
func (t *TransactionQuery) Keys(keys []string) *TransactionQuery {
	t.keys = append(t.keys, keys...)
	return t
}

// Key sets the "transaction_keys" parameter to the the given key. When querying
// the /transaction endpoint the key will be appended to the query path
// (i.e. /transaction/<key>).
func (t *TransactionQuery) Key(key string) *TransactionQuery {
	t.keys = []string{key}
	return t
}

// Types sets the "transaction_types" parameter with the given types to the
// query. Valid types are add,drop,commish,trade.
func (t *TransactionQuery) Types(types []string) *TransactionQuery {
	t.params = append(t.params, fmt.Sprintf("types=%s", strings.Join(types, ",")))
	return t
}

// Count adds the "count" parameter with the given count to the query. count is
// expected by Yahoo to be a positive integer.
func (t *TransactionQuery) Count(count int) *TransactionQuery {
	t.params = append(t.params, fmt.Sprintf("count=%d", count))
	return t
}

// TeamKey adds the "team_key" parameter with the provided key to the query.
func (t *TransactionQuery) TeamKey(key string) *TransactionQuery {
	t.params = append(t.params, fmt.Sprintf("team_key=%s", key))
	return t
}
