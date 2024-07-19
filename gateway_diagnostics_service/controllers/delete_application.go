package controllers

import (
	"fmt"
	"golang_service_template/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	
)


func DeleteApplication(c *gin.Context) {
    // Extract the ID from the URL parameters
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    // Assume pool is a global variable or passed in another way
    rowsAffected, err := services.DeleteApplicationData(pool, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unable to delete data: %v", err)})
        return
    }

    // Check if the application was found and deleted
    if rowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No application found with the provided ID."})
    } else {
        c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Deleted application with ID: %d", id)})
    }
}
