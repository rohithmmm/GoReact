package services

import (
	"ClockMe/database"
	"ClockMe/models"
	"time"
)

func GetTimesheet() ([]models.Timesheet, error) {
	var timesheet models.Timesheet
	//if err := db.DB.QueryRow("SELECT * FROM timesheet").Scan(&timesheet.Clock_in, &timesheet.Clock_out); err != nil {
	//	fmt.Println(err, "timesheet_service error")
	//}

	res, err := db.DB.Query("SELECT * from timesheet")
	if err != nil {
		return nil, err
	}
	var times []models.Timesheet
	for res.Next() {
		err := res.Scan(&timesheet.Clock_in, &timesheet.Clock_out)
		if err != nil {
			return nil, err
		}
		times = append(times, timesheet)
	}
	return times, nil

}

func AddTimesheet(timesheet *models.Timesheet) error {
	_, err := db.DB.Query("INSERT INTO timesheet (clock_in, clock_out) VALUES ($1,$2)", timesheet.Clock_in, timesheet.Clock_out)
	if err != nil {
		return err
	}
	return nil
}

var clockIn string

func ClockInService() string {
	clockIn = time.Now().Format("2006-01-02 15:04:05")
	_, err := db.DB.Query("INSERT INTO timesheet (clock_in,clock_out) VALUES ($1,$2)", time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 ")+"17:01:00")
	if err != nil {
		return "error"
	}
	return clockIn
}

func ClockOutService() string {
	clockOut := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.DB.Query("UPDATE timesheet SET clock_out = $1 WHERE clock_in = $2", clockOut, clockIn)
	if err != nil {
		return "error"
	}
	return clockOut
}
