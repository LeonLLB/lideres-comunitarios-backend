package controllers

import (
	"lideres-comunitarios-backend/models"
	"log"
	"net/http"
	"os"

	"lideres-comunitarios-backend/utilities"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type UserInputCore struct {
	Cedula   int    `json:"cedula" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	UserInputCore
}

type RegisterInput struct {
	UserInputCore
	Rol string `json:"rol"`
}

func SetCookie(c *gin.Context, name string, value string) {

	if os.Getenv("DEV") == "" || os.Getenv("DEV") == "1" {
		env_err := godotenv.Load(".env")

		if env_err != nil {
			log.Fatal("Cannot load local .env")
		}
	}

	domain := os.Getenv("FRONT_DOMAIN")

	c.SetCookie(name, value, 60*60, "/", domain, domain != "localhost", true)
}

func UserLogin(c *gin.Context) {

	var dto LoginInput

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dUser := models.Usuario{
		Cedula: dto.Cedula,
	}

	qUser, err := dUser.FindUsuario()

	loginError := gin.H{"error": "Usuario o clave invalida"}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, loginError)
		return
	}

	pass_err := qUser.CheckPassword(dto.Password)

	if pass_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, loginError)
		return
	}

	token, err := utilities.GenUserJWT(qUser.ID)

	if err != nil {
		log.Fatal(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	SetCookie(c, "x-token", token)

	c.JSON(http.StatusAccepted, gin.H{"token": token})
}

func RegisterUser(c *gin.Context) {

	var dto RegisterInput

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nUser := models.Usuario{
		Cedula:   dto.Cedula,
		Password: dto.Password,
		Rol:      dto.Rol,
	}

	cUser, err := nUser.SaveUsuario()
	cUser.Password = ""

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "data": cUser})

}

func Logout(c *gin.Context) {
	SetCookie(c, "x-token", "")
	c.JSON(http.StatusAccepted, gin.H{"success": true})
}

func GetTokenStatus(c *gin.Context) {
	id, err := utilities.ExtractTokenID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"auth": true,
	})
}
