package main

import "github.com/gin-gonic/gin"

func main() {
    // Creates a gin router with default middleware:
    // logger and recovery (crash-free) middleware
    router := gin.Default()

    router.GET("/someGet", func(c *gin.Context) {
        c.String(200, "pong")
    })
    // Listen and server on 0.0.0.0:8080
    router.Run(":8080")
}
