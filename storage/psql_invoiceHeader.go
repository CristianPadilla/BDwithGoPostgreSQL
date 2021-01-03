package storage

import (
	"database/sql"
	"fmt"

	"github.com/CristianPadilla/BDwithGoPostgreSQL/pkg/invoiceHeader"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP ,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
	)`

	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers (client) VALUES ($1) RETURNING
	id, created_at`
)

// object represents postgreSQL invoice header
type PsqlInvoiceHeader struct {
	db *sql.DB
}

//NewPsqlInvoiceHeader constructor
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

//Migrate implements interface  invoiceHeader.storage
func (p *PsqlInvoiceHeader) Migrate() error {
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

//CreateTx makes a transaction
func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceHeader.Model) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
}
