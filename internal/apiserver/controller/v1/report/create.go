package report

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/ggchangan/yugong/internal/apiserver/model"
	"github.com/ggchangan/yugong/internal/pkg/code"
	"github.com/ggchangan/yugong/internal/pkg/util/core"
)

func (r *ReportController) Create(c *gin.Context) {
	log.L(c).Info("report create function called.")

	var re model.Report

	if err := c.ShouldBindJSON(&re); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	//if errs := r.Validate(); len(errs) != 0 {
	//	core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

	//	return
	//}

	// Insert the user to the storage.
	if err := r.srv.Reports().Create(c, &re); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, r)
}
