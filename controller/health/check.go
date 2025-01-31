package health

import (
	"billing_engine/common"
	"billing_engine/lib/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *controller) Check(ctx *gin.Context) {
	if err := c.healthService.Check(); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SuccessfullyChecked, nil)

	return
}
