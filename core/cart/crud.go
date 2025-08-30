package cart

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"thosai-chutney/utils"
)

func CreateCart(conn *pgxpool.Pool, consId int, distId int, cartItems []CreateCartItem) error {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}

	for _, cartItem := range cartItems {
		_, err = tx.Exec(
			context.Background(),
			"INSERT INTO cart(cons_id, dist_id, item_id, quantity) VALUES ($1, $2, $3, $4) RETURNING SUM(quantity)",
			consId,
			distId,
			cartItem.ItemId,
			cartItem.Quantity,
		)
		if err != nil {
			return err
		}
		defer tx.Rollback(context.Background())
	}

	err = tx.Commit(context.Background())
	defer processDistCart(conn, distId)
	return err
}

func processDistCart(conn *pgxpool.Pool, distId int) {
	tx, err := conn.Begin(context.Background())
	utils.CheckError(err)
	defer tx.Rollback(context.Background())

	res, err := conn.Exec(
		context.Background(),
		"UPDATE cart SET quantity = (SELECT SUM(quantity) FROM cart WHERE dist_id IS NOT NULL AND dist_id = $1 AND cons_id IS NOT NULL AND NOT fulfilled)",
		distId,
	)
	utils.CheckError(err)
	if res.RowsAffected() == 0 {
		res, err = conn.Exec(
			context.Background(),
			"INSERT INTO cart(dist_id, quantity) VALUES ($1, (SELECT SUM(quantity) FROM cart WHERE dist_id = $1 AND cons_id IS NOT NULL AND NOT fulfilled))",
			distId,
		)
		utils.CheckError(err)
	}

	err = tx.Commit(context.Background())
	utils.CheckError(err)
}
