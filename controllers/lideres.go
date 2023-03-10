package controllers

import (
	"lideres-comunitarios-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LiderInput struct {
	Nombre    string `binding:"required" json:"nombre"`
	Apellido  string `binding:"required" json:"apellido"`
	Cedula    uint   `binding:"required" json:"cedula"`
	Apodo     string `binding:"required" json:"apodo"`
	Telefono  string `binding:"required" json:"telefono"`
	Email     string `binding:"required" json:"email"`
	Parroquia string `binding:"required" json:"parroquia"`
	Comunidad string `binding:"required" json:"comunidad"`
}

func CreateLider(c *gin.Context) {
	var dto LiderInput

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nLider := models.Lider{
		Nombre:    dto.Nombre,
		Apellido:  dto.Apellido,
		Cedula:    dto.Cedula,
		Apodo:     dto.Apodo,
		Telefono:  dto.Telefono,
		Email:     dto.Email,
		Parroquia: dto.Parroquia,
		Comunidad: dto.Comunidad,
	}

	cLider, err := nLider.SaveLider()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": true,
		"data":   cLider,
	})
}
