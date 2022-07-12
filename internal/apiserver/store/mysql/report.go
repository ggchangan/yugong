package mysql

import (
	"context"

	gorm "gorm.io/gorm"

	"github.com/marmotedu/errors"

	"github.com/ggchangan/yugong/internal/apiserver/model"
	"github.com/ggchangan/yugong/internal/pkg/code"
)

type reports struct {
	db *gorm.DB
}

func newReports(ds *datastore) *reports {
	return &reports{db: ds.db}
}

// Create creates a new report account.
func (u *reports) Create(ctx context.Context, report *model.Report) error {
	return u.db.Create(&report).Error
}

// Update updates an report account information.
func (u *reports) Update(ctx context.Context, report *model.Report) error {
	return u.db.Save(report).Error
}

// Delete deletes the report by the report identifier.
func (u *reports) Delete(ctx context.Context, id uint64) error {
	err := u.db.Where("id = ?", id).Delete(&model.Report{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	return nil
}

// Get return an report by the report identifier.
func (u *reports) Get(ctx context.Context, id uint64) (*model.Report, error) {
	report := &model.Report{}
	err := u.db.Where("id = ?", id).First(&report).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return report, nil
}
