package storage

import (
	"database/sql"
	"fmt"

	"github.com/CristianPadilla/BDwithGoPostgreSQL/pkg/invoice"
	"github.com/CristianPadilla/BDwithGoPostgreSQL/pkg/invoiceHeader"
	"github.com/CristianPadilla/BDwithGoPostgreSQL/pkg/invoiceItem"
)

// object represents postgreSQL.product table and gotta implement interface product.Storage
type PsqlInvoice struct {
	db                   *sql.DB
	invoiceHeaderStorage invoiceHeader.Storage
	invoiceItemsStorage  invoiceItem.Storage
}

//NewPsqlInvoice used for working with postgres-invoice
func NewPsqlInvoice(db *sql.DB, h invoiceHeader.Storage, i invoiceItem.Storage) *PsqlInvoice {
	return &PsqlInvoice{db: db, invoiceHeaderStorage: h, invoiceItemsStorage: i}
}

//Create to insert to products table
func (p *PsqlInvoice) Create(m *invoice.Model) error {
	//tx es la transaccion
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	if err = p.invoiceHeaderStorage.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return err
	}

	if err = p.invoiceItemsStorage.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println("factura creada con exito")
	return tx.Commit()

}
