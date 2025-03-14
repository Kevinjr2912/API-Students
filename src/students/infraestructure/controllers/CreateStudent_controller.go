package controllers

import (
	"apihex01/src/students/application/services"
	application "apihex01/src/students/application/useCases"
	"apihex01/src/students/domain/entities"
	"apihex01/src/students/infraestructure"
	"apihex01/src/students/infraestructure/responses"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateStudentController struct {
	useCase *application.CreateStudent
	event *services.Event
}

func NewCreateStudentController() *CreateStudentController {
	// MySQL
	mysql := infraestructure.GetMySQL()
	app := application.NewCreateStudent(mysql)

	// Rabbit
	rabbit := infraestructure.GetRabbit()
	event := services.NewEvent(rabbit)

	return &CreateStudentController{useCase: app, event: event}
}

func (cs_c *CreateStudentController) Run(ctx *gin.Context) {
	var s entities.StudentCredentials

	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if s.Student.Name == "" && s.Student.Age <= 0 && s.Student.PhoneNumber <= 0 && s.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Los campos están vacíos o son inválidos"})
		return
	}

	id, err := cs_c.useCase.Run(&s.Student)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	// Definimos los encabezados
	ctx.Writer.Header().Set("Content-Type", "application/vnd.api+json")
	
	ctx.Writer.Header().Set("Location", fmt.Sprintf("http://localhost:8080/students/%d", id))

	// Reasignamos el id del estudiante
	s.Student.Id = id

	// Enviamos la información de dicho estudiante al exchange
	cs_c.event.Run(&s)

	response := responses.NewResponseStudentCreated(&s.Student)

	ctx.JSON(http.StatusCreated, response)
}