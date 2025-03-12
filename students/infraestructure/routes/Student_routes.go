package routes

import (
	"apihex01/students/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRouter(router *gin.Engine) {

	// Definimos el grupo de router
	routes := router.Group("/students")
	{
		// Definimos las rutas
		routes.POST("", controllers.NewCreateStudentController().Run)
		routes.GET("", controllers.NewGetAllStudentsController().Run)
		routes.GET("/:id", controllers.NewFindByIdController().Run)
		routes.PUT("/:id", controllers.NewUpdateStudentController().Run)
		routes.DELETE("/:id", controllers.NewDeleteStudentController().Run)
	}

}
