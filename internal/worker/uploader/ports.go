package uploader

import (
	"context"
)

type service interface {
	ImportFromStorageToDatabase(ctx context.Context) error
}
