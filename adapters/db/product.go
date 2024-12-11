package db

import (
	"database/sql"

	"github.com/gabriel01-jpg/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

var db *sql.DB

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from product where id = ?")
	if err != nil {
		return nil, err
	}

	stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	return &product, nil
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int

	p.db.QueryRow("SELECT COUNT(*) FROM product WHERE id = ?", product.GetID()).Scan(&rows)

	if rows == 0 {
		_, err := p.create(product)

		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)

		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {

	var stmt, err = p.db.Prepare("INSERT INTO product (id, name, price, status) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {

	var stmt, err = p.db.Prepare("UPDATE product SET name = ?, price = ?, status = ? WHERE id = ?")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}
