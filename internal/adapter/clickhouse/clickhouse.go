package clickhouse

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/dto"
	"go.uber.org/zap"
	"os"
	"time"
)

type Repo struct {
	db   *sqlx.DB
	log  *zap.Logger
	rows *sql.Rows
}

/*
	func (s *Repo) ScanRow(rows *sql.Rows) (interface{}, error) {
		tick := Tick{}
		err := rows.Scan(&tick.InstrumentId, &tick.StockName, &tick.Timestamp, &tick.Value)
		if err != nil {
			return nil, err
		}

		return tick, nil
	}
*/
func NewRepo(
	log *zap.Logger,
	db *sqlx.DB,
) *Repo {
	return &Repo{db: db, log: log.Named("repo_clickhouse")}
}

func MustGetClickhouseConnection(ctx context.Context, clickHouseDSN string, log *zap.Logger) *sqlx.DB {

	db, err := sqlx.Open("clickhouse", clickHouseDSN)
	if err != nil {
		log.Error("failed to open db", zap.Error(err))
		os.Exit(1)
	}

	if err = db.PingContext(ctx); err != nil {
		log.Error("failed to ping db", zap.Error(err))
		os.Exit(1)
	}
	//rtr := adapter.RetryWithError(db.PingContext, "clickhouse_ping", 5, 5*time.Second, c.log)

	//if err = rtr(ctx); err != nil {
	//	log.Error("failed to ping db", zap.Error(err))
	//	os.Exit(1)
	//}

	//c.closer.AddCloser(db)
	log.Info("connected to clickhouse")

	return db
}

func (r *Repo) Prepare(ctx context.Context, startDate time.Time, endDate time.Time) error {
	startDateStr := fmt.Sprintf(startDate.Format(layout))
	endDateStr := fmt.Sprintf(endDate.Format(layout))

	stmt := CLKH.
		Select("case when stock_name=='NSE' THEN 1 ELSE 2 end", "instrument_id", "start_at", "maxMerge(max_value)", "minMerge(min_value)", "avgMerge(avg_value)", "quantileMerge(0.5)(med_value)").
		From("tickers.tickers_aggregates_1day_mv").
		Where("start_at >= toDateTime('"+startDateStr+" 00:00:00')").
		Where("start_at < toDateTime('"+endDateStr+" 00:00:00')").
		//Where(squirrel.GtOrEq{"start_at": startDate}).
		//Where(squirrel.LtOrEq{"start_at": startDate}).
		//Where(squirrel.Lt{"start_at": endDate}).
		//Where("t.timestamp < toInt64(toDateTime64('" + endDate + "', 0, 'UTC')) * 1000").
		Where(squirrel.NotEq{"instrument_id": "0"}).
		GroupBy("stock_name", "instrument_id", "start_at").
		OrderBy("stock_name", "instrument_id", "start_at")
	query, args, err := stmt.ToSql()
	if err != nil {
		return err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	r.rows = rows

	return nil
}

func (r *Repo) Scan() (any, error) {
	tick := dto.Tick1Day{}
	date := time.Time{}
	err := r.rows.Scan(&tick.StockName, &tick.InstrumentId, &date, &tick.Max, &tick.Min, &tick.Avg, &tick.Median)
	tick.Date = date.UnixMilli()
	if err != nil {
		return nil, err
	}
	return tick, nil
}

func (r *Repo) Next() bool {
	return r.rows.Next()
}
