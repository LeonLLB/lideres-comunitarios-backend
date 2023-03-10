package utilities

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GenUserJWT(userId uint) (string, error) {

	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatal("Cannot load local ENV vars")
	}

	token_lifespan, err := strconv.Atoi(os.Getenv("JWT_TOKEN_MIN_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_API_SECRET")))
}

func ExtractToken(c *gin.Context) string {
	token, err := c.Cookie("x-token")
	if err == nil {
		return token
	}

	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {

	token, err := TokenValid(c)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

func TokenValid(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metodo de cifrado inesperado: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func RevalidateToken(c *gin.Context) (string, error) {
	id, err := ExtractTokenID(c)
	if err != nil {
		return "", err
	}
	tkn, genErr := GenUserJWT(id)
	if err != nil {
		return "", err
	}
	return tkn, genErr
}
