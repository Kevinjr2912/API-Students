package main

import (
	studentInfra "apihex01/src/students/infraestructure"
	studentRoutes "apihex01/src/students/infraestructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	studentInfra.InitDepedencies()

	// // Creamos el router
	r := gin.Default()

	// CORS
	r.Use(cors.Default())
		
	studentRoutes.StudentRouter(r)

	// Levantamos el servidor
	r.Run()
}
