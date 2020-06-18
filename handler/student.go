package handler

import (
	"net/http"

	"github.com/iammarkps/triamudom-room-api/models"
	"github.com/labstack/echo/v4"
)

// Student handler
func (handler *Handler) Student(c echo.Context) error {
	ID := c.Param("id")
	User := &models.User{}

	handler.DB.Where(&models.User{ID: ID}).First(User)

	if User == (&models.User{}) {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	return c.JSON(http.StatusOK, User)
}
