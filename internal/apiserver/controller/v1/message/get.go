package message

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/ggchangan/yugong/internal/pkg/util/core"
)

// Get get an report message by the report identifier.
func (r *ReportMessageController) Get(c *gin.Context) {
	log.L(c).Info("get report message function called.")

	user, err := r.srv.ReportMessages().Get(c, c.GetUint64("id"))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}
