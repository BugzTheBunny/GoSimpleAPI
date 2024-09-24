package server

import (
	"fmt"
	"net/http"

	"github.com/BugzTheBunny/GoSimpleAPI/internal/middleware"
)

type Route struct {
	Method          string
	Path            string
	Handler         http.HandlerFunc
	ExtraMiddleware []Middleware
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HOME")
}

func GetRoutes() []Route {
	return []Route{
		{
			Method:          "GET",
			Path:            "/home",
			Handler:         HelloWorldHandler,
			ExtraMiddleware: []Middleware{middleware.LoggingMiddleware},
		},
	}
}

func RegisterRoutes(app *App) {

	for _, route := range GetRoutes() {
		fmt.Println(route)
		app.Handle(route.Method, route.Path, route.Handler, route.ExtraMiddleware...)
	}

}
