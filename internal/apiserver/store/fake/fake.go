// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fake

import (
	"fmt"
	"sync"

	"github.com/ggchangan/yugong/internal/apiserver/model"
	"github.com/ggchangan/yugong/internal/apiserver/store"
)

// ResourceCount defines the number of fake resources.
const ResourceCount = 1000

type datastore struct {
	sync.RWMutex
	reports        []*model.Report
	reportMessages []*model.ReportMessage
	stocks         []*model.Stock
}

func (ds *datastore) Stocks() store.StockStore {
	return newStocks(ds)
}

func (ds *datastore) Reports() store.ReportStore {
	return newReports(ds)
}

func (ds *datastore) ReportMessages() store.ReportMessageStore {
	return newReportMessages(ds)
}

func (ds *datastore) Close() error {
	return nil
}

var (
	fakeFactory store.Factory
	once        sync.Once
)

// GetFakeFactoryOr create fake store.
func GetFakeFactoryOr() (store.Factory, error) {
	once.Do(func() {
		fakeFactory = &datastore{
			reports:        FakeReports(ResourceCount),
			reportMessages: FakeReportMessages(ResourceCount),
			stocks:         FakeStocks(ResourceCount),
		}
	})

	if fakeFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v", fakeFactory)
	}

	return fakeFactory, nil
}

// FakeReports returns fake report data.
func FakeReports(count int) []*model.Report {
	// init some report records
	reports := make([]*model.Report, 0)
	for i := 1; i <= count; i++ {
		reports = append(reports, &model.Report{
			ObjectMeta: model.ObjectMeta{
				Name: fmt.Sprintf("report%d", i),
				ID:   uint64(i),
			},
			CreateUser: fmt.Sprintf("user%d", i),
		})
	}

	return reports
}

// FakeReportMessages returns fake report message data.
func FakeReportMessages(count int) []*model.ReportMessage {
	// init some report records
	reportMessages := make([]*model.ReportMessage, 0)
	for i := 1; i <= count; i++ {
		reportMessages = append(reportMessages, &model.ReportMessage{
			ObjectMeta: model.ObjectMeta{
				Name: fmt.Sprintf("reportMessage%d", i),
				ID:   uint64(i),
			},
			CreateUser: fmt.Sprintf("user%d", i),
		})
	}

	return reportMessages
}

// FakeStocks returns fake stock data.
func FakeStocks(count int) []*model.Stock {
	// init some report records
	stocks := make([]*model.Stock, 0)
	for i := 1; i <= count; i++ {
		stocks = append(stocks, &model.Stock{
			ObjectMeta: model.ObjectMeta{
				Name: fmt.Sprintf("report%d", i),
				ID:   uint64(i),
			},
			Price: float32(i),
		})
	}

	return stocks
}
