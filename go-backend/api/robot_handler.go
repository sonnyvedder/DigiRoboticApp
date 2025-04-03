package api

import (
    "github.com/gin-gonic/gin"
    "digiroboticapp/models"  // Using module path
)

func HandleRobotCommand(c *gin.Context) {
    var cmd models.RobotCommand
    if err := c.BindJSON(&cmd); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{
        "status": "received",
        "message": "Command processing initialized",
        "command": cmd.Command,
    })
}

func GetRobotState(c *gin.Context) {
    robotID := c.Param("id")
    robot := models.Robot{
        ID:     robotID,
        Status: "online",
        Position: models.Position{
            X: 0,
            Y: 0,
            Z: 0,
        },
    }
    c.JSON(200, robot)
}