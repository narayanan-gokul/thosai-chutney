package distributor

import (
	"context"
	"fmt"
	"thosai-chutney/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDistributor(conn *pgxpool.Pool, distributor *Distributor) {
	tx, err := conn.Begin(context.Background())
	utils.CheckError(err)
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(
		context.Background(),
		"INSERT INTO distributor(postcode, name, password) VALUES ($1, $2, $3)",
		distributor.Postcode,
		distributor.Name,
		distributor.Password,
	)
	utils.CheckError(err)

	err = tx.Commit(context.Background())
	utils.CheckError(err)
}
