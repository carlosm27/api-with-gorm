package main

import (
	//"fmt"
	"log"
	//"net/http"

	"github.com/carlosm27/go_projects/apiwithgorm/controllers"
	"github.com/carlosm27/go_projects/apiwithgorm/model"
	"github.com/gin-gonic/gin"

	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	model.Database()

	router := gin.Default()

	router.GET("/groceries", controllers.GetGroceries)
	router.GET("/grocery/:name", controllers.GetGrocery)
	router.POST("/grocery", controllers.PostGrocery)
	router.PUT("/grocery/:id", controllers.UpdateGrocery)
	router.DELETE("/grocery/:id", controllers.DeleteGrocery)

	log.Fatal(router.Run(":10000"))
}
