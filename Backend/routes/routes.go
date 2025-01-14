package routes

import (
	"ClockMe/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/timesheet", controllers.GetTimesheetController)
	e.POST("/clockIn", ClockInHandler)
	e.POST("/clockOut", ClockOutHandler)
	e.POST("/addEmployee", AddEmployee)
	e.POST("/getEmployeeByEmailAndPassword", GetEmployee)
}
