package distributor

import (
	"context"
	"thosai-chutney/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDistributor(conn *pgxpool.Pool, distributor *Distributor) int {
	tx, err := conn.Begin(context.Background())
	utils.CheckError(err)
	defer tx.Rollback(context.Background())

	row := tx.QueryRow(
		context.Background(),
		"INSERT INTO distributor(postcode, name, password) VALUES ($1, $2, $3) RETURNING dist_id",
		distributor.Postcode,
		distributor.Name,
		distributor.Password,
	)

	var dist_id int
	err = row.Scan(&dist_id)
	utils.CheckError(err)

	err = tx.Commit(context.Background())
	utils.CheckError(err)
	return dist_id
}
