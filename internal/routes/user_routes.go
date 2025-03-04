package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-go/internal/controllers"
)

func UserRoutes(g *echo.Group) {
	g.POST("/user", controllers.CreateUser)
	//g.GET("/user/:id", controllers.GetUser)
	//g.GET("/user", controllers.GetUsers)
	//g.PUT("/user/:id", UpdateUser)
	//g.DELETE("/user/:id", DeleteUser)
}
