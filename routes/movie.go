package routes

import (
	"github.com/abdulmanafc2001/mongodb/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.POST("/addmovie",controllers.AddMovie)
	router.GET("/allmovies",controllers.ViewAllMovies)
	router.PUT("/updatemovie/:id",controllers.UpdateMovie)
	router.GET("/movie/:id",controllers.GetOneMovie)
	router.DELETE("/movie/:id",controllers.DeleteOneMovie)
	router.DELETE("/movies",controllers.DeleteAllMovies)
}
