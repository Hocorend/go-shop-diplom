package models

import (
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID      uint32  `gorm:"primaryKey;autoIncrement" json:"id"`
	Login   *string `json:"login"`
	Deposit uint32  `json:"deposit"`
}

func (User) TableName() string {
	return "user_shop"
}

func MigrateUser(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("failed to migrate User model: %v", err)
	}
}
