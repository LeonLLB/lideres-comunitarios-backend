package controllers

import (
	"lideres-comunitarios-backend/models"
	"net/http"
	"strconv"

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

type LiderFiltrado struct {
	Parroquia string `form:"parroquia"`
	Comunidad string `form:"comunidad"`
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

func GetLideres(c *gin.Context) {

	var dto LiderFiltrado

	if err := c.ShouldBindQuery(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lider := models.Lider{}
	lideres, err := lider.FindLideres(models.Lider{Parroquia: dto.Parroquia, Comunidad: dto.Comunidad})

	var errRes gin.H
	var errCode int

	if err != nil {
		errRes = gin.H{"error": err.Error()}
		errCode = http.StatusInternalServerError
	} else if len(lideres) == 0 {
		errRes = gin.H{"error": "no hay lideres"}
		errCode = http.StatusNotFound
	}

	if err != nil || len(lideres) == 0 {
		c.AbortWithStatusJSON(errCode, errRes)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": lideres,
	})
}

func UpdateLider(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))

	if paramErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Id enviada no es valida"})
		return
	}

	var dto LiderInput
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uLider := models.Lider{
		ID:        uint(id),
		Nombre:    dto.Nombre,
		Apellido:  dto.Apellido,
		Cedula:    dto.Cedula,
		Apodo:     dto.Apodo,
		Telefono:  dto.Telefono,
		Email:     dto.Email,
		Parroquia: dto.Parroquia,
		Comunidad: dto.Comunidad,
	}

	if qErr := uLider.UpdateLider(); qErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": qErr.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"data":    uLider,
	})

}

func DeleteLider(c *gin.Context) {
	id, paramErr := strconv.Atoi(c.Param("id"))
	if paramErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID no valido"})
		return
	}

	dLider := models.Lider{ID: uint(id)}
	rows, err := dLider.DeleteLider()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rows == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "No existe ese lider"})
	}

	c.Status(http.StatusAccepted)

}
