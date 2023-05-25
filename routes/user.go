package routes

import (
	"dumbmerch/handler"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handler.HandlerUser(userRepository)
	e.GET("/user", h.FindUser)
	e.GET("/user/:id", h.FindUserId)
	e.DELETE("/user/:id", h.DeleteUser)
	e.POST("/user", h.CreateUser)
	// e.GET("/people/:id", handlers.GetPeople)
	// e.POST("/people", handlers.AddPeople)
	// e.DELETE("/people/:id", handlers.DeletePeople)
}
