package main

import (
	"log"
	"shazam_music_query/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// // start the server
	router := gin.Default()

	// set up the cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization "}
	router.Use(cors.New(config))

	mainController := controllers.MainController{}

	// router group
	routerGroup := router.Group("/api/v1")

	// set up the routes
	routerGroup.POST("/music-search", mainController.MusicSearchView)

	// start the server
	log.Fatal(router.Run(":8080"))

}
