package main

import (
	"github.com/abdulmanafc2001/mongodb/database"
	"github.com/abdulmanafc2001/mongodb/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	routes.UserRouter(router)

	router.Run(":9090")
}
