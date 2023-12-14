package usecase

import (
	"errors"
	"golang-test/models"
	"golang-test/payload/request"

	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) RegisterCustomer(input request.RegisterCustomer) (models.Customer, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	newCustomer := models.Customer{
		Name:         input.UserName,
		Email:        input.Email,
		PasswordHash: string(passwordHash),
	}

	findCustomerByEmail, _ := u.repositories.FindCustomerByEmail(input.Email)
	if len(findCustomerByEmail.Email) > 0 {
		return findCustomerByEmail, errors.New("Account has been registered")
	}

	if input.Password != input.ConfirmPassword {
		return findCustomerByEmail, errors.New("Confirm password wrong")
	}
	customer, err := u.repositories.CreateCustomer(newCustomer)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (u *usecase) GetCustomerById(userID int) (models.Customer, error) {
	customer, err := u.repositories.FindCustomerById(userID)
	if err != nil {
		return customer, err
	}
	return customer, nil
}
