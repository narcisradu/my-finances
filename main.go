package main

import (
	"flag"
	"log"
	"net/http"
)
import "github.com/gin-gonic/gin"

func main() {
	debugMode := flag.Bool("debug", false, "use debug mode")
	flag.Parse()

	if !*debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	heartbeat := router.Group("/heartbeat")
	{
		heartbeat.GET("/info", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{})
		})
	}

	err := router.Run()
	if err != nil {
		log.Println("Error on router run: ", err)
	}
}

func init() {
	log.Println("Starting My Finances Service...")
}
