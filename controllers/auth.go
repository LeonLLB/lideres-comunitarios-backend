package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Cedula   int    `json:"cedula" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserLogin(c *gin.Context) {

	var dto LoginInput

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"John": "Doe"})
}
