package store

import (
	"context"
	"github.com/ggchangan/yugong/internal/apiserver/model"
)

// ReportStore defines the user storage interface.
type ReportStore interface {
	Create(ctx context.Context, report *model.Report) error
	Update(ctx context.Context, report *model.Report) error
	Delete(ctx context.Context, id uint64) error
	Get(ctx context.Context, id uint64) (*model.Report, error)
}
