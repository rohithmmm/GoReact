package services

import (
	db "ClockMe/database"
	"ClockMe/models"
	_ "database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func AddEmployeeService(employee *models.Employee) error {
	if _, err := db.DB.Query("INSERT INTO employee (first_name, last_name, email, password) VALUES ($1,$2,$3,$4)", employee.FirstName, employee.LastName, employee.Email, employee.Password); err != nil {
		return err
	}
	return nil
}

func GetEmployeeService(email string, password string) (string, error) {
	var first_name string
	if err := db.DB.QueryRow("SELECT first_name from employee WHERE email = $1 AND password = $2", email, password).Scan(&first_name); err != nil {
		fmt.Println("Error in employee service")
		return "", err
	}
	return first_name, nil
}
