package storage

import (
	"database/sql"
	"fmt"

	"github.com/CristianPadilla/BDwithGoPostgreSQL/pkg/product"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(25),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP ,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`
	psqlCreateProduct = `INSERT INTO products(name, observations, price, created_at)
	VALUES ($1,$2,$3,$4) RETURNING id`

	psqlGetAllProduct = `SELECT id, name, observations, price, created_at, updated_at
	FROM products`

	psqlGetProductByID = `SELECT id, name, observations, price, created_at, updated_at
	FROM products ` + "WHERE id = $1 "

	psqlUpdateProduct = `UPDATE  products SET name=$1, observations=$2,
	 price=$3, updated_at=$4 ` + " WHERE id = $5 "

	psqlDeleteProduct = `DELETE FROM products WHERE id = $1 `
)

// object represents postgreSQL.product table and gotta implement interface product.Storage
type PsqlProduct struct {
	db *sql.DB
}

//NewPsqlProduct constructor
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

//Migrate implements interface Product.storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("migration executed")
	return nil
}

//Create to insert to products table
func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	//el scan sirve para recibir el id que me devuelve la sentencia en el RETURNING
	//el manejo de strings nulos stringToNulL() es necesario para leer campos que pueden ser
	// nulos en la bd, es decir que no tengan el NOT NULL
	err = stmt.QueryRow(m.Name, stringToNull(m.Observations), m.Price, m.CreatedAt).Scan(&m.ID)
	if err != nil {
		return err
	}
	fmt.Println("product created, id: ", m.ID)
	return nil

}

//GetAll implements storage interface
func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil

}

// GetByID to get a product
func (p *PsqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return scanRowProduct(stmt.QueryRow(id))
}

// Update to update a product
func (p *PsqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	//stringToNull en caso de que si me envian un valor vacio, lo guarde en la bd como null
	//y no como cadena vacia
	res, err := stmt.Exec(m.Name, stringToNull(m.Observations), m.Price, timeToNull(m.UpdatedAt), m.ID)
	if err != nil {
		return err
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAfected == 0 {
		return fmt.Errorf("no existe el id: %d en la tabla productos", m.ID)
	}
	fmt.Println("update succesfull")
	return nil
}

// Delete to delete a product
func (p *PsqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	fmt.Print("se elimino con exito")
	return nil

}
