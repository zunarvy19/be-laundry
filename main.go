package main

import (
	"laundry-api/config"
	"laundry-api/controllers"
	"os"

	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//load env
	if os.Getenv("ENV") != "production" {
		godotenv.Load()
	}

	//connect to database
	config.ConnectDB()

	//init gin
	r := gin.Default()

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
