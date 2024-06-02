package clickhouse

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/klimenkoOleg/migrate-stock-data/internal/adapter/parquet"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/dto"
	"time"
)

type TickParquetService struct {
}

func NewTickParquetService() *TickParquetService {
	return &TickParquetService{}
}

const layout = "2006-01-02"

func (s *Repo) WriteByMonth(ctx context.Context) error {

	t := time.Date(2023, time.November, 20, 0, 0, 0, 0, time.UTC)

	for i := 1; i <= 1; i++ {
		nextT := t.AddDate(0, 0, 10)
		//fmt.Println(t)
		startDate := fmt.Sprintf(t.Format(layout))
		endDate := fmt.Sprintf(nextT.Format(layout))

		fmt.Println(startDate)
		fmt.Println(endDate)

		//err := s.WriteSingleFile(ctx, t, nextT, startDate+".parquet")
		//if err != nil {
		//	return err
		//}
		//fmt.Println(fileName)
		t = nextT
	}

	return nil
}

func (s *Repo) CreateQuery(_ context.Context, startDate time.Time, endDate time.Time) squirrel.SelectBuilder {
	//startInt := startDate.UnixMilli()
	//endInt := endDate.UnixMilli()
	//select
	//	from tickers.tickers_aggregates_1day_mv where
	//	start_at >=  toDateTime('2023-01-01 00:00:00') AND
	//	start_at < toDateTime('2023-02-01 00:00:00')
	//	-- and instrument_id = '216328676147593246'
	//	group by stock_name,  instrument_id, start_at
	//	order by stock_name,  instrument_id, start_at
	//	limit 10;

	startDateStr := fmt.Sprintf(startDate.Format(layout))
	endDateStr := fmt.Sprintf(endDate.Format(layout))

	//stock_name int,
	//	instrument_id
	//tick_date bigint,
	//max_value bigint,
	//min_value bigint,
	//avg_value double precision,
	//med_value double precision

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

	return stmt
}

//func (s *Repo) WriteSingleFile(ctx context.Context, startDate string, endDate string, fileName string) error {
//}

/*
	func (s *Repo) Scan(builder *sql.Rows) (any, error) {
		tick := dto.Tick1Day{}
		date := time.Time{}
		err := builder.Scan(&tick.StockName, &tick.InstrumentId, &date, &tick.Max, &tick.Min, &tick.Avg, &tick.Median)
		tick.Date = date.UnixMilli()
		if err != nil {
			return nil, err
		}
		return tick, nil
	}
*/
func (s *Repo) WriteParquet() error {
	var err error

	writer, err := parquet.NewTickParquetWriter("test.parquet")
	if err != nil {
		return err
	}
	defer func() {
		errClose := writer.Close()
		if errClose != nil {
			err = errors.Join(err, errClose)
		}
	}()

	num := 100
	for i := 1; i < num; i++ {
		/*tick := &Tick{
			InstrumentId: int64(i),
			StockName:    int32(i%2) + 1,
			Timestamp:    int64(i*1000) + 1,
			Value:        int64(i+1010) + 1,
		}*/
		tick1 := &dto.Tick1Day{

			InstrumentId: int64(i),
			StockName:    int32(i%2) + 1,
			Date:         time.Now().UnixMilli(),
			Max:          10,
			Min:          1,
			Avg:          1,
			Median:       1,
		}

		err = writer.Write(tick1)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
func (s *Repo) FetchTickers5m(ctx context.Context, userID uint64) (string, error) {

	stmt := CLKH.
		Select("InstrumentId", "max_value").
		From("tickers.tickers_aggr_5min").
		//Where(squirrel.NotEq{"cid": nil}).
		//Where(squirrel.Eq{"user_id": userID}).
		//OrderBy("Timestamp DESC").
		Limit(10)

	query, args, err := stmt.ToSql()
	if err != nil {
		return "", err
	}
	//spanTag(trace.SpanFromContext(ctx), query, args)
	row := s.db.QueryRowContext(ctx, query, args...)

	var id, maxValue int
	err = row.Scan(&id, &maxValue)
	if err != nil {
		return "", err
	}
	//if cid == "" {
	//	return "", errors.New("cid is empty")
	//}

	return "ok", nil
}
*/
