package services

import (
	db "ClockMe/database"
	"ClockMe/models"
	_ "database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"os"
	"time"

	_ "github.com/joho/godotenv"
	_ "github.com/labstack/echo-jwt/v4"
)

func AddEmployeeService(employee *models.Employee) error {

	var first_name string
	if err := db.DB.QueryRow("SELECT first_name from employee WHERE email = $1", employee.Email).Scan(&first_name); err == nil || first_name != "" {
		fmt.Println("Employee already exists")
		return errors.New("Employee already exists")
	}

	if _, err := db.DB.Query("INSERT INTO employee (first_name, last_name, email, password) VALUES ($1,$2,$3,$4)", employee.FirstName, employee.LastName, employee.Email, employee.Password); err != nil {
		return err
	}
	return nil
}

var TokenString string

func GetEmployeeService(email string, password string) (string, string) {
	var first_name string
	if err := db.DB.QueryRow("SELECT first_name from employee WHERE email = $1 AND password = $2", email, password).Scan(&first_name); err != nil {
		fmt.Println("Error in employee service")
		return "Incorrect email and password", "err"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"time":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	TokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "Failed to create a token", ""
	}

	return first_name, TokenString
}
