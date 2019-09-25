package main

import (
	"fmt"
	"net/http"
	"os"

	"db1"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
)

func startServer() {
	server := echo.New()
	assetHandler := http.FileServer(rice.MustFindBox("../front-end/build").HTTPBox())
	// apiGroup := sever.Group("/api")
	// apiGroup.GET("/submit" )
	server.GET("/", echo.WrapHandler(assetHandler))
	server.GET("/static/*", echo.WrapHandler(assetHandler))
	server.Start(":8000")
}

func main() {
	dbURL := os.Getenv("MONGO_URL")
	dbName := os.Getenv("DB_NAME")
	m := map[string]interface{}{
		"MONGO_URL": dbURL,
		"DB_NAME":   dbName,
	}
	fmt.Println(m)
	db1.Init(m)
	startServer()
}
