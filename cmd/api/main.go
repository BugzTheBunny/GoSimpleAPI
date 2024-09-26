package main

import (
	"github.com/BugzTheBunny/GoSimpleAPI/internal/db"
	"github.com/BugzTheBunny/GoSimpleAPI/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
