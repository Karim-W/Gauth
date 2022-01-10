package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	FirstName string
	LastName  string
}
