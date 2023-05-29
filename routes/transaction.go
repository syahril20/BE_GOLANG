package routes

import (
	"dumbmerch/handler"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handler.HandlerTransaction(TransactionRepository)
	e.GET("/transaction", h.FindTransaction)
	e.GET("/transaction/:id", h.FindTransactionId)
	e.DELETE("/transaction/:id", h.DeleteTransaction)
	e.POST("/transaction", h.CreateTransaction)
	e.PATCH("/transaction/:id", h.UpdateTransaction)
}
