package service

import "github.com/ggchangan/yugong/internal/apiserver/store"

//go:generate mockgen -self_package=github.com/marmotedu/iam/internal/apiserver/service/v1 -destination mock_service.go -package v1 github.com/marmotedu/iam/internal/apiserver/service/v1 Service,UserSrv,SecretSrv,PolicySrv

// Service defines functions used to return resource interface.
type Service interface {
	Reports() ReportSrv
	ReportMessages() ReportMessageSrv
}

type service struct {
	store store.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Reports() ReportSrv {
	return newReports(s)
}

func (s *service) ReportMessages() ReportMessageSrv {
	return newReportMessages(s)
}
