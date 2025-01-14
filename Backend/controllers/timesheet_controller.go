package controllers

import (
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

func ClockInController(c echo.Context) error {
	var clockInTime string
	if clockInTime = services.ClockInService(); clockInTime == "error" {
		fmt.Println("Error")
		return nil
	}
	return c.JSON(http.StatusCreated, "Clocked in at "+clockInTime)
}

func ClockOutController(c echo.Context) error {
	var clockOutTime string
	if clockOutTime = services.ClockOutService(); clockOutTime == "error" {
		fmt.Println("Error")
		return nil
	}
	return c.JSON(http.StatusCreated, "Clocked out at "+clockOutTime)
}
