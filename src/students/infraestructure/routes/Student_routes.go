package routes

import (
	"apihex01/src/students/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRouter(router *gin.Engine) {

	routes := router.Group("/students") // Definimos el grupo de router
	{
		// Definimos las rutas
		routes.POST("", controllers.NewCreateStudentController().Run)
		routes.GET("", controllers.NewGetAllStudentsController().Run)
		routes.GET("/:id", controllers.NewFindByIdController().Run)
		routes.PUT("/:id", controllers.NewUpdateStudentController().Run)
		routes.DELETE("/:id", controllers.NewDeleteStudentController().Run)
	}

}
