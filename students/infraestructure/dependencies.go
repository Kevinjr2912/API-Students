package infraestructure

import (
	"apihex01/students/application/services"
	application "apihex01/students/application/useCases"
	"apihex01/students/infraestructure/controllers"
)

var (
	myql *MySQL
	rabbit *Rabbit
)

func Init() {
	myql = NewMySQL()
	rabbit = NewRabbitMq()
}

// Casos de uso

// Crear un estudiante
func CreateStudentController() *controllers.CreateStudentController {
	useCaseCreateStudent := application.NewCreateStudent(myql)
	event := services.NewEvent(rabbit)

	return controllers.NewCreateStudentController(useCaseCreateStudent, event)
}

// Obtener todos los application
func GetAllStudentsController() *controllers.GetAllStudentsController {
	useCaseGetAllStudents := application.NewGetAllStudents(myql)

	return controllers.NewGetAllStudentsController(useCaseGetAllStudents)
}

// Encontrar un estudiante por medio del id
func FindByIdController() *controllers.FindByIdController {
	useCaseFindById := application.NewFindById(myql)

	return controllers.NewFindByIdController(useCaseFindById)
}

// Actualizar la información de un estudiante
func UpdateStudentController() *controllers.UpdateStudentController {
	useCaseUpdateStudent := application.NewUpdateStudent(myql)

	return controllers.NewUpdateStudentController(useCaseUpdateStudent)
}

// Eliminar un estudiante
func DeleteStudentController() *controllers.DeleteStudentController {
	useCaseDeleteStudent := application.NewDeleteStudent(myql)

	return controllers.NewDeleteStudentController(useCaseDeleteStudent)
}