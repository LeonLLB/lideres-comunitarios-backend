package models

import (
	"errors"
	"fmt"
)

type Seguidor struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Nombre    string `gorm:"not null" json:"nombre"`
	Apellido  string `gorm:"not null" json:"apellido"`
	Cedula    uint   `gorm:"not null;unique" json:"cedula"`
	Apodo     string `gorm:"not null" json:"apodo"`
	Telefono  string `gorm:"not null" json:"telefono"`
	Email     string `gorm:"not null" json:"email"`
	Parroquia string `gorm:"not null" json:"parroquia"`
	Comunidad string `gorm:"not null" json:"comunidad"`
	LiderID   uint   `gorm:"not null" json:"liderId"`
}

func (s *Seguidor) TableName() string {
	return "seguidores"
}

func (s *Seguidor) SaveSeguidor() (*Seguidor, error) {
	var lider Lider
	DB.Find(&lider, &Lider{ID: s.LiderID})

	if lider.ID != s.LiderID {
		return &Seguidor{}, errors.New("no such lider exists")
	}

	err := DB.Create(&s).Error

	if err != nil {
		return &Seguidor{}, err
	}

	return s, nil

}

func (s *Seguidor) UpdateSeguidor() error {

	var lider Lider
	DB.Find(&lider, &Lider{ID: s.LiderID})
	fmt.Print(lider.ID)
	fmt.Print(s.LiderID)
	fmt.Print(lider.ID == s.LiderID)
	if lider.ID != s.LiderID {
		return errors.New("no such lider exists")
	}

	err := DB.Model(&Seguidor{}).Where(&Seguidor{ID: s.ID}).Updates(&s).Error

	return err

}

func (s *Seguidor) DeleteSeguidor() (int, error) {

	res := DB.Model(&Seguidor{}).Where(&Seguidor{
		ID: s.ID,
	}).Delete(&s)

	return int(res.RowsAffected), res.Error

}

func (s *Seguidor) BuscarSeguidor() error {
	err := DB.Where(&Seguidor{ID: s.ID}).First(&s).Error

	if err != nil {
		return err
	}
	return nil
}
