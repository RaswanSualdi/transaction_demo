package repository

import (
	"golang-test/models"

	"gorm.io/gorm"
)

type Repositories interface {
	FindCustomerById(customerID int) (models.Customer, error)
	// SaveTransaction(transactioReq request.TransactionRequest)
	FindCustomerByEmail(email string) (models.Customer, error)
	DeleteCustomerById(ID int) (models.Customer, error)
	UpdateCustomerByID(customer models.Customer) (models.Customer, error)
	CreateCustomer(customer models.Customer) (models.Customer, error)
	CreateMerchant(merchant models.Merchant) (models.Merchant, error)
	CreateBank(bank models.Bank) (models.Bank, error)
	// FindCustomerByBankID(bankID int) models.Bank
	FindMerchantByAccountBank(accountBank string) (models.Merchant, error)
	FindBankByAccountNumber(accountBank string) (models.Bank, error)
	UpdateAccountBank(bank models.Bank, amount int64) (models.Bank, error)
}

type repositories struct {
	db *gorm.DB
}

func NewRepositories(db *gorm.DB) *repositories {
	return &repositories{db}
}
