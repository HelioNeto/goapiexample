package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializando o Router
	router := gin.Default()

	initilizeRoutes(router)

	//Rodando o Servidor
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
