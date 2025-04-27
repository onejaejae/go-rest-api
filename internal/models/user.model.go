package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName *string `gorm:"type:varchar(50)"`
	LastName  *string `gorm:"type:varchar(50)"`
	Email     string  `gorm:"type:varchar(100);not null"`
	Gender    *string `gorm:"type:varchar(10)"`
	Password  string  `gorm:"type:varchar(255);not null"`
}
