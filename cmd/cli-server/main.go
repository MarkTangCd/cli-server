package main

import (
	"cli-server/internal/cli-server/controller"
	"cli-server/internal/cli-server/store"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	templateRouter := r.Group("/template")
	templateRouter.GET("/", controller.TemplateList)
	templateRouter.POST("/create", controller.CreateTemplate)
	templateRouter.POST("/update", controller.UpdateTemplate)
	templateRouter.GET("/delete/:value", controller.DeleteTemplate)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	client := store.InitDB()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
