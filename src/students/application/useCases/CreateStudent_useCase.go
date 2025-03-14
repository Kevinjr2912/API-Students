package application

import (
	"apihex01/src/students/domain/entities"
	"apihex01/src/students/domain/repositories"
)

type CreateStudent struct {
	db repositories.IStudent
}

func NewCreateStudent(db repositories.IStudent) *CreateStudent {
	return &CreateStudent{db: db}
}

func (cs *CreateStudent) Run(student *entities.Student) (id int64, err error) {
	return cs.db.CreateStudent(student)
}