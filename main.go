package main

import (
	"apihex01/core"
	studentInfra "apihex01/students/infraestructure"
	studentRoutes "apihex01/students/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	studentInfra.InitDepedencies()

	// // Creamos el router
	r := gin.Default()

	// CORS
	core.Init(r)

	studentRoutes.StudentRouter(r)

	// // Levantamos el servidor
	r.Run()
}
