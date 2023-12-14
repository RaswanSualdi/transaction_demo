package repository

import (
	"fmt"
	"golang-test/models"
)

func (r *repositories) CreateBank(bank models.Bank) (models.Bank, error) {
	err := r.db.Create(&bank).Error
	if err != nil {
		return bank, err
	}
	return bank, nil
}

func (r *repositories) FindBankByAccountNumber(accountBank string) (models.Bank, error) {
	var bank models.Bank
	err := r.db.Where("account_number = ?", accountBank).Find(&bank).Error
	if err != nil {
		return bank, err
	}
	return bank, nil
}

func (r *repositories) UpdateAccountBank(bank models.Bank, amount int64) (models.Bank, error) {
	account := bank.AccountNumber
	currentSaldo := bank.Saldo
	currentSaldo += amount
	err := r.db.Model(&models.Bank{}).Where("account_number = ?", account).Update("saldo", currentSaldo).Error
	if err != nil {
		return bank, err
	}
	fmt.Println(bank.Saldo)
	return bank, nil
}
