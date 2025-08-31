package main

import (
	"context"
	"fmt"
	"os"
	"net/http"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"thosai-chutney/core/consumer"
	"thosai-chutney/core/supplier"
	"thosai-chutney/core/distributor"
	"thosai-chutney/core/item"
	"thosai-chutney/core/cart"
	"thosai-chutney/core/auth"
	"thosai-chutney/core/shipment"
	"thosai-chutney/core/allocate"
)

func main() {
	fmt.Println("Initializing database connection.")
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file: %v", err)
	}

	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to database!")

	fileServer := http.FileServer(http.Dir("./static"))

	consumerRouter := consumer.ConsumerRouter(conn)
	supplierRouter := supplier.SupplierRouter(conn)
	distributorRouter := distributor.DistributorRouter(conn)
	cartRouter := cart.CartRouter(conn)
	itemRouter := item.ItemRouter(conn)
	shipmentRouter := shipment.ShipmentRouter(conn)
	allocationRouter := allocate.AllocationRouter(conn)

	router := http.NewServeMux()
	router.Handle("/", fileServer)
	router.Handle("/consumer/", http.StripPrefix("/consumer", consumerRouter))
	router.Handle("/supplier/", http.StripPrefix("/supplier", supplierRouter))
	router.Handle("/distributor/", http.StripPrefix("/distributor", distributorRouter))
	router.Handle("/cart/", http.StripPrefix("/cart", auth.AuthMiddleware(cartRouter)))
	router.Handle("/item/", http.StripPrefix("/item", auth.AuthMiddleware(itemRouter)))
	router.Handle("/shipment/", http.StripPrefix("/shipment", auth.AuthMiddleware(shipmentRouter)))
	router.Handle("/allocation/", http.StripPrefix("/allocation", auth.AuthMiddleware(allocationRouter)))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Printf("Server up and listening on port %s.\n", server.Addr[1:])
	log.Fatal(server.ListenAndServe())
}
