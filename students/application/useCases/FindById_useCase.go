package application

import (
	"apihex01/students/domain/entities"
	"apihex01/students/domain/repositories"
)

type FindById struct {
	db repositories.IStudent
}

func NewFindById(db repositories.IStudent) *FindById {
	return &FindById{db: db}
}

func (fbi *FindById) Run(id int64) (student *entities.Student, err error) {
	return fbi.db.FindById(id)
}