package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quminzhi/go-examples/gin-example/controllers"
	"github.com/quminzhi/go-examples/gin-example/middlewares"
	"log"
)

func main() {
	server := gin.Default()
	server.Use(middlewares.AuthMiddleware())

	controller := controllers.NewVideoController()
	group := server.Group("/videos")
	group.Use(middlewares.Logger())
	group.GET("/", controller.GetAll)
	group.PUT("/:id", controller.Update)
	group.POST("/", controller.Create)
	group.DELETE("/:id", controller.Delete)

	log.Fatalln(server.Run("localhost:8080"))
}
