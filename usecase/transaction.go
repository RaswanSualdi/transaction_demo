package usecase

import (
	"fmt"
	"golang-test/models"
	"golang-test/payload/request"
)

func (u *usecase) Transaction(input request.TransactionRequest) (models.Bank, models.Merchant, error) {
	bank, _ := u.repositories.FindBankByAccountNumber(input.NumberAccount)

	// currentSaldo := bank.Saldo
	// fmt.Println(currentSaldo)
	// currentSaldo += input.Amount
	// fmt.Println(currentSaldo)
	updatedAccountBank, _ := u.repositories.UpdateAccountBank(bank, input.Amount)
	// fmt.Println(updatedAccountBank.Saldo)
	fmt.Println(updatedAccountBank.Saldo)
	merchant, _ := u.repositories.FindMerchantByAccountBank(input.NumberAccount)

	return updatedAccountBank, merchant, nil

}
