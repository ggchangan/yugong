package store

import (
	"context"
	"github.com/ggchangan/yugong/internal/apiserver/model"
)

// StockStore defines the stock storage interface.
type StockStore interface {
	Get(ctx context.Context, id uint64) (*model.Stock, error)
}
