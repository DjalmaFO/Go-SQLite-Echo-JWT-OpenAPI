package main

import (
	"log"
	"os"
)

func createTables() {
	createTableUser()
}

func createTableUser() {
	stm, err := db.Prepare(createTableUsersSql)
	if err != nil {
		os.Remove(nomeArquivo)
		log.Fatalf("Falha ao criar tabela de usuários: Erro - %s", err.Error())
	}

	_, err = stm.Exec()
	if err != nil {
		os.Remove(nomeArquivo)
		log.Fatalf("Falha ao criar tabela de usuários: Erro - %s", err.Error())
	}
}

func createFirstUser() {
	name, ok := os.LookupEnv("FIRST_USER_NAME")
	if !ok {
		os.Remove(nomeArquivo)
		log.Fatal("Variável FIRST_USER_NAME não encontrada")
	}

	login, ok := os.LookupEnv("FIRST_USER_LOGIN")
	if !ok {
		os.Remove(nomeArquivo)
		log.Fatal("Variável FIRST_USER_LOGIN não encontrada")
	}

	password, ok := os.LookupEnv("FIRST_USER_PASSWORD")
	if !ok {
		os.Remove(nomeArquivo)
		log.Fatal("Variável FIRST_USER_PASSWORD não encontrada")
	}

	hash, err := HashPassword(password)
	if err != nil {
		os.Remove(nomeArquivo)
		log.Fatal(err.Error())
	}

	r, err := db.Exec(createUserSql, name, login, hash)
	if err != nil {
		os.Remove(nomeArquivo)
		log.Fatal(err.Error())
	}

	if ra, _ := r.RowsAffected(); ra == 0 {
		os.Remove(nomeArquivo)
		log.Fatal("Falha ao gerar novo usuário")
	}
}
