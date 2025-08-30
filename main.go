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

	router := http.NewServeMux()
	router.Handle("/", fileServer)
	router.Handle("/consumer/", http.StripPrefix("/consumer", consumerRouter))
	router.Handle("/supplier/", http.StripPrefix("/supplier", supplierRouter))
	router.Handle("/distributor/", http.StripPrefix("/distributor", distributorRouter))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Printf("Server up and listening on port %s.\n", server.Addr[1:])
	log.Fatal(server.ListenAndServe())
}
