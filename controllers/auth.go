package controllers

import (
	"lideres-comunitarios-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

func UserLogin(c *gin.Context) {

	var dto LoginInput

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"John": "Doe"})
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

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "data": cUser})

}
