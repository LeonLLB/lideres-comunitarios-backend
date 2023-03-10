package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Cedula   int    `gorm:"unique; not null" json:"cedula"`
	Password string `gorm:"not null" json:"password"`
	Rol      string `gorm:"not null;default:S" json:"rol"`
}

func (u *Usuario) TableName() string {
	return "usuarios"
}

func (u *Usuario) SaveUsuario() (*Usuario, error) {

	err := DB.Create(&u).Error
	fmt.Print(err.Error())
	if err != nil {
		return &Usuario{}, err
	}

	return u, nil

}

func (u *Usuario) BeforeSave() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}
