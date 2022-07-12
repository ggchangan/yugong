package report

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/ggchangan/yugong/internal/pkg/util/core"
)

// Get get an report by the report identifier.
func (r *ReportController) Get(c *gin.Context) {
	log.L(c).Info("get report function called.")

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user, err := r.srv.Reports().Get(c, uint64(id))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}
