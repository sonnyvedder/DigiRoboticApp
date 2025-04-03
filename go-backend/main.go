package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "os"
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
        port = "8080" // default port
    }

    // Health check endpoint
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "ok",
        })
    })

    // Robot API routes
    api := r.Group("/api")
    {
        api.POST("/robot/command", handleRobotCommand)
        api.GET("/robot/state/:id", getRobotState)
    }

    log.Printf("Starting robot control server on :%s", port)
    r.Run(":" + port)
}

func handleRobotCommand(c *gin.Context) {
    // Initial command handler
    c.JSON(200, gin.H{
        "status": "received",
        "message": "Command processing initialized",
    })
}

func getRobotState(c *gin.Context) {
    robotID := c.Param("id")
    c.JSON(200, gin.H{
        "robot_id": robotID,
        "status": "online",
        "position": map[string]float64{
            "x": 0,
            "y": 0,
            "z": 0,
        },
    })
}