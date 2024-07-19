package services

import (
    "context"
    "github.com/jackc/pgx/v4/pgxpool"
)

type Message struct {
    ApplicationName string
    Message         string
}

func GetMessagesData(pool *pgxpool.Pool, applicationNumber int) ([]Message, error) {
    query := `
        SELECT a.application_name, i.message
        FROM applications a
        INNER JOIN info i ON a.application_number = i.application_number
        WHERE a.application_number = $1;
    `

    rows, err := pool.Query(context.Background(), query, applicationNumber)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var messages []Message
    for rows.Next() {
        var message Message
        err := rows.Scan(&message.ApplicationName, &message.Message)
        if err != nil {
            return nil, err
        }
        messages = append(messages, message)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return messages, nil
}
