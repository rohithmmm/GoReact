package services

import (
	"ClockMe/database"
	"ClockMe/models"
	"fmt"
)

func GetTimesheet() (*models.Timesheet, error) {
	var timesheet models.Timesheet

	if err := db.DB.QueryRow("SELECT * FROM timesheet").Scan(&timesheet.Clock_in, &timesheet.Clock_out); err != nil {
		fmt.Println(err, "timesheet_service error")
	}

	return &timesheet, nil

}

func AddTimesheet(timesheet *models.Timesheet) error {
	_, err := db.DB.Query("INSERT INTO timesheet (clock_in, clock_out) VALUES ($1,$2)", timesheet.Clock_in, timesheet.Clock_out)
	if err != nil {
		return err

	}
	return nil
}
