package fake

import (
	"context"

	"github.com/marmotedu/errors"

	"github.com/ggchangan/yugong/internal/apiserver/model"
	"github.com/ggchangan/yugong/internal/pkg/code"
	reflectutil "github.com/ggchangan/yugong/internal/pkg/util/reflect"
)

type reportMessages struct {
	ds *datastore
}

func (r reportMessages) Create(ctx context.Context, report *model.ReportMessage) error {
	r.ds.Lock()
	defer r.ds.Unlock()

	for _, re := range r.ds.reportMessages {
		if report.Name == re.Name {
			return errors.WithCode(code.ErrUserAlreadyExist, "record already exist")
		}
	}

	if len(r.ds.reportMessages) > 0 {
		report.ID = r.ds.reportMessages[len(r.ds.reportMessages)-1].ID + 1
	}
	r.ds.reportMessages = append(r.ds.reportMessages, report)

	return nil
}

func (r reportMessages) Update(ctx context.Context, report *model.ReportMessage) error {
	r.ds.Lock()
	defer r.ds.Unlock()

	for _, re := range r.ds.reportMessages {
		if report.Name == re.Name {
			if _, err := reflectutil.CopyObj(report, re, nil); err != nil {
				return errors.Wrap(err, "copy user failed")
			}
		}
	}

	return nil
}

func (r reportMessages) Delete(ctx context.Context, id uint64) error {
	r.ds.Lock()
	defer r.ds.Unlock()

	reports := r.ds.reportMessages
	r.ds.reportMessages = make([]*model.ReportMessage, 0)
	for _, re := range reports {
		if re.ID == id {
			continue
		}

		r.ds.reportMessages = append(r.ds.reportMessages, re)
	}

	return nil
}

func (r reportMessages) Get(ctx context.Context, id uint64) (*model.ReportMessage, error) {
	r.ds.RLock()
	defer r.ds.RUnlock()

	for _, re := range r.ds.reportMessages {
		if re.ID == id {
			return re, nil
		}
	}

	return nil, errors.WithCode(code.ErrUserNotFound, "record not found")
}

func newReportMessages(ds *datastore) *reportMessages {
	return &reportMessages{ds: ds}
}
