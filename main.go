package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	connStr := "postgres://user:password@localhost:5432/utilities_store"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM SmartphoneUtilities")
	if err != nil {
		log.Fatalf("Failed to execute query: %v\n", err)
	}
	defer rows.Close()

	fmt.Println("Data from SmartphoneUtilities:")
	for rows.Next() {
		var (
			utilityID                                                     int
			utilityName, brand, modelCompatibility, description, imageURL string
			price                                                         float64
			quantityInStock                                               int
			dateAdded                                                     time.Time // Note that we're using Go's time.Time type for the DATE column
		)

		err = rows.Scan(&utilityID, &utilityName, &brand, &modelCompatibility, &price, &quantityInStock, &description, &imageURL, &dateAdded)
		if err != nil {
			log.Fatalf("Failed to scan row: %v\n", err)
		}
		fmt.Printf("ID: %d, Name: %s, Brand: %s, Compatible Models: %s, Price: %f, Quantity in Stock: %d, Description: %s, Image URL: %s, Date Added: %s\n",
			utilityID, utilityName, brand, modelCompatibility, price, quantityInStock, description, imageURL, dateAdded)
	}

	// Check for errors from iterating over rows.
	if rows.Err() != nil {
		log.Fatalf("Error while iterating rows: %v\n", rows.Err())
	}
}
