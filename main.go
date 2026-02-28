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
	r.GET("/", controllers.CheckApi)

	//register routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	//laundrypackages routes
	r.GET("/packages", controllers.GetPackages)
	r.POST("/packages", controllers.CreatePackage)
	r.PUT("/packages/:id", controllers.UpdatePackage)
	r.DELETE("/packages/:id", controllers.DeletePackage)

	//contacts routes
	r.GET("/contacts", controllers.GetContacts)
	r.POST("/contacts", controllers.CreateContact)
	r.PUT("/contacts/:id", controllers.UpdateContact)
	r.DELETE("/contacts/:id", controllers.DeleteContact)

	//webcontent routes
	r.GET("/webcontent", controllers.GetWebContent)
	r.POST("/webcontent", controllers.CreateWebContent)
	r.PUT("/webcontent/:id", controllers.UpdateWebContent)
	r.DELETE("/webcontent/:id", controllers.DeleteWebContent)

	//run server
	r.Run(":9010")
}
