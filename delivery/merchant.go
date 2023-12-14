package delivery

import (
	"golang-test/payload/request"
	"golang-test/payload/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *delivery) RegisterMerchant(c *gin.Context) {
	var input request.RegisterMerchant
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "payload error",
		})
		return
	}
	merchant, err := d.usecase.RegisterMerchant(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.APIresponse("Failed register merchant", http.StatusBadRequest, "error", merchant))
		return
	}
	c.JSON(http.StatusCreated, response.APIresponse("Success created merchant", http.StatusCreated, "success", merchant))
}
