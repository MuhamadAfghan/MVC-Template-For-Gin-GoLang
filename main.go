package main

import (
	"github.com/gin-gonic/gin"
	"go-health-tracker/controllers"
	"go-health-tracker/models"
	"log"
)

func main() {
	dsn := "root:@tcp(localhost:3306)/golang_digiup_ujikom?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta"
	// init db
	models.InitDB(dsn)

	// load routes
	r := gin.Default()

	// load views
	r.LoadHTMLGlob("views/*")

	// ruting
	r.GET("/", controllers.ShowForm)
	r.POST("/submit", controllers.SubmitForm)

	// run server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
