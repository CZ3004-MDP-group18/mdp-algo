package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"mdp_algo/simulator"
	"net/http"
)

func main() {
	fmt.Println("Hello MDP")

	router := gin.Default()
	// Enable all origins
	router.Use(cors.Default())

	// Serve static files from the frontend/dist directory.
	router.StaticFS("/", http.Dir("./frontend/dist"))
	router.POST("/health_check", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{}) })
	router.POST("/set_arena", simulator.HandleSetArena)
	router.POST("/hamilton_path", simulator.HandleGetHamiltonPath)

	fmt.Println("Server listening on port 3000")
	log.Fatal(router.Run(":3000"))
}
