package models

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDevDatabase() {

	var err error

	DB, err = gorm.Open(sqlite.Open("lideres.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		panic("Can't open local database\n")
	} else {
		fmt.Print("Local database connection has been succesful\n")
	}

	DB.AutoMigrate(&Usuario{}, &Lider{}, &Seguidor{})
}

func InitProdDatabase() {
	//TODO: INICIAR BASE DE DATOS DE PRODUCCION
	fmt.Print("Opening production database\n")
	InitDevDatabase()
}

func GetDbInstance() *gorm.DB {
	return DB
}
