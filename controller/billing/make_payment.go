package billing

import (
	"billing_engine/common"
	"billing_engine/library/response"
	billingModel "billing_engine/model/billing"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *controller) MakePayment(ctx *gin.Context) {
	var reqBody billingModel.CreateReqBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if reqBody.AmountPaid != 110000 {
		response.Error(ctx, http.StatusBadRequest, "amount_paid must be 110000")
		return
	}
	if err := c.billingService.Create(&reqBody); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SuccessfullyCreated, nil)
}
