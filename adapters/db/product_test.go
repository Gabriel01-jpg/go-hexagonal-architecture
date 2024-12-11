package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/gabriel01-jpg/go-hexagonal/adapters/db"
	"github.com/gabriel01-jpg/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")

	createTable(Db)
	createProduct(Db)

}

func createTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE product (
		"id" VARCHAR(36) PRIMARY KEY,
		"name" VARCHAR(255),
		"price" FLOAT,
		"status" VARCHAR(16)
	);`

	stmt, error := db.Prepare(createTableSQL)
	if error != nil {
		log.Fatal(error.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insertProductSQL := `INSERT INTO product (id, name, price, status) VALUES ("1", "Product 1", 10.5, "enabled");`
	stmt, err := db.Prepare(insertProductSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()

	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("1")
	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.5, product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()

	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = application.ENABLED

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

}
