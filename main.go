package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"thosai-chutney/core/consumer"
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
	var p = consumer.Consumer{"Jane", "Doe", 2014, 1} 
	consumer.CreateConsumer(conn, p)
}
