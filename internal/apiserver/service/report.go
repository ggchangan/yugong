package service

import (
	"context"

	"github.com/ggchangan/yugong/internal/pkg/code"
	"github.com/marmotedu/errors"

	"github.com/ggchangan/yugong/internal/apiserver/model"
	"github.com/ggchangan/yugong/internal/apiserver/store"
)

// ReportSrv defines functions used to handle report request.
type ReportSrv interface {
	Create(ctx context.Context, report *model.Report) error
	Update(ctx context.Context, report *model.Report) error
	Delete(ctx context.Context, id uint64) error
	Get(ctx context.Context, id uint64) (*model.Report, error)
}

type reportService struct {
	store store.Factory
}

func newReports(srv *service) *reportService {
	return &reportService{store: srv.store}
}

var _ ReportSrv = (*reportService)(nil)

func (r reportService) Create(ctx context.Context, report *model.Report) error {
	if err := r.store.Reports().Create(ctx, report); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (r reportService) Update(ctx context.Context, report *model.Report) error {
	if err := r.store.Reports().Update(ctx, report); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (r reportService) Delete(ctx context.Context, id uint64) error {
	if err := r.store.Reports().Delete(ctx, id); err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

func (r reportService) Get(ctx context.Context, id uint64) (*model.Report, error) {
	report, err := r.store.Reports().Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return report, nil
}
