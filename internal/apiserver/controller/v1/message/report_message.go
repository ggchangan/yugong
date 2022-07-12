package message

import (
	"github.com/ggchangan/yugong/internal/apiserver/service"
	"github.com/ggchangan/yugong/internal/apiserver/store"
)

// ReportMessageController create a report message handler used to handle request for report message resource.
type ReportMessageController struct {
	srv service.Service
}

// NewReportMessageController creates a user handler.
func NewReportMessageController(store store.Factory) *ReportMessageController {
	return &ReportMessageController{
		srv: service.NewService(store),
	}
}
