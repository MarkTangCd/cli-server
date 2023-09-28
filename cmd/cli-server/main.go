package main

import (
	"cli-server/internal/cli-server/controller"
	"cli-server/internal/cli-server/store"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Header)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Init() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	g := gin.Default()
	g.Use(gin.Recovery())
	g.Use(CORS())
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	templateRouter := g.Group("/template")
	templateRouter.GET("/list", controller.TemplateList)
	templateRouter.POST("/create", controller.CreateTemplate)
	templateRouter.POST("/update", controller.UpdateTemplate)
	templateRouter.GET("/delete/:value", controller.DeleteTemplate)

	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return g
}

func main() {
	client := store.InitDB()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	g := Init()
	// Listen and Server in 0.0.0.0:8080
	g.Run(":8080")
}
