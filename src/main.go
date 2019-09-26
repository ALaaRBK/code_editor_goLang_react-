package main

import (
	"bytes"
	"db1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path"
	"schemas"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
)

var db *mongo.Database
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func compileAndRun(code string, stdin string) (string, string) {
	data := []byte(code)
	uploadDir := os.Getenv("UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = os.Getenv("PWD")
	}
	cppFilePath := path.Join(uploadDir, RandStringRunes(20)+".cpp")
	outFilePath := path.Join(uploadDir, RandStringRunes(20)+".out")
	stdinFilePath := path.Join(uploadDir, RandStringRunes(20)+".in")
	ioutil.WriteFile(cppFilePath, data, 0644)
	ioutil.WriteFile(stdinFilePath, []byte(stdin), 0644)
	out, err := exec.Command("g++", cppFilePath, "-o", outFilePath).CombinedOutput()
	if err != nil {
		fmt.Println("failed to compile \n%s", string(out))
		os.Remove(cppFilePath)
		return "", string(out)
	}
	subProcess := exec.Command(outFilePath)
	std, _ := subProcess.StdinPipe()
	stdoutt, _ := subProcess.StdoutPipe()
	subProcess.Start()
	io.WriteString(std, stdin)
	buf := new(bytes.Buffer)
	buf.ReadFrom(stdoutt)
	output := buf.String()
	subProcess.Wait()
	if err != nil {
		fmt.Println("failed to compile \n%s", string(out))
		os.Remove(cppFilePath)
		os.Remove(outFilePath)
		return "", string(out)
	}
	os.Remove(outFilePath)
	os.Remove(stdinFilePath)
	fmt.Println(string(output))
	return cppFilePath, string(output)
}
func submitHandler(ctx echo.Context) error {
	body := ctx.Request().Body
	decodedBody := schemas.Submission{CreatedAt: time.Now()}
	json.NewDecoder(body).Decode(&decodedBody)
	if strings.TrimSpace(decodedBody.CodeText) == "" {
		return ctx.String(400, "can process empty code text")
	}
	res, output := compileAndRun(decodedBody.CodeText, decodedBody.Stdin)
	if res == "" {
		return ctx.String(http.StatusBadRequest, output)
	}
	decodedBody.FilePath = res
	schemas.Save(db, decodedBody)
	return ctx.String(http.StatusOK, output)
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
	db = db1.DB
	startServer()
}
