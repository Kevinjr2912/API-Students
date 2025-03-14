package controllers

import (
	application "apihex01/src/students/application/useCases"
	"apihex01/src/students/infraestructure"
	"apihex01/src/students/infraestructure/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllStudentsController struct {
	useCase *application.GetAllStudents
}

func NewGetAllStudentsController() *GetAllStudentsController {
	mysql := infraestructure.GetMySQL()
	app := application.NewGetAllStudents(mysql)

	return &GetAllStudentsController{useCase: app}
}

func (gas_c *GetAllStudentsController) Run(ctx *gin.Context) {
	students, err := gas_c.useCase.Run()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	response := responses.NewResponseGetAllStudents(students)

	ctx.JSON(http.StatusOK, response)
}