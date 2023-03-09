package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Cedula   int    `gorm:"unique; not null" json:"cedula"`
	Password string `gorm:"not null" json:"password"`
	Rol      string `gorm:"not null;default:S" json:"rol"`
}

func (u *Usuario) TableName() string {
	return "usuarios"
}
