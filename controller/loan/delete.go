package loan

import (
	"billing_engine/common"
	"billing_engine/library/response"
	"billing_engine/model/loan"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *controller) Delete(ctx *gin.Context) {
	var reqPath loan.DeleteReqPath
	if err := ctx.ShouldBindUri(&reqPath); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := c.loanService.Delete(&reqPath); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SuccessfullyDeleted, nil)
}
