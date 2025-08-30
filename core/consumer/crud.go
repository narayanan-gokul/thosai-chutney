package consumer

import (
	"context"
	"thosai-chutney/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConsumer(conn *pgxpool.Pool, consumer *Consumer) int {
	tx, err := conn.Begin(context.Background())
	utils.CheckError(err)
	defer tx.Rollback(context.Background())

	row := tx.QueryRow(
		context.Background(),
		"INSERT INTO consumer(postcode, first_name, last_name, password) VALUES ($1, $2, $3, $4) RETURNING cons_id",
		consumer.Postcode,
		consumer.FirstName,
		consumer.LastName,
		consumer.Password,
	)

	var cons_id int
	err = row.Scan(&cons_id)
	utils.CheckError(err)

	err = tx.Commit(context.Background())
	utils.CheckError(err)
	return cons_id
}

func FindConsumer(conn *pgxpool.Pool, consId int) *Consumer {
	tx, err := conn.Begin(context.Background())
	utils.CheckError(err)
	defer tx.Rollback(context.Background())

	var consumer Consumer

	row := tx.QueryRow(
		context.Background(),
		"SELECT * FROM consumer WHERE cons_id = $1",
		consId,
	)

	err = row.Scan(
		&consumer.ConsId,
		&consumer.Postcode,
		&consumer.FirstName,
		&consumer.LastName,
		&consumer.Password,
	)
	utils.CheckError(err)

	return &consumer
}
