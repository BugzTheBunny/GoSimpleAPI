package server

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type App struct {
	mux         *http.ServeMux
	middlewares []Middleware
}

func (a *App) ListenAndServe(address string) error {
	return http.ListenAndServe(address, a.mux)
}

func NewApp() *App {
	return &App{
		mux:         http.NewServeMux(),
		middlewares: []Middleware{},
	}
}

func (a *App) Use(mw Middleware) {
	a.middlewares = append(a.middlewares, mw)
}

func (a *App) Handle(method string, path string, handler http.Handler, extraMiddleware ...Middleware) {
	finalHandler := handler
	handlerPath := fmt.Sprint(method, " ", path)
	fmt.Println(handlerPath)
	// Adding extra middleware that might be used.
	for i := len(extraMiddleware) - 1; i >= 0; i-- {
		finalHandler = extraMiddleware[i](finalHandler)
	}

	// Appending middlaware that were added by App.Use().
	for i := len(a.middlewares) - 1; i >= 0; i-- {
		finalHandler = a.middlewares[i](finalHandler)
	}

	a.mux.Handle(handlerPath, finalHandler)
}
