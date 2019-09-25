package main

import (
	"db1"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/labstack/echo"
)

func indexHandler(ctx echo.Context) error {
	dat, err := ioutil.ReadFile("front-end/build/index.html")
	var data interface{}
	return ctx.Render(200, string(dat), &data)
}
func initRouter(server echo) {
	server.GET("/", indexHandler)
	// apiGroup := server.Group("/api")

}
func startServer() {
	server := echo.New()
	initRouter(&server)
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
