package main

import (
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	e := echo.New()

	mysql.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/uploads", "uploads")
	e.Logger.Fatal(e.Start("localhost:5000"))
	fmt.Println("Running :)")

}
