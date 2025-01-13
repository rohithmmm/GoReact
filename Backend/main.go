package main

import (
	"ClockMe/database"
	"ClockMe/routes"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

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
	e.Logger.Fatal(e.Start(":8080"))

}
