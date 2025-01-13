package routes

import (
	"ClockMe/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/timesheet", controllers.GetTimesheetController)
	e.POST("/addTime", AddTimesheet)
	e.POST("/addEmployee", AddEmployee)
	e.POST("/getEmployeeByEmailAndPassword", GetEmployee)
}
