package borrower

import (
	"billing_engine/library/response"
	borrowerModel "billing_engine/model/borrower"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *controller) Detail(ctx *gin.Context) {
	var reqPath borrowerModel.DetailReqPath
	if err := ctx.ShouldBindUri(&reqPath); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resData, statusCode, err := c.borrowerService.Detail(&reqPath)
	if err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Success(ctx, statusCode, "", resData)
}
