package uploadticks

import (
	"context"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/dto"
)

type ObjectReader interface {
	//HasNext() bool
	Read([]dto.Tick1Day) (int, error)
	Close() error
}

type Repo interface {
	WriteTick(ctx context.Context, tick []dto.Tick1Day) error
	Close()
	Flush(ctx context.Context) error
}
