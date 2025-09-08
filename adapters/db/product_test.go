package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/leomoritz/fullcycle-arq-hexagonal-app/adapters/db"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

// Função parar criar o processo de conexão com o banco de dados
func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:") // o arquivo fica em memória apenas para teste (mais rápido)
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

func TestProductDb_Get(t *testing.T) {
	setup()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())

}
