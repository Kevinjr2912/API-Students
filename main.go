package main

import (
	"apihex01/src/core"
	studentInfra "apihex01/src/students/infraestructure"
	studentRoutes "apihex01/src/students/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	studentInfra.InitDepedencies()

	// // Creamos el router
	r := gin.Default()

	// CORS
	core.ConfigCORS(r)
		
	studentRoutes.StudentRouter(r)

	// Levantamos el servidor
	r.Run()
}
