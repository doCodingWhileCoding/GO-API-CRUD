package storage

/*
import (
	"database/sql"

	"github.com/doCodingWhileCoding/GO-API-CRUD/model"
)

//MemoryContents used for work with person-comunity
type MemoryContents struct {
	db *sql.DB
}

//NewMemoryContents return a new pointer of MySQLInvoiceItem
func NewMemoryContents(db *sql.DB) *MemoryContents {
	return &MemoryContents{db}
}

//CreateTx implements the interface InvoiceItem.storage
func (p *MemoryContents) CreateTx(tx *sql.Tx, m *model.Communities, personID uint) error {
	stmt, err := tx.Prepare(mySQLCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		result, err := stmt.Exec(headerID, item.ProductID)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		item.ID = uint(id)

	}

	return nil
}*/
