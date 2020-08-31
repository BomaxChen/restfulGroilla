package model

import (
	"time"

	"gorm.io/gorm"
)

// User Bean
type User struct {
	//gorm.Model
	ID         int       `gorm:"unique" json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email"`
	DateCreate time.Time `json:"dateCreate"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
