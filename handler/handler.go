package handler

import "github.com/doCodingWhileCoding/GO-API-CRUD/model"

// Storage .
type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetByID(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}
