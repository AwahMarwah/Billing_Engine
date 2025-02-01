package loan

import (
	"billing_engine/common"
	"billing_engine/library/response"
	loanModel "billing_engine/model/loan"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (c *controller) Create(ctx *gin.Context) {
	var reqBody loanModel.CreateReqBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(reqBody, "loan controller")
	if statusCode, err := c.loanService.Create(&reqBody); err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SuccessfullyCreated, nil)
}
