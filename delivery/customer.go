package delivery

import (
	"golang-test/payload/request"
	"golang-test/payload/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *delivery) RegisterCustomer(c *gin.Context) {
	var input request.RegisterCustomer
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "payload error",
		})
		return
	}
	customer, err := d.usecase.RegisterCustomer(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIresponse("register failed", http.StatusBadRequest, "Error", response.FormatRegisterResponse(customer)))
		return
	}
	response := response.APIresponse("register success", http.StatusCreated, "success", response.FormatRegisterResponse(customer))
	c.JSON(http.StatusOK, response)
}

func (d *delivery) Login(c *gin.Context) {
	var input request.Login
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "payload error",
		})
		return
	}

	customer, err := d.usecase.Login(input)
	token, _ := d.usecase.GenerateToken(customer.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIresponse("login failed", http.StatusBadRequest, "Error", response.FormatLoginCustomerResponse(customer, token)))
		return
	}

	response := response.APIresponse("register success", http.StatusOK, "success", response.FormatLoginCustomerResponse(customer, token))
	c.JSON(http.StatusOK, response)

}
