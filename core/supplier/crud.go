package supplier

import (
	"context"
	"fmt"
	"thosai-chutney/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateSupplier(conn *pgxpool.Pool, supplier *Supplier) {
	tx, err := conn.Begin(context.Background())
	utils.CheckError(err)
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(
		context.Background(),
		"INSERT INTO supplier(postcode, name, password) VALUES ($1, $2, $3)",
		supplier.Postcode,
		supplier.Name,
		supplier.Password,
	)
	utils.CheckError(err)

	err = tx.Commit(context.Background())
	utils.CheckError(err)
}
