package test

import (
	"encoding/json"
	"golang_service_template/app"
	"golang_service_template/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	funcGetUsersDbResponse   func() ([]models.User, *models.ApiError)
	funcGetUserDbResponse    func() (*models.User, *models.ApiError)
	funcCreateUserDbResponse func() (*string, *models.ApiError)
	funcDeleteUserDbResponse func() (*string, *models.ApiError)
)

type userDbMock struct{}

func init() {

	//database.UserDb = &userDbMock{}

}

func (d *userDbMock) GetUsers() ([]models.User, *models.ApiError) {

	return funcGetUsersDbResponse()

}

func (d *userDbMock) GetUser(userId int) (*models.User, *models.ApiError) {

	return funcGetUserDbResponse()

}

func (d *userDbMock) CreateUser(user models.CreateUserRequest) (*string, *models.ApiError) {

	return funcCreateUserDbResponse()

}

func (d *userDbMock) DeleteUser(userId int) (*string, *models.ApiError) {

	return funcDeleteUserDbResponse()

}

func TestGetUserNoErrors(t *testing.T) {

	assert := assert.New(t)

	app.SetupCognitoRoutes()

	//database.UserDb = &userDbMock{}

	funcGetUserDbResponse = func() (*models.User, *models.ApiError) {

		dbResponse := models.User{Id: 45, Name: "malcolm", Email: "malcolm@spintly.com"}

		return &dbResponse, nil
	}

	resWriter := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/user/1", nil)
	req.Header.Set("Content-Type", "application/json")
	app.Router.ServeHTTP(resWriter, req)

	getUserResponse := models.GetUserResponse{}
	json.Unmarshal(resWriter.Body.Bytes(), &getUserResponse)

	assert.Equal(200, resWriter.Code)
	assert.Equal(45, getUserResponse.Message.Id)
	assert.Equal("pallavi", getUserResponse.Message.Name)
	assert.Equal("malcolm@spintly.com", getUserResponse.Message.Email)

}
