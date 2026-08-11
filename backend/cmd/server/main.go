package main

import (
	"Stork/internal/handler"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Stork backend...")

	router := gin.Default()
	router.POST("/send", handler.SendEmailHandler)

	router.Run(":5000")
}
