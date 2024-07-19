package services

import (
	"golang_service_template/database"
	"golang_service_template/models"
)

func GetUsers() ([]models.User, *models.ApiError) {
	return database.UserDb.GetUsers()
}

func GetUser(userId int) (*models.User, *models.ApiError) {
	return database.UserDb.GetUser(userId)
}

func CreateUser(user models.CreateUserRequest) (*string, *models.ApiError) {
	return database.UserDb.CreateUser(user)
}

func DeleteUser(userId int) (*string, *models.ApiError) {
	return database.UserDb.DeleteUser(userId)
}
