package fake

import (
	"context"

	"github.com/marmotedu/errors"

	"github.com/ggchangan/yugong/internal/apiserver/model"
	"github.com/ggchangan/yugong/internal/pkg/code"
	reflectutil "github.com/ggchangan/yugong/internal/pkg/util/reflect"
)

type reports struct {
	ds *datastore
}

func newReports(ds *datastore) *reports {
	return &reports{ds: ds}
}

func (r reports) Create(ctx context.Context, report *model.Report) error {
	r.ds.Lock()
	defer r.ds.Unlock()

	for _, re := range r.ds.reports {
		if report.Name == re.Name {
			return errors.WithCode(code.ErrUserAlreadyExist, "record already exist")
		}
	}

	if len(r.ds.reports) > 0 {
		report.ID = r.ds.reports[len(r.ds.reports)-1].ID + 1
	}
	r.ds.reports = append(r.ds.reports, report)

	return nil
}

func (r reports) Update(ctx context.Context, report *model.Report) error {
	r.ds.Lock()
	defer r.ds.Unlock()

	for _, re := range r.ds.reports {
		if report.Name == re.Name {
			if _, err := reflectutil.CopyObj(report, re, nil); err != nil {
				return errors.Wrap(err, "copy user failed")
			}
		}
	}

	return nil
}

func (r reports) Delete(ctx context.Context, id uint64) error {
	r.ds.Lock()
	defer r.ds.Unlock()

	reports := r.ds.reports
	r.ds.reports = make([]*model.Report, 0)
	for _, re := range reports {
		if re.ID == id {
			continue
		}

		r.ds.reports = append(r.ds.reports, re)
	}

	return nil
}

func (r reports) Get(ctx context.Context, id uint64) (*model.Report, error) {
	r.ds.RLock()
	defer r.ds.RUnlock()

	for _, re := range r.ds.reports {
		if re.ID == id {
			return re, nil
		}
	}

	return nil, errors.WithCode(code.ErrUserNotFound, "record not found")
}
