package cart

import (
	"context"
	"net/http"
	"encoding/json"
	"thosai-chutney/utils"


	"github.com/jackc/pgx/v5/pgxpool"
)

func createConsumerCart(conn *pgxpool.Pool) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var createCartRequest CreateCartRequest
		err := json.NewDecoder(request.Body).Decode(&createCartRequest)
		utils.CheckError(err)

		tx, err := conn.Begin(context.Background())
		utils.CheckError(err)
		defer tx.Rollback(context.Background())

		userId, _ := request.Context().Value("userId").(int)

		err = CreateCart(conn, userId, createCartRequest.DistId, createCartRequest.Items)
		err = tx.Commit(context.Background())
		utils.CheckError(err)
		writer.WriteHeader(http.StatusCreated)
	}
}

func CartRouter(conn *pgxpool.Pool) *http.ServeMux {
	cartRouter := http.NewServeMux()
	cartRouter.HandleFunc("/create", createConsumerCart(conn))
	return cartRouter
}
