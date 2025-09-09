package main

import (
	"database/sql"

	productDb "github.com/leomoritz/fullcycle-arq-hexagonal-app/adapters/db"
	"github.com/leomoritz/fullcycle-arq-hexagonal-app/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Instanciando o adapter
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := productDb.NewProductDb(db)

	// Criando o serviço utilizando adaptador de banco sqlite3 para persistência
	productService := application.NewProductService(productDbAdapter)

	// Salvando e ativando produto
	product, _ := productService.Create("Product Example", 30)
	productService.Enable(product)
}
