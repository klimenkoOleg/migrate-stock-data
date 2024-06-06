package uploader

import (
	"context"
)

type Service interface {
	ImportFromStorageToDatabase(ctx context.Context) error
}
