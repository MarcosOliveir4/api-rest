package routes

import (
	"api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	r.GET("/", controllers.Home)
	r.GET("/api/personalidades", controllers.TodasPersonalidades)
	r.GET("/api/personalidades/:id", controllers.RetornaUmaPersonalidade)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
