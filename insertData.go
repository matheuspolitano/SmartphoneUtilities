package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v4"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "utilities_store"
)

func randomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	for i := 0; i < 100; i++ {
		utilityName := "Utility " + randomString(rand.Intn(5)+5)
		brand := "Brand " + randomString(rand.Intn(3)+3)
		modelCompatibility := "Model " + randomString(rand.Intn(3)+3)
		price := rand.Float64() * 100
		quantityInStock := rand.Intn(100)
		description := "Description " + randomString(rand.Intn(10)+10)
		imageURL := "http://example.com/" + randomString(5) + ".jpg"

		_, err := conn.Exec(context.Background(),
			"INSERT INTO SmartphoneUtilities (utility_name, brand, model_compatibility, price, quantity_in_stock, description, image_url) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			utilityName, brand, modelCompatibility, price, quantityInStock, description, imageURL)

		if err != nil {
			log.Fatalf("Failed to insert data: %v\n", err)
		}
	}

	fmt.Println("Random data inserted successfully!")
}
