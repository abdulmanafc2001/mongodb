package controllers

import (
	helper "github.com/abdulmanafc2001/mongodb/helpers"
	"github.com/abdulmanafc2001/mongodb/models"
	"github.com/gin-gonic/gin"
)

func AddMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.Bind(&movie); err != nil {
		c.JSON(500, gin.H{
			"error": "Binding error",
		})
		return
	}
	res, err := helper.AddOneMovie(movie)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully inserted movie",
		"id":      res,
	})
}
