package controllers

import (
	"api-rest/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "Vlw natalina",
	})
}

func TodasPersonalidades(context *gin.Context) {
	context.JSON(http.StatusOK, models.Personalidades)
}

func RetornaUmaPersonalidade(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, personalidade := range models.Personalidades {
		idConvertido := strconv.Itoa(personalidade.Id)
		if idConvertido == id {
			ctx.JSON(http.StatusOK, personalidade)
		}
	}
}
