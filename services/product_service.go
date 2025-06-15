package services

import (
	"go-shop-diplom/models"
	"go-shop-diplom/storage"
)

func FindAllProducts() (*[]models.Product, error) {
	products := &[]models.Product{}
	err := storage.DB.Find(products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func FindProductByName(name *string) (*models.Product, error) {
	product := &models.Product{}
	err := storage.DB.Where("name = ?", name).First(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
