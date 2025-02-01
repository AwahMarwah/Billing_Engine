package loan

import (
	"billing_engine/library/pagination"
	"billing_engine/library/response"
	loanModel "billing_engine/model/loan"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *controller) List(ctx *gin.Context) {
	var reqQuery loanModel.ListReqQuery
	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	reqQuery.Offset = pagination.Offset(&reqQuery.Limit, &reqQuery.Page)
	resData, count, err := c.loanService.List(&reqQuery)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessWithPage(ctx, http.StatusOK, "", resData, reqQuery.Page, reqQuery.Limit, count)
}
