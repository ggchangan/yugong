package mysql

import (
	"context"
	"github.com/ggchangan/yugong/internal/apiserver/model"
	"github.com/ggchangan/yugong/internal/pkg/code"
	"github.com/marmotedu/errors"
	"gorm.io/gorm"
)

type stocks struct {
	db *gorm.DB
}

func newStocks(ds *datastore) *stocks {
	return &stocks{db: ds.db}
}

// Get return an report by the report identifier.
func (u *stocks) Get(ctx context.Context, id uint64) (*model.Stock, error) {
	stock := &model.Stock{}
	err := u.db.Where("id = ?", id).First(&stock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrStockNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return stock, nil
}
