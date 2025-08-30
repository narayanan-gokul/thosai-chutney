package supplier

import (
	"context"
	"thosai-chutney/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateSupplier(conn *pgxpool.Pool, supplier *Supplier) int {
	tx, err := conn.Begin(context.Background())
	utils.CheckError(err)
	defer tx.Rollback(context.Background())

	row := tx.QueryRow(
		context.Background(),
		"INSERT INTO supplier(postcode, name, password) VALUES ($1, $2, $3) RETURNING supp_id",
		supplier.Postcode,
		supplier.Name,
		supplier.Password,
	)

	var supp_id int
	err = row.Scan(&supp_id)
	utils.CheckError(err)

	err = tx.Commit(context.Background())
	utils.CheckError(err)
	return supp_id
}

func FindSupplier(conn *pgxpool.Pool, suppId int) *Supplier {
	tx, err := conn.Begin(context.Background())
	utils.CheckError(err)
	defer tx.Rollback(context.Background())

	var supplier Supplier

	row := tx.QueryRow(
		context.Background(),
		"SELECT * FROM supplier WHERE supp_id = $1",
		suppId,
	)

	err = row.Scan(
		&supplier.SuppId,
		&supplier.Postcode,
		&supplier.Name,
		&supplier.Password,
	)
	utils.CheckError(err)

	return &supplier
}
