package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/lans97/cassist-go/internal/database"
	"github.com/lans97/cassist-go/internal/models"
)

// CRUD Operations for User

// Create a new user
func CreateUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(&user)

	res := database.DB.Create(&user)
	if res.Error != nil {
		return c.JSON(500, res.Error)
	}

	return c.JSON(200, user)
}
