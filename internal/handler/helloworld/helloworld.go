package helloworld

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"netflixRental/internal/models"
)

func HelloWorld(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, models.HelloWorld{
		"Hello World!",
	})
}
