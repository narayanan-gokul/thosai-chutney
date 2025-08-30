package consumer

import (
	"context"
	"thosai-chutney/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConsumer(conn *pgxpool.Pool, consumer *Consumer) {
	tx, err := conn.Begin(context.Background())
	utils.CheckError(err)
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(
		context.Background(),
		"INSERT INTO consumer(postcode, first_name, last_name, password) VALUES ($1, $2, $3, $4)",
		consumer.Postcode,
		consumer.FirstName,
		consumer.LastName,
		consumer.Password,
	)
	utils.CheckError(err)

	err = tx.Commit(context.Background())
	utils.CheckError(err)
}
