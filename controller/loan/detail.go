package loan

import (
	"billing_engine/library/response"
	loanModel "billing_engine/model/loan"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *controller) Detail(ctx *gin.Context) {
	var reqPath loanModel.DetailReqPath
	if err := ctx.ShouldBindUri(&reqPath); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resData, statusCode, err := c.loanService.Detail(&reqPath)
	if err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Success(ctx, statusCode, "", resData)

}
