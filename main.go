package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"laundry-api/config"
	"laundry-api/controllers"
)

func main() {
	//load env
	required := []string{"DB_HOST", "DB_USER", "DB_PASSWORD"}

	for _, v := range required {
		if os.Getenv(v) == "" {
			log.Fatalf("%s is required", v)
		}
	}

	//connect to database
	config.ConnectDB()

	//init gin
	r := gin.Default()

	// enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Check API
	r.GET("/api/", controllers.CheckApi)

	//register routes
	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)

	//laundrypackages routes
	r.GET("/api/packages", controllers.GetPackages)
	r.POST("/api/packages", controllers.CreatePackage)
	r.PUT("/api/packages/:id", controllers.UpdatePackage)
	r.DELETE("/api/packages/:id", controllers.DeletePackage)

	//contacts routes
	r.GET("/api/contacts", controllers.GetContacts)
	r.POST("/api/contacts", controllers.CreateContact)
	r.PUT("/api/contacts/:id", controllers.UpdateContact)
	r.DELETE("/api/contacts/:id", controllers.DeleteContact)

	//webcontent routes
	r.GET("/api/webcontent", controllers.GetWebContent)
	r.POST("/api/webcontent", controllers.CreateWebContent)
	r.PUT("/api/webcontent/:id", controllers.UpdateWebContent)
	r.DELETE("/api/webcontent/:id", controllers.DeleteWebContent)

	//run server
	r.Run(":9010")
}
