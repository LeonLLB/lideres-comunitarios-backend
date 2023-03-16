package main

import (
	"fmt"
	"lideres-comunitarios-backend/controllers"
	"lideres-comunitarios-backend/middlewares"
	"lideres-comunitarios-backend/models"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file, perhaps it doesn't exist, switching to OS variables")
	}

	pass := middlewares.GenSecretRoutePassword()
	fmt.Printf("SECRET GENERATED PASSWORD %s\n", pass)

	models.InitDatabase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://lideres-comunitarios.vercel.app"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := r.Group("/auth")
	auth.POST("/login", controllers.UserLogin)
	auth.POST("/register", middlewares.ValidateSecretRoutePassword, controllers.RegisterUser)
	auth.POST("/logout", controllers.Logout)
	auth.POST("/revalidate", middlewares.ValidateAnyUser, controllers.GetTokenStatus)
	// auth.POST("/protected", middlewares.ValidateToken, middlewares.ValidateIfAdmin, middlewares.RevalidateUsrToken, func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"hello": "We wanted to talk about your cars extended warranty"})
	// })

	lideres := r.Group("/lideres")
	lideres.POST("/", middlewares.ValidateIfAdmin, middlewares.RevalidateUsrToken, controllers.CreateLider)
	lideres.GET("/", middlewares.ValidateAnyUser, middlewares.RevalidateUsrToken, controllers.GetLideres)
	lideres.GET("/:id", middlewares.ValidateAnyUser, middlewares.RevalidateUsrToken, controllers.GetLider)
	lideres.PUT("/:id", middlewares.ValidateIfAdmin, middlewares.RevalidateUsrToken, controllers.UpdateLider)
	lideres.DELETE("/:id", middlewares.ValidateIfAdmin, middlewares.RevalidateUsrToken, controllers.DeleteLider)

	seguidores := r.Group("/seguidores")
	seguidores.GET("/:id", middlewares.ValidateIfAdmin, middlewares.RevalidateUsrToken, controllers.GetSeguidor)
	seguidores.POST("/", middlewares.ValidateIfAdmin, middlewares.RevalidateUsrToken, controllers.CreateSeguidor)
	seguidores.PUT("/:id", middlewares.ValidateIfAdmin, middlewares.RevalidateUsrToken, controllers.UpdateSeguidor)
	seguidores.DELETE("/:id", middlewares.ValidateIfAdmin, middlewares.RevalidateUsrToken, controllers.DeleteSeguidor)

	r.Run(":" + os.Getenv("PORT"))
}
