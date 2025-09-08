package db_test

import (
	"database/sql"
	"log"
)

var Db *sql.DB

// Função parar criar o processo de conexão com o banco de dados
func setup() {
	Db, _ = sql.Open("slite3", ":memory:") // o arquivo fica em memória apenas para teste (mais rápido)
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products(
			"id" string,
			"name" string,
			"price" float,
			"status" string
			);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES("abc", "Product Test", 0, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}
