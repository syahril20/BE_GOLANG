package routes

import (
	"dumbmerch/handler"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handler.HandlerUser(userRepository)
	e.GET("/users", h.FindUser)
	e.GET("/users/:id", middleware.Auth(h.FindUserId))
	e.GET("/user/:id", middleware.Auth(h.GetTransByUsers))
	e.DELETE("/users/:id", middleware.Auth(h.DeleteUser))
	e.POST("/users", middleware.Auth(h.CreateUser))
	e.PATCH("/users/:id", middleware.Auth(h.UpdateUser))
}
