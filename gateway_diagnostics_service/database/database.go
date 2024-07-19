package database

import (
	"context"
	"fmt"
	"golang_service_template/logger"
	"golang_service_template/models"
	"golang_service_template/utils"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func InitializeDatabasePool() error {

	var err error

	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	logger.Log.Debug(databaseUrl)

	dbconfig, _ := pgxpool.ParseConfig(databaseUrl)
	dbPool, err = pgxpool.NewWithConfig(context.Background(), dbconfig)

	if err != nil {

		logger.Log.Error(err.Error())
		utils.ProcessError(models.ProcessErrorMessage{Priority: 1, Error: err}, databaseUrl)
		return err

	}

	return nil

}

func PingDatabasePool() error {

	return dbPool.Ping(context.Background())

}

func CloseDatabasePool() {

	dbPool.Close()

}
