package main

import (
	"db1"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"schemas"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
)

type Code struct {
	Lang     string `json:"lang"`
	CodeText string `json:"codeText"`
	Stdin    string `json:stdin`
}

func submitHandler(ctx echo.Context) error {
	body := ctx.Request().Body
	decodedBody := Code{}
	err := json.NewDecoder(body).Decode(&decodedBody)
	schemas.JustTest()
	// err := json.Unmarshal(body, &decodedBody)
	fmt.Println(err)
	fmt.Println(decodedBody)
	return ctx.String(http.StatusOK, "safwan ok")
}
func startServer() {
	server := echo.New()
	assetHandler := http.FileServer(rice.MustFindBox("../front-end/build").HTTPBox())
	apiGroup := server.Group("/api")
	apiGroup.POST("/submit", submitHandler)
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
	// x := .Collection("submissions")
	startServer()
}
