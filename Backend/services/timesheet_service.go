package services

import (
	"ClockMe/database"
	"ClockMe/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
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

var clockIn string
var emailK string

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Get the secret key
	secretKey := os.Getenv("SECRET")

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	// Handle parsing errors
	if err != nil {
		return nil, err
	}

	// Extract and return the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Optionally, check for expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, errors.New("token has expired")
			}
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ClockInService(email string, token string) string {
	emailK = email
	// Validate the token
	claims, err := ValidateToken(token)
	if err != nil {
		return "Invalid token"
	}

	// Verify that the email in the token matches the email provided
	if claims["email"] != email {
		return "Unauthorized"
	}
	clockIn = time.Now().Format("2006-01-02 15:04:05")
	_, err = db.DB.Exec("INSERT INTO timesheet (email, clock_in,clock_out) VALUES ($1,$2,$3)", email, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 ")+"17:01:00")
	if err != nil {
		return "error"
	}
	return clockIn
}

func ClockOutService(email string, token string) string {
	emailK = email
	// Validate the token
	claims, err := ValidateToken(token)
	if err != nil {
		return "Invalid token"
	}

	// Verify that the email in the token matches the email provided
	if claims["email"] != email {
		return "Unauthorized"
	}

	clockOut := time.Now().Format("2006-01-02 15:04:05")
	_, err = db.DB.Query("UPDATE timesheet SET email = $3, clock_out = $1 WHERE clock_in = $2", clockOut, clockIn, emailK)
	if err != nil {
		return "error"
	}
	return clockOut
}
