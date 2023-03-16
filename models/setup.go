package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {

	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatal("Cannot load local ENV vars")
	}
	fmt.Print("Opening production database\n")

	DBHost := os.Getenv("DBHOST")
	DBUser := os.Getenv("DBUSER")
	DBPassword := os.Getenv("DBPASSWORD")
	DBName := os.Getenv("DBNAME")
	DBPort := os.Getenv("DBPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Caracas", DBHost, DBUser, DBPassword, DBName, DBPort)
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		panic("Can't open postgres database\n")
	} else {
		fmt.Print("database connection has been succesful\n")
	}
	DB.AutoMigrate(&Usuario{}, &Lider{}, &Seguidor{})
}

func GetDbInstance() *gorm.DB {
	return DB
}
