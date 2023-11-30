package model

import "time"

type User struct {
	ID        string `gorm:"primarykey"`
	Name      string
	Age       string
	Email     string `gorm:"unique"`
	Phone     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:false"`
}
