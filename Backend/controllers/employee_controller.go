package controllers

import (
	"ClockMe/models"
	"ClockMe/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddEmployeeController(employee *models.Employee, c echo.Context) error {
	if err := services.AddEmployeeService(employee); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, employee)
}

func GetEmployeeController(email string, password string, c echo.Context) error {
	var first_name string
	var err error
	if first_name, err = services.GetEmployeeService(email, password); err != nil {
		return err
	}
	if first_name == "" {
		return c.JSON(http.StatusBadRequest, "Incorrect email or password")
	}
	return c.JSON(http.StatusOK, "Hello "+first_name)
}
