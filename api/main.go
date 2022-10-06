package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/al26p/lets_go_sports/api/controllers"
	"github.com/al26p/lets_go_sports/api/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world")

	// Initialize default router
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000", "http://192.168.0.12:3000"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Connect to DB
	models.ConnectDB()

	// Define routes
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hello world !"})
	})
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.PATCH("/users/:id", controllers.PatchUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run()
}
