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
	e.GET("/users", h.FindUser)
	e.GET("/users/:id", h.FindUserId)
	e.DELETE("/users/:id", h.DeleteUser)
	e.POST("/users", h.CreateUser)
	e.PATCH("/users/:id", h.UpdateUser)
}
