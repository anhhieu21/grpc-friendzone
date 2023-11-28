package model

import "time"

type User struct {
	ID        string `gorm:"primarykey"`
	Name      string
	Age       string
	Email     string
	Phone     string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:false"`
}
