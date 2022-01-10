package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
}
