package borrower

import (
	"billing_engine/common"
	"billing_engine/library/response"
	borrowerModel "billing_engine/model/borrower"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (c *controller) Create(ctx *gin.Context) {
	var reqBody borrowerModel.CreateReqBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(reqBody, "controller")
	if err := c.borrowerService.Create(&reqBody); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SuccessfullyCreated, nil)
}
