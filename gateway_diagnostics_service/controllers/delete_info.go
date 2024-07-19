package controllers

import (
	"fmt"
	"golang_service_template/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteInfo(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    rowsAffected, err := services.DeleteInfoData(pool, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unable to delete data: %v", err)})
        return
    }

    if rowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No info found with the provided ID"})
    } else {
        c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Deleted info with ID: %d", id)})
    }
}
