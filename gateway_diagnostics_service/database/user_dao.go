package database

import (
	"context"
	"fmt"
	"golang_service_template/logger"
	"golang_service_template/models"
	"net/http"
)

type userDb struct{}

type userDbInterface interface {
	GetUsers() ([]models.User, *models.ApiError)
	GetUser(userId int) (*models.User, *models.ApiError)
	CreateUser(user models.CreateUserRequest) (*string, *models.ApiError)
	DeleteUser(userId int) (*string, *models.ApiError)
}

var UserDb userDbInterface

func init() {
	UserDb = &userDb{}
}

func (d *userDb) GetUsers() ([]models.User, *models.ApiError) {

	//rows, err := dbPool.Query(context.Background(), "select id,name,email,age from users")
	rows, err := dbPool.Query(context.Background(), "select id,name,email from users")

	if err != nil {
		logger.Log.Error(err.Error())
		apiError := models.ApiError{
			StatusCode: http.StatusInternalServerError,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1008,
					ErrorMessage: fmt.Sprintf("%s", err),
				},
			},
		}
		return nil, &apiError
	}

	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		//err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age)
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			logger.Log.Error(err.Error())
			apiError := models.ApiError{
				StatusCode: http.StatusInternalServerError,
				ApplicationError: models.ApplicationError{
					Type: "error",
					Message: models.ApplicationErrorMessage{
						ErrorCode:    1009,
						ErrorMessage: fmt.Sprintf("%s", err),
					},
				},
			}
			return nil, &apiError
		}
		users = append(users, user)
	}

	return users, nil

}

func (d *userDb) GetUser(userId int) (*models.User, *models.ApiError) {

	user := models.User{}

	err := dbPool.QueryRow(context.Background(), "select id, name, email from users where id=$1", userId).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		logger.Log.Error(err.Error())
		apiError := models.ApiError{
			StatusCode: http.StatusInternalServerError,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1008,
					ErrorMessage: fmt.Sprintf("%s", err),
				},
			},
		}
		return nil, &apiError
	}

	return &user, nil

}

func (d *userDb) CreateUser(user models.CreateUserRequest) (*string, *models.ApiError) {

	var id int

	fmt.Printf("Name value:%v", user.Name)
	fmt.Printf("Name type:%T", user.Name)
	fmt.Printf("Email value:%v", user.Email)
	fmt.Printf("Email type:%T", user.Email)

	sqlStatement := `insert into users (name,email) values ($1,$2) returning id`
	err := dbPool.QueryRow(context.Background(), sqlStatement, user.Name, user.Email).Scan(&id)

	if err != nil {
		logger.Log.Error(err.Error())
		apiError := models.ApiError{
			StatusCode: http.StatusInternalServerError,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1008,
					ErrorMessage: fmt.Sprintf("%s", err),
				},
			},
		}
		return nil, &apiError
	}

	successMsg := fmt.Sprintf("User with user id %d addded!", id)

	return &successMsg, nil

}

func (d *userDb) DeleteUser(userId int) (*string, *models.ApiError) {

	res, err := dbPool.Exec(context.Background(), "delete from users where id=$1", userId)

	if err != nil {
		logger.Log.Error(err.Error())
		apiError := models.ApiError{
			StatusCode: http.StatusInternalServerError,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1008,
					ErrorMessage: fmt.Sprintf("%s", err),
				},
			},
		}
		return nil, &apiError
	}

	if res.RowsAffected() != 1 {
		apiError := models.ApiError{
			StatusCode: http.StatusInternalServerError,
			ApplicationError: models.ApplicationError{
				Type: "error",
				Message: models.ApplicationErrorMessage{
					ErrorCode:    1008,
					ErrorMessage: fmt.Sprintf("Could not delete user with user id %d!", userId),
				},
			},
		}
		return nil, &apiError
	}

	successMsg := fmt.Sprintf("User with user id %d deleted!", userId)

	return &successMsg, nil

}
