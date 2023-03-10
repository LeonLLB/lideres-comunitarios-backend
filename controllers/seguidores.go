package controllers

import (
	"lideres-comunitarios-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SeguidorInput struct {
	Nombre    string `binding:"required" json:"nombre"`
	Apellido  string `binding:"required" json:"apellido"`
	Cedula    uint   `binding:"required" json:"cedula"`
	Apodo     string `binding:"required" json:"apodo"`
	Telefono  string `binding:"required" json:"telefono"`
	Email     string `binding:"required" json:"email"`
	Parroquia string `binding:"required" json:"parroquia"`
	Comunidad string `binding:"required" json:"comunidad"`
	LiderID   uint   `binding:"required" json:"liderId"`
}

type SeguidorFiltrado struct {
	Parroquia string `form:"parroquia"`
	Comunidad string `form:"comunidad"`
}

func CreateSeguidor(c *gin.Context) {
	var dto SeguidorInput

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nSeguidor := models.Seguidor{
		Nombre:    dto.Nombre,
		Apellido:  dto.Apellido,
		Cedula:    dto.Cedula,
		Apodo:     dto.Apodo,
		Telefono:  dto.Telefono,
		Email:     dto.Email,
		Parroquia: dto.Parroquia,
		Comunidad: dto.Comunidad,
		LiderID:   dto.LiderID,
	}

	cSeguidor, err := nSeguidor.SaveSeguidor()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": true,
		"data":   cSeguidor,
	})
}

func UpdateSeguidor(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))

	if paramErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Id enviada no es valida"})
		return
	}

	var dto SeguidorInput
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uSeguidor := models.Seguidor{
		ID:        uint(id),
		Nombre:    dto.Nombre,
		Apellido:  dto.Apellido,
		Cedula:    dto.Cedula,
		Apodo:     dto.Apodo,
		Telefono:  dto.Telefono,
		Email:     dto.Email,
		Parroquia: dto.Parroquia,
		Comunidad: dto.Comunidad,
		LiderID:   dto.LiderID,
	}

	if qErr := uSeguidor.UpdateSeguidor(); qErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": qErr.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"data":    uSeguidor,
	})

}

func DeleteSeguidor(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID no valido"})
		return
	}

	dSeguidor := models.Seguidor{ID: uint(id)}
	rows, err := dSeguidor.DeleteSeguidor()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "No existe ese seguidor"})
	}

	c.Status(http.StatusAccepted)

}
