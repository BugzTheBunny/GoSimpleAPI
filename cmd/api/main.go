package main

import (
	"log"

	"github.com/BugzTheBunny/GoSimpleAPI/internal/middleware"
	"github.com/BugzTheBunny/GoSimpleAPI/internal/server"
)

func main() {
	app := server.NewApp()

	app.Use(middleware.LoggingMiddleware)

	server.RegisterRoutes(app)

	log.Fatal(app.ListenAndServe(":8080"))
}
