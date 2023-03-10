package middlewares

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var secretPassword string

func GenSecretRoutePassword() string {
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

func ValidateSecretRoutePassword(c *gin.Context) {

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
