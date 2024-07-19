package controllers

import (
   
    "fmt"
    "net/http"
    "golang_service_template/services"
    "github.com/gin-gonic/gin"
)
func InsertInfo(c *gin.Context) {
    var input struct {
        Type              string `json:"type"`
        ApplicationNumber int    `json:"application_number"`
        MessageType       string `json:"message_type"`
        Message           string `json:"message"`
        Details           string `json:"details"`
    }

    // Decode the JSON input
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Insert the information data into the database
    id, err := services.InsertInfoData(pool, input.Type, input.ApplicationNumber, input.MessageType, input.Message, input.Details)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unable to insert data: %v", err)})
        return
    }

    // Return the created ID
    c.JSON(http.StatusCreated, gin.H{"id": id})
}
