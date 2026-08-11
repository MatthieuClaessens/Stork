package handler

import (
	"Stork/internal/email"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendEmailHandler(c *gin.Context) {
	var mail email.Email
	if err := c.ShouldBindJSON(&mail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := email.SendRequest(mail, "localhost", "1025")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully sent anonymous mail"})
}
