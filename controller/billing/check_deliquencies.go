package billing

import (
	"billing_engine/library/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *controller) CheckDeliquencies(ctx *gin.Context) {
	resData, err := c.billingService.Check()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "", resData)
}
