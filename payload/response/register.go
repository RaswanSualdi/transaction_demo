package response

import "golang-test/models"

type RegisterResponse struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

func FormatRegisterResponse(customer models.Customer) RegisterResponse {
	registerResponse := RegisterResponse{
		UserName: customer.Name,
		Email:    customer.Email,
	}

	return registerResponse
}
