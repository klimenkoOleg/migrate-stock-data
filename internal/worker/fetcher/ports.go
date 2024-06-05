package fetcher

import (
	"context"
	"time"
)

type service interface {
	FetchFromDBtoParquet(ctx context.Context, startDate time.Time, endDate time.Time) error
}
