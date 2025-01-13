package controllers

import (
	"ClockMe/models"
	"ClockMe/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetTimesheetController(c echo.Context) error {
	timesheet, err := services.GetTimesheet()
	if err != nil {
		fmt.Println("error:", err)
	}
	return c.JSON(http.StatusOK, timesheet)
}

func PostTimesheetController(timesheet *models.Timesheet, c echo.Context) error {
	if err := services.AddTimesheet(timesheet); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, timesheet)
}
