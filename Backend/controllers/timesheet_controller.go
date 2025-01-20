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

func ClockInController(email string, token string, c echo.Context) error {
	var clockInTime string
	if clockInTime = services.ClockInService(email, token); clockInTime == "error" {
		return c.JSON(http.StatusInternalServerError, "ClockIn Error")
	} else if clockInTime == "Invalid token" {
		return c.JSON(http.StatusForbidden, "Unauthorized access")
	} else if clockInTime == "Unauthorized" {
		return c.JSON(http.StatusForbidden, "Unauthorized access")
	}
	return c.JSON(http.StatusCreated, "Clocked in at "+clockInTime)
}

func ClockOutController(email string, token string, c echo.Context) error {
	var clockOutTime string
	if clockOutTime = services.ClockOutService(email, token); clockOutTime == "error" {
		return c.JSON(http.StatusInternalServerError, "ClockOut Error")
	} else if clockOutTime == "" {
		return c.JSON(http.StatusForbidden, "Unauthorized access")
	}
	return c.JSON(http.StatusCreated, "Clocked in at "+clockOutTime)
}
