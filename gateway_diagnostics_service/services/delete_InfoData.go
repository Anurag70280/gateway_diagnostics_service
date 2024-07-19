package services

import (
    "context"
    "github.com/jackc/pgx/v4/pgxpool"
)

func DeleteInfoData(pool *pgxpool.Pool, id int) (int64, error) {
    query := `DELETE FROM info WHERE id = $1`

    res, err := pool.Exec(context.Background(), query, id)
    if err != nil {
        return 0, err
    }

    return res.RowsAffected(), nil
}