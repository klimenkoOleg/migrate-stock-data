package clickhouse

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var CLKH = sq.StatementBuilder.PlaceholderFormat(sq.Question)

func rebind(query string, args ...interface{}) (string, []interface{}, error) {
	query, args, err := sqlx.In(query, args...)
	query = sqlx.Rebind(sqlx.BindType("clickhouse"), query)
	return query, args, err
}
