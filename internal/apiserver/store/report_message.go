package store

import (
	"context"

	"github.com/ggchangan/yugong/internal/apiserver/model"
)

// ReportMessageStore defines the user storage interface.
type ReportMessageStore interface {
	Create(ctx context.Context, report *model.ReportMessage) error
	Update(ctx context.Context, report *model.ReportMessage) error
	Delete(ctx context.Context, id uint64) error
	Get(ctx context.Context, id uint64) (*model.ReportMessage, error)
}
