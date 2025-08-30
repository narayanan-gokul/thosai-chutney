package main

import (
	"context"
	"fmt"
	"os"
	"net/http"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

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

	router := http.NewServeMux()
	router.Handle("/", fileServer)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Printf("Server up and listening on port %s.\n", server.Addr[1:])
	log.Fatal(server.ListenAndServe())
}
