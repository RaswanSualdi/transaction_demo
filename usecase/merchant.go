package usecase

import (
	"fmt"
	"golang-test/models"
	"golang-test/payload/request"
	"math/rand"
)

func (u *usecase) RegisterMerchant(input request.RegisterMerchant) (models.Merchant, error) {

	randomNumber := fmt.Sprintf("%d", rand.Intn(900000)+100000)
	bank := models.Bank{
		BankName:      input.Bank.BankName,
		Saldo:         input.Bank.Saldo,
		AccountNumber: randomNumber,
	}
	merchant := models.Merchant{
		Name:         input.Name,
		Address:      input.Address,
		MerchantCode: randomNumber,
		Bank:         bank,
	}

	newMerchant, err := u.repositories.CreateMerchant(merchant)
	if err != nil {
		return models.Merchant{}, fmt.Errorf("failed to create merchant: %w", err)
	}

	return newMerchant, nil
}
