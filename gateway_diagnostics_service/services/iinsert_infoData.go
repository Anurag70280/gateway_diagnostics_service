package services

import (
    "context"
    "github.com/jackc/pgx/v4/pgxpool"
)

func InsertInfoData(pool *pgxpool.Pool, infoType string, applicationNumber int, messageType string, message string, details string) (int, error) {
    query := `
        INSERT INTO info (type, application_number, message_type, message, details)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`

    var id int
    err := pool.QueryRow(context.Background(), query, infoType, applicationNumber, messageType, message, details).Scan(&id)
    if err != nil {
        return 0, err
    }

    return id, nil
}