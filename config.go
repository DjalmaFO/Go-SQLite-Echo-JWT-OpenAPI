package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func configurar() {
	e = echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	var ok bool
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	jwtSecret, ok = os.LookupEnv("SECRET")
	if !ok {
		log.Fatal("Variável SECRET não encontrada")
	}

	serverPort, ok = os.LookupEnv("SERVER_PORT")
	if !ok {
		log.Fatal("Variável SERVER_PORT não encontrada")
	}

	_, err = os.Stat(nomeArquivo)
	if os.IsNotExist(err) {
		file, err = os.Create(nomeArquivo)
		if err != nil {
			log.Fatalf("Falha ao gerar arquivo %s: %s \n", nomeArquivo, err.Error())
		}
		newDB = true
	} else {
		file, err = os.Open(nomeArquivo)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	db, err = sqlx.Connect(driver, nomeArquivo)
	if err != nil {
		os.Remove(nomeArquivo)
		log.Fatal(err.Error())
	}

	if err := db.Ping(); err != nil {
		os.Remove(nomeArquivo)
		log.Fatal(err.Error())
	}

	if newDB {
		createTables()
		createFirstUser()
	}
}
