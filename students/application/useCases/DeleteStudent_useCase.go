package application

import "apihex01/students/domain/repositories"

type DeleteStudent struct {
	db repositories.IStudent
}

func NewDeleteStudent(db repositories.IStudent) *DeleteStudent {
	return &DeleteStudent{db: db}
}

func (ds *DeleteStudent) Run(id int64) (err error) {
	return ds.db.DeleteStudent(id)
}