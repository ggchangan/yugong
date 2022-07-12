package service

import (
	"context"

	"github.com/ggchangan/yugong/internal/apiserver/model"
	"github.com/ggchangan/yugong/internal/apiserver/store"
)

type ReportMessageSrv interface {
	Create(ctx context.Context, report *model.ReportMessage) error
	Update(ctx context.Context, report *model.ReportMessage) error
	Delete(ctx context.Context, id uint64) error
	Get(ctx context.Context, id uint64) (*model.ReportMessage, error)
}

type reportMessageService struct {
	store store.Factory
}

func (r reportMessageService) Create(ctx context.Context, report *model.ReportMessage) error {
	panic("implement me")
}

func (r reportMessageService) Update(ctx context.Context, report *model.ReportMessage) error {
	panic("implement me")
}

func (r reportMessageService) Delete(ctx context.Context, id uint64) error {
	panic("implement me")
}

func (r reportMessageService) Get(ctx context.Context, id uint64) (*model.ReportMessage, error) {
	panic("implement me")
}

func newReportMessages(srv *service) *reportMessageService {
	return &reportMessageService{store: srv.store}
}

var _ ReportMessageSrv = (*reportMessageService)(nil)
