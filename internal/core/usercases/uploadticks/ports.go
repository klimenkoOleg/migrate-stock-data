package uploadticks

import (
	"context"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/dto"
)

type ObjectReader interface {
	Read(*[]dto.Tick1Day) (int, error)
	Close() error
}

type Repo interface {
	WriteTick(tick []dto.Tick1Day) error
	Flush(ctx context.Context) error
	Close() error
}
