package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Cedula   int    `gorm:"unique; not null" json:"cedula"`
	Password string `gorm:"not null" json:"password"`
}

func (u *Usuario) TableName() string {
	return "usuarios"
}
