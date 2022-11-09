package yfantasy

import (
	"fmt"
	"strings"
)

type TransactionQuery struct {
	query
}

func Transaction() *TransactionQuery {
	return &TransactionQuery{query{resource: "transaction"}}
}

func Transactions() *TransactionQuery {
	return &TransactionQuery{query{resource: "transaction", isCollection: true}}
}

func (t *TransactionQuery) Keys(keys []string) *TransactionQuery {
	t.keys = append(t.keys, keys...)
	return t
}

func (t *TransactionQuery) Types(types []string) *TransactionQuery {
	t.params = append(t.params, fmt.Sprintf("types=%s", strings.Join(types, ",")))
	return t
}

func (t *TransactionQuery) Count(count int) *TransactionQuery {
	t.params = append(t.params, fmt.Sprintf("count=%d", count))
	return t
}

func (t *TransactionQuery) TeamKey(key string) *TransactionQuery {
	t.params = append(t.params, fmt.Sprintf("team_key=%s", key))
	return t
}
