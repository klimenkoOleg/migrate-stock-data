package uploadticks

import (
	"context"
	"fmt"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/dto"
)

type TickDumperService struct {
	repo         Repo
	objectReader ObjectReader
	batchSize    int
	maxRecords   int
}

func NewTickDumperService(Repo Repo, ObjectReader ObjectReader, batchSize, maxRecords int) *TickDumperService {
	return &TickDumperService{repo: Repo, objectReader: ObjectReader, batchSize: batchSize, maxRecords: maxRecords}
}

func (s *TickDumperService) ImportFromStorageToDatabase(ctx context.Context) error {

	ticks := make([]dto.Tick1Day, s.batchSize)
	count := 1

	for {
		//ticks := make([]dto.Tick1Day, len())
		numRead, err := s.objectReader.Read(ticks)
		if err != nil {
			return fmt.Errorf("object reader failed, %w", err)
		}
		if numRead == 0 {
			break
		}
		//if err != nil {
		//	if errors.Is(err, io.EOF) {
		//		break
		//	}
		//	return fmt.Errorf("object reader failed, %w", err)
		//}

		if count == 1 {
			fmt.Printf("%+v\n", ticks[0])
		}

		ticks = ticks[:numRead]

		//ticks = append(ticks, *tick)

		//if count%s.batchSize == 0 {
		err = s.repo.WriteTick(ctx, ticks)
		if err != nil {
			return fmt.Errorf("repo write failed, %w", err)
		}
		ticks = ticks[:0]
		//}

		//if count > s.maxRecords {
		//	break
		//}
		count++
	}

	return nil
}
