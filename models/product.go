package models

import (
	"go-shop-diplom/dto"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	ID   uint32  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name *string `json:"name"`
	Cost uint32  `json:"cost"`
}

func MigrateProduct(db *gorm.DB) {
	err := db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalf("failed to migrate Product model: %v", err)
	}
}

func (Product) TableName() string {
	return "product"
}

func MapProductToProductDTO(product Product) *dto.ProductDTO {
	return &dto.ProductDTO{
		Name: product.Name,
		Cost: product.Cost,
	}
}

func MapProductsToProductDTOs(products *[]Product) []*dto.ProductDTO {
	productDTOs := make([]*dto.ProductDTO, len(*products))
	for i, product := range *products {
		productDTOs[i] = MapProductToProductDTO(product)
	}
	return productDTOs
}

func MapProductDTOToProduct(dto *dto.ProductDTO) *Product {
	return &Product{
		Name: dto.Name,
		Cost: dto.Cost,
	}
}

func MapProductDTOToProducts(dto *[]dto.ProductDTO) *[]Product {
	products := make([]Product, len(*dto))
	for i, productDTO := range *dto {
		products[i] = *MapProductDTOToProduct(&productDTO)
	}
	return &products
}
