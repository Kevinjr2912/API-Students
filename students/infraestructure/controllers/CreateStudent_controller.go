package controllers

import (
	"apihex01/students/application/services"
	application "apihex01/students/application/useCases"
	"apihex01/students/domain/entities"
	"apihex01/students/infraestructure/responses"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateStudentController struct {
	useCase *application.CreateStudent
	event *services.Event
}

func NewCreateStudentController(useCase *application.CreateStudent, event *services.Event) *CreateStudentController {
	return &CreateStudentController{useCase: useCase, event: event}
}

func (cs_c *CreateStudentController) Run(ctx *gin.Context) {
	var student entities.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if student.Name == "" && student.Age <= 0 && student.PhoneNumber <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Los campos están vacíos o son inválidos"})
		return
	}

	id, err := cs_c.useCase.Run(&student)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	// Definimos los encabezados
	ctx.Writer.Header().Set("Content-Type", "application/vnd.api+json")
	
	ctx.Writer.Header().Set("Location", fmt.Sprintf("http://localhost:8080/students/%d", id))

	// Reasignamos el id del estudiante
	student.Id = id

	// Enviamos la información de dicho estudiante al broker
	cs_c.event.Run(&student)

	response := responses.NewResponseStudentCreated(&student)

	ctx.JSON(http.StatusCreated, response)
}