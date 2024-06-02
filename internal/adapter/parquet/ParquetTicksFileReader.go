package parquet

import (
	"fmt"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/dto"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/source"
	"time"
)

type TickParquetReader struct {
	fr source.ParquetFile
	pr *reader.ParquetReader
	//totalCount int64
	//num       int
	startDate *time.Time
	endDate   *time.Time
}

func NewTickParquetReader(fileName string, startDate, endDate *time.Time) (*TickParquetReader, error) {
	fr, err := local.NewLocalFileReader(fileName)
	if err != nil {
		return nil, fmt.Errorf("can't open file, %w", err)

	}

	pr, err := reader.NewParquetReader(fr, new(dto.Tick1Day), 4)
	if err != nil {
		return nil, fmt.Errorf("can't create parquet reader, %w", err)
	}

	//startDT := datetime.Datetime{Month: 1, Day: 1, Year: 2024}
	//start := startDT.Time()
	//endDT := datetime.Datetime{Month: 5, Day: 5, Year: 2024}
	//end := endDT.Time()

	return &TickParquetReader{fr: fr, pr: pr, startDate: startDate, endDate: endDate}, nil
}

//func (p *TickParquetReader) HasNext() bool {
//	return p.num > 0
//}

func (p *TickParquetReader) Read(ticks []dto.Tick1Day) (int, error) {
	//if p.num <= 0 {
	//	return nil, errors.New("no data")
	//}

	//for p.num > 0 {
	//ticks = make([]dto.Tick1Day, 1_000)
	err := p.pr.Read(&ticks)
	if err != nil {
		return 0, fmt.Errorf("read error, %w", err)
	}
	//p.num--

	target := 0

	for _, tick := range ticks {
		date := time.UnixMilli(tick.Date)
		if p.startDate != nil && p.endDate != nil && date.After(*p.startDate) && date.Before(*p.endDate) {
			ticks[target] = tick
			target++
			//return &ticks[0], nil
		}
	}

	//tick := ticks[0]
	//date := time.UnixMilli(tick.Date)
	//if p.startDate != nil && p.endDate != nil && date.After(*p.startDate) && date.Before(*p.endDate) {
	//	return &ticks[0], nil
	//}
	//}

	return target, nil
}

func (p *TickParquetReader) Close() error {
	p.pr.ReadStop()
	return p.fr.Close()
}
