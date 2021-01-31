package model

// Community estructura de una comunidad
type Community struct {
	ID uint
	// Name nombre de una comunidad. Ej: EDteam
	Name string `json:"name"`
}

// Communities slice de comunidades
type Communities []*Community

//Storage .
/*type Storage interface {
	Create(community *Community) error
	Update(community *Community) error
	Delete(ID int) error
	GetByID(ID int) (Community, error)
	GetAll() (Communities, error)
}*/
