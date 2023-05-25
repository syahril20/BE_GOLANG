package main

import (
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))
	e.Logger.Fatal(e.Start("localhost:5000"))
	fmt.Println("Running :)")

}
