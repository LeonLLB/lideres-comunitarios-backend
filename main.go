package main

import (
	"lideres-comunitarios-backend/controllers"
	"lideres-comunitarios-backend/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file, perhaps it doesn't exist, switching to OS variables")
	}

	if os.Getenv("DEV") == "1" {
		models.InitDevDatabase()
	} else {
		//TODO: CONECTARSE A BASE DE DATOS EN PRODUCCION
		models.InitProdDatabase()
	}

	r := gin.Default()

	auth := r.Group("/auth")
	auth.POST("/login", controllers.UserLogin)
	auth.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"john": "doe"})
	})

	r.Run(":" + os.Getenv("PORT"))
}
