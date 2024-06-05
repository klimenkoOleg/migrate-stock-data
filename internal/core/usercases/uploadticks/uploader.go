package uploadticks

import (
	"context"
	"errors"
	"fmt"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/dto"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"io"
)

type TickUploaderService struct {
	repo         Repo
	objectReader ObjectReader
	batchSize    int
}

func NewTickDumperService(Repo Repo, ObjectReader ObjectReader, batchSize int) *TickUploaderService {
	return &TickUploaderService{repo: Repo, objectReader: ObjectReader, batchSize: batchSize}
}

func (s *TickUploaderService) ImportFromStorageToDatabase(ctx context.Context) error {
	ticks := make([]dto.Tick1Day, 0, s.batchSize)
	for {
		ticks = ticks[:cap(ticks)]
		numRead, err := s.objectReader.Read(&ticks)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("object reader failed, %w", err)
		}
		ticks = ticks[:numRead]
		err = s.repo.WriteTick(ticks)
		if err != nil {
			return fmt.Errorf("repo write failed, %w", err)
		}
	}
	if err := s.repo.Flush(ctx); err != nil {
		log.Error("repo flush failed", zap.Error(err))
	}

	return nil
}
