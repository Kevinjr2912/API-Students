package controllers

import (
	application "apihex01/students/application/useCases"
	"apihex01/students/infraestructure/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllStudentsController struct {
	useCase *application.GetAllStudents
}

func NewGetAllStudentsController(useCase *application.GetAllStudents) *GetAllStudentsController {
	return &GetAllStudentsController{useCase: useCase}
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