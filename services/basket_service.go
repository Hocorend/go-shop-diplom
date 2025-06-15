package services

import (
	"go-shop-diplom/models"
	"go-shop-diplom/storage"
)

func FindBasketForUserAndProduct(user *models.User, product *models.Product) (*models.Basket, error) {
	basket := &models.Basket{}
	err := storage.DB.Where("user_id = ? AND product_id = ?", user.ID, product.ID).First(basket).Error
	if err != nil {
		return nil, err
	}

	return basket, nil
}

func UpdateBasket(basket *models.Basket) error {
	err := storage.DB.Model(&models.Basket{}).Where("id = ?", basket.ID).Updates(basket).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateBasket(basket *models.Basket) error {
	err := storage.DB.Create(basket).Error
	if err != nil {
		return err
	}
	return nil
}

func FindAllBasketsForUser(user *models.User) ([]models.Basket, error) {
	var baskets []models.Basket
	err := storage.DB.Preload("Product").Preload("User").Where("user_id = ?", user.ID).Find(&baskets).Error
	if err != nil {
		return nil, err
	}
	return baskets, nil
}

func DeleteAllBasketsForUser(user *models.User) error {
	err := storage.DB.Where("user_id = ?", user.ID).Delete(&models.Basket{}).Error
	if err != nil {
		return err
	}
	return nil
}
