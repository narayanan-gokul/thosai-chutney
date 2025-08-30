package consumer

import (
	"context"
	"fmt"
	"thosai-chutney/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConsumer(conn *pgxpool.Pool, user Consumer) {
	fmt.Println("Creating transaction")
	tx, err := conn.Begin(context.Background())
	utils.CheckError(err)
	defer tx.Rollback(context.Background())
	fmt.Println("Creating transaction")

	fmt.Println("Creating transaction")
	row, err := tx.Query(
		context.Background(),
		"INSERT INTO consumer(postcode, first_name, last_name, password)",
	)
	defer row.Close()
	tx.Commit(context.Background())
}
