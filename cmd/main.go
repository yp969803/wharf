package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wharf/wharf/conf"
	"github.com/wharf/wharf/internal/routes"
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
	routes.ContainerRoutes(router)
	routes.ImageRoutes(router)
	routes.VolumeRoutes(router)
	routes.NetworkRoutes(router)
	routes.AuthRoutes(router)
	conf.InitDir()
	go conf.InitPassword()
	go store.InitStore()
	conf.InitCache()
	router.Run(":" + port)
}
