package routes

import (
	"apihex01/students/infraestructure"

	"github.com/gin-gonic/gin"
)

func StudentRouter(router *gin.Engine) {

	// Instanciamos los controladores
	createStudentController := infraestructure.CreateStudentController().Run
	getAllStudentsController := infraestructure.GetAllStudentsController().Run
	findByIdController := infraestructure.FindByIdController().Run
	updateStudentController := infraestructure.UpdateStudentController().Run
	deleteStudentController := infraestructure.DeleteStudentController().Run

	// Definimos el grupo de router
	routes := router.Group("/students")
	{
		// Definimos las rutas
		routes.POST("", createStudentController)
		routes.GET("", getAllStudentsController)
		routes.GET("/:id", findByIdController)
		routes.PUT("/:id", updateStudentController)
		routes.DELETE("/:id", deleteStudentController)
	}

}
