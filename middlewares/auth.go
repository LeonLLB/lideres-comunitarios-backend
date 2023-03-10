package middlewares

import (
	"lideres-comunitarios-backend/models"
	"lideres-comunitarios-backend/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateToken(c *gin.Context) {
	token, err := c.Cookie("x-token")

	var resErr gin.H
	if err != nil {
		resErr = gin.H{"error": err.Error()}
	} else if token == "" {
		resErr = gin.H{"error": "no token"}
	}

	_, tokenErr := utilities.TokenValid(c)
	if tokenErr != nil {
		resErr = gin.H{"error": "nonvalid token"}
	}

	if err != nil || token == "" || tokenErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, resErr)
		return
	}
	c.Next()
}

func validateUser(c *gin.Context, wr string) bool {
	//Suponiendo que este middleware se llama despues del ValidateToken

	id, err := utilities.ExtractTokenID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	qUser := models.Usuario{
		ID: id,
	}
	usr, qErr := qUser.FindUsuario()

	if qErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": qErr.Error()})
		return false
	} else if usr.Rol != wr && wr != "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "unauthorized"})
		return false
	}
	return true
}

func ValidateIfAdmin(c *gin.Context) {
	if res := validateUser(c, "A"); res {
		c.Next()
	}
}
func ValidateAnyUser(c *gin.Context) {
	if res := validateUser(c, ""); res {
		c.Next()
	}
}
