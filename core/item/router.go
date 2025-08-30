package item

import (
	"net/http"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
)

func getItems(conn *pgxpool.Pool) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		items := GetItems(conn)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(items)
	}
}

func ItemRouter(conn *pgxpool.Pool) *http.ServeMux {
	itemRouter := http.NewServeMux()
	itemRouter.HandleFunc("/", getItems(conn))
	return itemRouter
}
