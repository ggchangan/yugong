package report

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/ggchangan/yugong/internal/pkg/util/core"
)

// Delete delete an report by the report identifier.
// Only administrator can call this function.
func (r *ReportController) Delete(c *gin.Context) {
	log.L(c).Info("delete user function called.")

	if err := r.srv.Reports().Delete(c, c.GetUint64("id")); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
