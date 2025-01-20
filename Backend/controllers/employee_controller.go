package controllers

import (
	"ClockMe/models"
	"ClockMe/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddEmployeeController(employee *models.Employee, c echo.Context) error {
	if err := services.AddEmployeeService(employee); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, employee)
}

func GetEmployeeController(email string, password string, c echo.Context) error {
	var first_name, token string
	if first_name, token = services.GetEmployeeService(email, password); token == "err" {
		return c.JSON(http.StatusForbidden, "Incorrect email or password")
	}
	if token == "" {
		return c.JSON(http.StatusForbidden, "Failed to create a token")
	}
	responseMap := map[string]interface{}{
		"email":      email,
		"first_name": first_name,
		"token":      token,
	}
	return c.JSON(http.StatusOK, responseMap)
}
