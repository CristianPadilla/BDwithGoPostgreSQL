package storage

import (
	"database/sql"
	"fmt"

	"github.com/CristianPadilla/BDwithGoPostgreSQL/pkg/invoiceItem"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL, 
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP ,
		CONSTRAINT invoice_item_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_item_invoice_header_id_fk FOREIGN KEY (invoice_header_id)
		REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_item_product_id_fk FOREIGN KEY (product_id)
		REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`

	psqlCreateInvoiceItem = `INSERT INTO invoice_items (invoice_header_id,product_id)
	 VALUES ($1,$2) RETURNING id, created_at`
)

// object represents postgreSQL invoiceItem
type psqlInvoiceItem struct {
	db *sql.DB
}

//NewPsqlInvoiceItem constructor
func NewPsqlInvoiceItem(db *sql.DB) *psqlInvoiceItem {
	return &psqlInvoiceItem{db}
}

//Migrate implements interface  invoiceItem.storage
func (p *psqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
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

//CreateTx makes a transaction
func (p *psqlInvoiceItem) CreateTx(tx *sql.Tx, headerID uint, ms invoiceItem.Models) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, item := range ms {
		err = stmt.QueryRow(headerID, item.ProductID).Scan(&item.ID, &item.CreatedAt)
		if err != nil {
			return err
		}
	}
	return nil
	// return stmt.QueryRow(m.).Scan(&m.ID, &m.CreatedAt)
}
