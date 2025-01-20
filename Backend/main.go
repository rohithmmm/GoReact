package main

import (
	"ClockMe/database"
	"ClockMe/routes"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func main() {

	fmt.Println("Hello World")
	db.InitDB()
	defer db.DB.Close()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()
	routes.RegisterRoutes(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Add CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},                    // Allow only your frontend's origin
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE}, // Allow specific HTTP methods
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Logger.Fatal(e.Start(":8080"))

}
