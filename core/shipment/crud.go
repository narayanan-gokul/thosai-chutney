package shipment

import (
	"context"
	"thosai-chutney/core/cart"
	"thosai-chutney/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateShipment(conn *pgxpool.Pool, suppId int, items []cart.CreateCartItem) {
	tx, err := conn.Begin(context.Background())
	utils.CheckError(err)
	defer tx.Rollback(context.Background())

	for _, item := range items {
		_, err :=  tx.Exec(
			context.Background(),
			"INSERT INTO shipment(item_id, quantity, supp_id) VALUES ($1, $2, $3)",
			item.ItemId,
			item.Quantity,
			suppId,
		)
		utils.CheckError(err)
	}

	err =  tx.Commit(context.Background())
	utils.CheckError(err)
}
