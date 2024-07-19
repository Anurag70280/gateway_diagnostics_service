package controllers

import (
	"fmt"
	"golang_service_template/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMessages(c *gin.Context) {
    // Retrieve application_number from query parameters
    applicationNumberStr := c.Query("application_number")
    applicationNumber, err := strconv.Atoi(applicationNumberStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application number"})
        return
    }

    messages, err := services.GetMessagesData(pool, applicationNumber)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unable to retrieve data: %v", err)})
        return
    }

    c.JSON(http.StatusOK, messages)
}

