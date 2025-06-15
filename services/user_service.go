package services

import (
	"go-shop-diplom/models"
	"go-shop-diplom/storage"
)

func UpdateDeposit(user *models.User) (*models.User, error) {

	err := storage.DB.Model(&models.User{}).Where("id = ?", user.ID).Update("deposit", user.Deposit).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserByLogin(login *string) (*models.User, error) {
	user := &models.User{}
	err := storage.DB.Where("login = ?", login).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil

}
