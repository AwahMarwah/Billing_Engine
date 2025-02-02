package billing

import (
	"billing_engine/library/response"
	"billing_engine/model/billing"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *controller) GetOutstanding(ctx *gin.Context) {
	var reqPath billing.DetailReqPath
	if err := ctx.ShouldBindUri(&reqPath); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resData, statusCode, err := c.billingService.Detail(&reqPath)
	if err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Success(ctx, statusCode, "", resData)
}
