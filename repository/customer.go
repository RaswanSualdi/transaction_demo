package repository

import "golang-test/models"

func (r *repositories) CreateCustomer(customer models.Customer) (models.Customer, error) {
	err := r.db.Create(&customer).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *repositories) FindCustomerById(customerID int) (models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("ID = ?", customerID).Find(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *repositories) FindCustomerByEmail(email string) (models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("email = ?", email).Find(&customer).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *repositories) DeleteCustomerById(ID int) (models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("id = ?", ID).Delete(&customer).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (r *repositories) UpdateCustomerByID(customer models.Customer) (models.Customer, error) {
	err := r.db.Save(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}
