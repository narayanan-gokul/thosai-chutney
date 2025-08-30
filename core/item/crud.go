package item

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"thosai-chutney/utils"
)

func GetItems(conn *pgxpool.Pool) []Item {
	var items []Item

	tx, err := conn.Begin(context.Background())
	utils.CheckError(err)
	defer tx.Rollback(context.Background())

	rows, err := tx.Query(
		context.Background(),
		"SELECT * FROM item",
	)
	defer rows.Close()
	utils.CheckError(err)

	for rows.Next() {
		var rowData Item
		err = rows.Scan(&rowData.ItemId, &rowData.Name, &rowData.MaxCap)
		utils.CheckError(err)
		items = append(items, rowData)
	}

	err = tx.Commit(context.Background())
	utils.CheckError(err)

	return items
}
