package delivery

import (
	"golang-test/payload/request"
	"golang-test/payload/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *delivery) Transaction(c *gin.Context) {
	var input request.TransactionRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "payload error",
		})
		return
	}
	bank, customer, merchant, err := d.usecase.Transaction(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.APIresponse("Failed register merchant",
			http.StatusBadRequest, "error", response.FormatTransactionResponse(bank, customer, merchant, input)))
		return
	}

	c.JSON(http.StatusOK, response.APIresponse("Success Transfer",
		http.StatusOK, "success", response.FormatTransactionResponse(bank, customer, merchant, input)))
}
