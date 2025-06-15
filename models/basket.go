package models

import (
	"gorm.io/gorm"
	"log"
)

type Basket struct {
	ID        uint32   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint32   `json:"user_id" gorm:"notNull"`
	User      *User    `json:"user" gorm:"foreignKey:UserID"`
	ProductID uint32   `json:"product_id" gorm:"notNull"`
	Product   *Product `json:"product" gorm:"foreignKey:ProductID"`
	Count     uint32   `json:"count" gorm:"notNull"`
}

func MigrateBasket(db *gorm.DB) {
	err := db.AutoMigrate(&Basket{})
	if err != nil {
		log.Fatalf("failed to migrate Basket model: %v", err)
	}
}

func (Basket) TableName() string {
	return "basket"
}
