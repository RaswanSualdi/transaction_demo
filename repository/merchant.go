package repository

import "golang-test/models"

func (r *repositories) CreateMerchant(merchant models.Merchant) (models.Merchant, error) {
	err := r.db.Create(&merchant).Error
	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

func (r *repositories) FindMerchantByAccountBank(accountBank string) (models.Merchant, error) {
	var merchant models.Merchant
	err := r.db.Joins("Bank").Where("account_number = ?", accountBank).First(&merchant).Error
	if err != nil {
		return merchant, err
	}
	return merchant, nil

}
