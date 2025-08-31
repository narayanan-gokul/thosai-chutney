package item

import (
	"context"
	"net/http"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
	"thosai-chutney/utils"
)

func getItems(conn *pgxpool.Pool) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		items := GetItems(conn)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(items)
	}
}

func addItems(conn *pgxpool.Pool) func (writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		tx, err := conn.Begin(context.Background())
		utils.CheckError(err)
		defer tx.Rollback(context.Background())

		var items []Item
		err = json.NewDecoder(request.Body).Decode(&items)
		utils.CheckError(err)
		defer request.Body.Close()

		for _, item := range items {
			_, err = tx.Exec(
				context.Background(),
				"INSERT INTO item(name, max_cap) VALUES ($1, $2)",
				item.Name,
				item.MaxCap,
			)
			utils.CheckError(err)
		}
		err = tx.Commit(context.Background())
		utils.CheckError(err)
	}
}

func ItemRouter(conn *pgxpool.Pool) *http.ServeMux {
	itemRouter := http.NewServeMux()
	itemRouter.HandleFunc("/", getItems(conn))
	itemRouter.HandleFunc("/add", addItems(conn))
	return itemRouter
}
