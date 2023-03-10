package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Cedula   int    `gorm:"unique; not null" json:"cedula"`
	Password string `gorm:"not null" json:"password"`
	Rol      string `gorm:"not null;default:S" json:"rol"`
}

func (u *Usuario) TableName() string {
	return "usuarios"
}

func (u *Usuario) SaveUsuario() (*Usuario, error) {

	err := DB.Create(&u).Error

	if err != nil {
		return &Usuario{}, err
	}

	return u, nil

}

func (u *Usuario) BeforeSave(_ *gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}

func (u *Usuario) FindUsuario() (*Usuario, error) {

	err := DB.Where(&Usuario{ID: u.ID}).Or(&Usuario{Cedula: u.Cedula}).First(&u).Error

	if err != nil {
		return &Usuario{}, err
	}
	return u, err

}

func (u *Usuario) CheckPassword(pass string) error {

	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
}
