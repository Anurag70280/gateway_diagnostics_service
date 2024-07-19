package services

import (
    "context"
    "github.com/jackc/pgx/v4/pgxpool"
)

func InsertApplicationData(pool *pgxpool.Pool, appType string, applicationNumber int, applicationName string) (int, error) {
    query := `
        INSERT INTO applications (type, application_number, application_name)
        VALUES ($1, $2, $3)
        RETURNING id`

    var id int
    err := pool.QueryRow(context.Background(), query, appType, applicationNumber, applicationName).Scan(&id)
    if err != nil {
        return 0, err
    }

    return id, nil
}