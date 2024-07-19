package controllers

import (
	
	"fmt"
	"golang_service_template/services"
	"net/http"
    "github.com/gin-gonic/gin"
	
)


func InsertApplication(c *gin.Context) {
    var input struct {
        Type              string `json:"type"`
        ApplicationNumber int    `json:"application_number"`
        ApplicationName   string `json:"application_name"`
    }

    // Decode the JSON input
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Insert the application data into the database
    id, err := services.InsertApplicationData(pool, input.Type, input.ApplicationNumber, input.ApplicationName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unable to insert data: %v", err)})
        return
    }

    // Return the created ID
    c.JSON(http.StatusCreated, gin.H{"id": id})
}
