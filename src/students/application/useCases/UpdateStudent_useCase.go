package application

import (
	"apihex01/src/students/domain/entities"
	"apihex01/src/students/domain/repositories"
)

type UpdateStudent struct {
	db repositories.IStudent
}

func NewUpdateStudent(db repositories.IStudent) *UpdateStudent {
	return &UpdateStudent{db: db}
}

func (up *UpdateStudent) Run(id int64, student *entities.Student) (err error) {
	return up.db.UpdateStudent(id, student)
}