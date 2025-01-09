package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wharf/wharf/conf"
	"github.com/wharf/wharf/internal/routes"
	"github.com/wharf/wharf/pkg/auth"
	"github.com/wharf/wharf/pkg/store"
)

func main() {

	port := "9001"

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}

	router := gin.New()
	router.Use(cors.New(corsConfig))
	router.Use(gin.Logger())

	dockerClient := conf.InitDockerClient()
	api := router.Group("/api/protected")
	{
		api.Use(auth.MiddleWare())
		routes.UserRoutes(api)
		routes.ContainerRoutes(api, dockerClient)
		routes.ImageRoutes(api, dockerClient)
		routes.VolumeRoutes(api, dockerClient)
		routes.NetworkRoutes(api, dockerClient)
	}

	routes.AuthRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		filePath := "./static/" + path
		if !strings.HasPrefix(path, "/api") && !strings.HasPrefix(path, "/docs") {
			if _, err := os.Stat(filePath); err == nil {
				c.File(filePath)
			} else {
				c.File("./static/index.html")
			}

		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		}
	})

	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	router.GET("/docs/api", func(c *gin.Context) {
		c.File("./docs/api_doc.html")
	})

	conf.InitDir()
	go conf.InitPassword()
	go store.InitStore()
	conf.InitCache()
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

}
