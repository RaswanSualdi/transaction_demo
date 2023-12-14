package response

import (
	"golang-test/models"
	"golang-test/payload/request"
)

type TransactionResponse struct {
	From          string `json:"from"`
	Email         string `json:"email"`
	To            string `json:"to"`
	Amount        int64  `json:"Amount"`
	BankName      string `json:"bankName"`
	AccountNumber string `json:"accountNumber"`
}

func FormatTransactionResponse(bank models.Bank, customer models.Customer, merchant models.Merchant, request request.TransactionRequest) TransactionResponse {
	transactionResponse := TransactionResponse{
		From:          customer.Name,
		Email:         customer.Email,
		To:            merchant.Name,
		Amount:        request.Amount,
		BankName:      bank.BankName,
		AccountNumber: bank.AccountNumber,
	}

	return transactionResponse

}
