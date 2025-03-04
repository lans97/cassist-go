package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lans97/cassist-go/internal/database"
	"github.com/lans97/cassist-go/internal/routes"
)

func main() {
	database.InitDB()
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/", "/var/www/html")

	api := e.Group("/api")

	routes.UserRoutes(api)

	e.GET("/*", func(c echo.Context) error {
		return c.File("/var/www/html/index.html")
	})

	e.Logger.Fatal(e.Start(":42069"))
}
