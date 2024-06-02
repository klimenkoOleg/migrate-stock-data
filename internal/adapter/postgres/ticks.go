package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/dto"
	"go.uber.org/zap"
	"log"
	"time"
)

type DBDX interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

type Repo struct {
	db   DBDX
	rows *pgx.Rows
	log  *zap.Logger
}

func NewRepo(db DBDX, log *zap.Logger) *Repo {
	return &Repo{db: db, log: log}
}

//const addTick = `-- name: CreatePayment :exec
//INSERT INTO tick1d
//(stock_name, instrument_id, tick_date, max, min, avg, median)
//VALUES
//    ($1, $2, $3, $4, $5, $6, $7);
//`

func (r *Repo) WriteTick(_ context.Context, ticks []dto.Tick1Day) error {
	columns := []string{"stock_name", "instrument_id", "tick_date", "max", "min", "avg", "median"}

	sq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Insert("tick1d").Columns(columns...)

	for _, tick := range ticks {

		stockName := "BSE"
		if tick.StockName == 1 {
			stockName = "NSE"
		}

		rowVals := []interface{}{
			stockName,
			tick.InstrumentId,
			tick.Date,
			tick.Max,
			tick.Min,
			tick.Avg,
			tick.Median,
		}
		sq = sq.Values(rowVals...)
	}

	sql, args, err := sq.ToSql()
	if err != nil {
		return fmt.Errorf("error building SQL for inserting tick, sql=%v, err=%w", err, sql)
	}

	//_, err := r.db.Exec(ctx, addTick,
	//	stockName,
	//	tick.InstrumentId,
	//	tick.Date,
	//	tick.Max,
	//	tick.Min,
	//	tick.Avg,
	//	tick.Median,
	//)

	//vals, _ := d.Values()
	//values := append(vals, d.AvgSquarePrice, d.CreatedAt, d.LatestDeveloperItemStartTime)

	//queryBuilder := sq.StatementBuilder.
	//	PlaceholderFormat(sq.Dollar).
	//	Insert("development").
	//	Columns(columns...).
	//	Values(values...)
	//query, args, err := queryBuilder.ToSql()
	//require.NoError(t, err)

	_, err = r.db.Exec(context.Background(), sql, args...)
	if err != nil {
		return fmt.Errorf("error inserting tick, sql=%v, args=%+v, err=%w", sql, args, err)
	}
	return err
}

func (r *Repo) Close() {
}

func (r *Repo) Flush(_ context.Context) error {
	return nil
}

func (r *Repo) Prepare(ctx context.Context, startDate time.Time, endDate time.Time) error {
	stmt := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("stock_name", "instrument_id", "tick_date", "max", "min", "avg", "median").
		From("tick1d").
		Where(squirrel.GtOrEq{"tick_date": startDate.UnixMilli()}).
		Where(squirrel.Lt{"tick_date": endDate.UnixMilli()}).
		Where(squirrel.NotEq{"instrument_id": "0"})
	query, args, err := stmt.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL, %w", err)
	}
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query, %w", err)
	}
	r.rows = &rows

	return nil
}

func (r *Repo) Scan() (any, error) {
	tick := dto.Tick1Day{}
	if r.rows == nil {
		return tick, fmt.Errorf("scan before initialization: call Prepare first")
	}

	var stockName string
	err := (*r.rows).Scan(&stockName, &tick.InstrumentId, &tick.Date, &tick.Max, &tick.Min, &tick.Avg, &tick.Median)
	if err != nil {
		return nil, err
	}
	tick.StockName = 1 // 1 for "NSE"
	if stockName == "BSE" {
		tick.StockName = 2 // 2 for "BSE"
	}

	return tick, nil
}

func (r *Repo) Next() bool {
	if r.rows == nil {
		log.Fatal("Scan before initialization: call Prepare first")
	}
	return (*r.rows).Next()
}
