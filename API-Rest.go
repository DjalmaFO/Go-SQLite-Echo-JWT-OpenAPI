package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

var newDB bool
var db *sqlx.DB
var e *echo.Echo
var file *os.File
var envFile string = ".env"
var driver string = "sqlite3"
var serverPort, jwtSecret string
var nomeArquivo string = "go_sqlite3.db"

func main() {
	configurar()
	defer db.Close()
	defer file.Close()

	configurarRotas()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", serverPort)))
}
