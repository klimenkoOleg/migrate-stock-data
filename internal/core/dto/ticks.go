package dto

// to exchange between handler/controller/command and service
// Services, Use cases should not convert to/from this DTOs. They should use domain models.
// This is a common practice in DDD.
// Adapters are responsible for converting domain models to/from DTOs.
type Tick struct {
	InstrumentId int64 `parquet:"name=InstrumentId, type=INT64"`
	StockName    int32 `parquet:"name=StockName, type=INT32, encoding=PLAIN"`
	Timestamp    int64 `parquet:"name=Timestamp, type=INT64"`
	Value        int64 `parquet:"name=Value, type=INT64"`
}

type Tick1Day struct {
	StockName    int32   `parquet:"name=Stock, type=INT32, encoding=PLAIN"`
	InstrumentId int64   `parquet:"name=instrument_id, type=INT64"`
	Date         int64   `parquet:"name=date, type=INT64"`
	Max          int64   `parquet:"name=Value, type=INT64"`
	Min          int64   `parquet:"name=min_value, type=INT64"`
	Avg          float64 `parquet:"name=avg_value, type=DOUBLE"`
	Median       float64 `parquet:"name=median_value, type=DOUBLE"`
}
