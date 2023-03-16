package models

type Lider struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Nombre     string     `gorm:"not null" json:"nombre"`
	Apellido   string     `gorm:"not null" json:"apellido"`
	Cedula     uint       `gorm:"not null;unique" json:"cedula"`
	Apodo      string     `gorm:"not null" json:"apodo"`
	Telefono   string     `gorm:"not null" json:"telefono"`
	Email      string     `gorm:"not null" json:"email"`
	Parroquia  string     `gorm:"not null" json:"parroquia"`
	Comunidad  string     `gorm:"not null" json:"comunidad"`
	Seguidores []Seguidor `json:"seguidores" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (l *Lider) TableName() string {
	return "lideres"
}

func (l *Lider) SaveLider() (*Lider, error) {

	err := DB.Create(&l).Error

	if err != nil {
		return &Lider{}, err
	}

	return l, nil

}

func (l *Lider) FindLider() (*Lider, error) {

	err := DB.Where(&Lider{ID: l.ID}).Or(&Lider{Cedula: l.Cedula}).Preload("Seguidores").First(&l).Error

	if err != nil {
		return &Lider{}, err
	}
	return l, err

}

func (l *Lider) FindLideres(dto Lider) ([]Lider, error) {
	var lideres []Lider
	err := DB.Where("parroquia LIKE ? AND comunidad LIKE ?", "%"+dto.Parroquia+"%", "%"+dto.Comunidad+"%").Find(&lideres).Error
	if err != nil {
		return []Lider{}, err
	}
	return lideres, nil
}

func (l *Lider) UpdateLider() error {

	err := DB.Model(&Lider{}).Where(&Lider{ID: l.ID}).Updates(&l).Error

	return err

}

func (l *Lider) DeleteLider() (int, error) {

	res := DB.Model(&Lider{}).Where(&Lider{
		ID: l.ID,
	}).Delete(&l)

	return int(res.RowsAffected), res.Error

}
