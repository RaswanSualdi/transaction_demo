package usecase

import (
	"golang-test/models"
	"golang-test/payload/request"
	"golang-test/repository"

	"github.com/golang-jwt/jwt/v5"
)

type Usecase interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	Login(input request.Login) (models.Customer, error)
	RegisterCustomer(input request.RegisterCustomer) (models.Customer, error)
	GetCustomerById(userID int) (models.Customer, error)
	RegisterMerchant(input request.RegisterMerchant) (models.Merchant, error)
	Transaction(input request.TransactionRequest) (models.Bank, models.Customer, models.Merchant, error)
}

type usecase struct {
	repositories repository.Repositories
}

func NewUseCase(repository repository.Repositories) *usecase {
	return &usecase{repositories: repository}
}
