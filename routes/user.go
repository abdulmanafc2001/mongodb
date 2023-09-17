package routes

import (
	"github.com/abdulmanafc2001/mongodb/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.POST("/addmovie",controllers.AddMovie)
}
