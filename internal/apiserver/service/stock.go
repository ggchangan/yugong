package service

import (
	"context"
	"github.com/ggchangan/yugong/internal/apiserver/model"
	"github.com/ggchangan/yugong/internal/apiserver/store"
)

type StockSrv interface {
	Get(ctx context.Context, id uint64) (*model.Stock, error)
}

type stockService struct {
	store store.Factory
}

func newStockService(srv *service) *stockService {
	return &stockService{store: srv.store}
}

var _ StockSrv = (*stockService)(nil)

func (s stockService) Get(ctx context.Context, id uint64) (*model.Stock, error) {
	stock, err := s.store.Stocks().Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return stock, nil
}
