package main

import (
	"fmt"
	"lideres-comunitarios-backend/controllers"
	"lideres-comunitarios-backend/middlewares"
	"lideres-comunitarios-backend/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file, perhaps it doesn't exist, switching to OS variables")
	}

	pass := middlewares.GenSecretRoutePassword()
	fmt.Printf("SECRET GENERATED PASSWORD %s\n", pass)

	if os.Getenv("DEV") == "1" {
		models.InitDevDatabase()
	} else {
		//TODO: CONECTARSE A BASE DE DATOS EN PRODUCCION
		models.InitProdDatabase()
	}

	r := gin.Default()

	auth := r.Group("/auth")
	auth.POST("/login", controllers.UserLogin)
	auth.POST("/register", middlewares.ValidateSecretRoutePassword, controllers.RegisterUser)
	auth.POST("/logout", controllers.Logout)
	// auth.POST("/protected", middlewares.ValidateToken, middlewares.ValidateIfAdmin, middlewares.RevalidateUsrToken, func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"hello": "We wanted to talk about your cars extended warranty"})
	// })

	lideres := r.Group("/lideres")
	lideres.POST("/", middlewares.ValidateIfAdmin, middlewares.RevalidateUsrToken, controllers.CreateLider)

	r.Run(":" + os.Getenv("PORT"))
}
