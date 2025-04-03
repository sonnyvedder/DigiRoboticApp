package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "os"
    "digiroboticapp/api"  // Using module path
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Printf("Error loading .env file: %v", err)
    }
}

func main() {
    r := gin.Default()
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    apiGroup := r.Group("/api")
    {
        apiGroup.POST("/robot/command", api.HandleRobotCommand)
        apiGroup.GET("/robot/state/:id", api.GetRobotState)
    }

    log.Printf("Starting robot control server on :%s", port)
    r.Run(":" + port)
}