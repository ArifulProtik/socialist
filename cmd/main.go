package main

import (
	"github.com/ArifulProtik/socialist/internals/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e *echo.Echo
	c *controllers.Controller
}

func NewServer() *Server {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Static("/public", "public")
	return &Server{
		e: app,
		c: controllers.NewController(),
	}
}

func main() {
	server := NewServer()
	server.Routes()
	server.e.Logger.Fatal(server.e.Start(":1269"))
}
