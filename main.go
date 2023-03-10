package main

import (
	"fmt"
	"lideres-comunitarios-backend/controllers"
	"lideres-comunitarios-backend/models"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var secretPassword string

func genSecretRoutePassword() string {
	chars := "abcdefghijklmnopqrstuvwxyz"
	var rPass string
	for i := 0; i < 11; i++ {
		char := strings.Split(chars, "")[rand.Intn(len(chars))]
		rPass = rPass + char
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(rPass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print("Couldn't generate secret password for specific functions\n")
		return ""
	}
	secretPassword = string(pass)
	return rPass
}

type secretRouteInput struct {
	Key string `json:"key"`
}

func validateSecretRoutePassword(c *gin.Context) {

	var input secretRouteInput
	response := gin.H{"Unauthorized": "This route is meant to be used with the proper key"}

	if err := c.ShouldBindHeader(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, response)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(secretPassword), []byte(input.Key))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, response)
		return
	}

	c.Next()
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file, perhaps it doesn't exist, switching to OS variables")
	}

	pass := genSecretRoutePassword()
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
	auth.POST("/register", validateSecretRoutePassword, controllers.RegisterUser)
	auth.POST("/protected", func(c *gin.Context) {
		token, err := c.Cookie("x-token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	r.Run(":" + os.Getenv("PORT"))
}
