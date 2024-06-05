package fetchticks

import (
	"context"
	"fmt"
	"time"
)

type TickFetcherService struct {
	repo         ClickhouseReader
	objectWriter ObjectWriter
}

func NewTickFetcherService(repo ClickhouseReader, objectWriter ObjectWriter) *TickFetcherService {
	return &TickFetcherService{repo: repo, objectWriter: objectWriter}
}

type ClickhouseReader interface {
	Next() bool
	Scan() (any, error)
	Prepare(ctx context.Context, startDate time.Time, endDate time.Time) error
}

func (s *TickFetcherService) FetchFromDBtoParquet(ctx context.Context, startDate time.Time, endDate time.Time) error {
	err := s.repo.Prepare(ctx, startDate, endDate)
	if err != nil {
		return fmt.Errorf("failed fetch, %w", err)
	}

	count := 1
	for s.repo.Next() {
		data, err := s.repo.Scan()
		if err != nil {
			return fmt.Errorf("failed scan, %w", err)
		}

		err = s.objectWriter.Write(data)
		if err != nil {
			return fmt.Errorf("filed write to storage, %w", err)
		}
		if count%1000 == 0 {
			fmt.Printf("count: %d\n", count)
		}

		count++
	}

	return nil
}
