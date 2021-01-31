package storage

/*
import (
	"database/sql"
	"fmt"

	"github.com/doCodingWhileCoding/GO-API-CRUD/model"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	MigrateComunity = `CREATE TABLE IF NOT EXIST comunities (
			Name VARCHAR(50) NOT NULL,
			id INT AUTO_INCREMENT NOT NULL PRIMARY KEY
		)`
	CreateComunity   = `INSERT INTO comunities(name) VALUES(?)`
	GetAllComunities = `SELECT * FROM comunities`
	GetComunityByID  = GetAllComunities + " WHERE id = ?"
	UpdateComunity   = `UPDATE comunities SET name = ? WHERE id = ?`
	DeleteComunity   = `DELETE FROM comunities WHERE id = ?`
)

//MemoryComunities used for work with person-comunity
type MemoryComunities struct {
	db *sql.DB
}

//NewMemoryComunities return a new pointer of MySQLInvoiceItem
func NewMemoryComunities(db *sql.DB) *MemoryComunities {
	return &MemoryComunities{db}
}

//Migrate .
func (p *MemoryComunities) Migrate() error {
	stmt, err := p.db.Prepare(MigrateComunity)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("migración del InvoiceItemo ejecutada")
	return nil
}

//Create .
func (p *MemoryComunities) Create(m *model.Community) error {
	stmt, err := p.db.Prepare(CreateComunity)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(m.Name)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = uint(id)

	fmt.Printf("Se creo la comunidad correctamente con ID: %d\n", m.ID)
	return nil
}

//CreateTx .
func (p *MemoryComunities) CreateTx(tx *sql.Tx, m *model.Communities) error {
	stmt, err := p.db.Prepare(CreateComunity)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := 0; i < len(*m); i++ {
		result, err := stmt.Exec((*m)[i].Name)

		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		(*m)[i].ID = uint(id)
	}

	return nil
}

//GetAll implements the interface product.storage
func (p *MemoryComunities) GetAll() (model.Communities, error) {
	stmt, err := p.db.Prepare(GetAllComunities)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Communities, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

//GetByID implements the interface product.storage
func (p *MemoryComunities) GetByID(id uint) (*model.Community, error) {
	stmt, err := p.db.Prepare(GetComunityByID)
	if err != nil {
		return &model.Community{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

//Update implements the interface product.storage
func (p *MemoryComunities) Update(m *model.Community) error {
	stmt, err := p.db.Prepare(UpdateComunity)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.ID,
		m.Name,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con el id: %d", m.ID)
	}

	fmt.Println("se actualizo el producto correctamente")
	return nil
}

//Delete implements the interface product.storage
func (p *MemoryComunities) Delete(id uint) error {
	stmt, err := p.db.Prepare(DeleteComunity)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe la comunidad con el id: %d", id)
	}

	fmt.Println("se eliminó el producto correctamente")
	return nil
}
func scanRowProduct(s scanner) (*model.Community, error) {
	m := &model.Community{}
	err := s.Scan(
		&m.ID,
		&m.Name,
	)
	if err != nil {
		return &model.Community{}, err
	}
	return m, nil
}
*/
