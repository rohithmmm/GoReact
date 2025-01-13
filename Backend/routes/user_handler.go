package routes

import (
	"ClockMe/controllers"
	"ClockMe/models"
	"fmt"
	"github.com/labstack/echo/v4"
)

func AddTimesheet(e echo.Context) error {
	var timesheet *models.Timesheet
	if err := e.Bind(&timesheet); err != nil {
		return err
	}
	err := controllers.PostTimesheetController(timesheet, e)
	return err
}

func AddEmployee(e echo.Context) error {
	var employee *models.Employee
	if err := e.Bind(&employee); err != nil {
		fmt.Println(err)
	}
	if err := controllers.AddEmployeeController(employee, e); err != nil {
		return err
	}
	return nil
}

type EmailPassStruct struct {
	Email string `json:"email" sql:"email"`
	Pass  string `json:"password" sql:"password"`
}

func GetEmployee(e echo.Context) error {
	var emailPassStruct EmailPassStruct
	if err := e.Bind(&emailPassStruct); err != nil {
		fmt.Println(err)
	}
	if err := controllers.GetEmployeeController(emailPassStruct.Email, emailPassStruct.Pass, e); err != nil {
		return err
	}
	return nil
}
