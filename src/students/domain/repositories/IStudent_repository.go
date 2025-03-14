package repositories

import "apihex01/src/students/domain/entities"

type IStudent interface {
	CreateStudent(student *entities.Student) (id int64, err error)
	GetAllStudents() (studentsArray *[]entities.Student, err error)
	FindById(id int64) (student *entities.Student, err error)
	UpdateStudent(id int64, student *entities.Student) (err error)
	DeleteStudent(id int64) (err error)
}