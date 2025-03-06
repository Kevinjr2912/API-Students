package application

import (
	"apihex01/students/domain/entities"
	"apihex01/students/domain/repositories"
)

type GetAllStudents struct {
	db repositories.IStudent
}

func NewGetAllStudents(db repositories.IStudent) *GetAllStudents {
	return &GetAllStudents{db: db}
}

func (gas *GetAllStudents) Run() (studentsArray *[]entities.Student, err error) {
	return gas.db.GetAllStudents()
}