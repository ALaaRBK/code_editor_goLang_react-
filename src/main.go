package main

import (
	"db1"
	"fmt"
	"os"
)

func main() {
	dbURL := os.Getenv("MONGO_URL")
	dbName := os.Getenv("DB_NAME")
	m := map[string]interface{}{
		"MONGO_URL": dbURL,
		"DB_NAME":   dbName,
	}
	fmt.Println(m)
	db1.Init(m)
}
