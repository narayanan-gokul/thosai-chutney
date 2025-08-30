package shipment

import (
	"encoding/json"
	"thosai-chutney/utils"
	"thosai-chutney/core/cart"
	"net/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

func createShipment(conn *pgxpool.Pool) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		suppId, _ := request.Context().Value("userId").(int)

		var items []cart.CreateCartItem
		err := json.NewDecoder(request.Body).Decode(&items)
		utils.CheckError(err)
		defer request.Body.Close()

		CreateShipment(conn, suppId, items)
	}
}

func ShipmentRouter(conn *pgxpool.Pool) *http.ServeMux {
	shipmentRouter := http.NewServeMux()
	shipmentRouter.HandleFunc("/create", createShipment(conn))
	return shipmentRouter
}
