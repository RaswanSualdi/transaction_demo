package response

import (
	"golang-test/models"
)

type CustomerResponse struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"username"`
	Token    string `json:"token"`
}

func FormatLoginCustomerResponse(customer models.Customer, token string) CustomerResponse {
	customerResponse := CustomerResponse{
		UserID:   customer.ID,
		UserName: customer.Email,
		Token:    token,
	}

	return customerResponse
}
