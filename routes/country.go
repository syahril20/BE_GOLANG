package routes

import (
	"dumbmerch/handler"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func CountryRoutes(e *echo.Group) {
	CountryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handler.HandlerCountry(CountryRepository)
	e.GET("/country", h.FindCountry)
	e.GET("/country/:id", h.FindCountryId)
	e.DELETE("/country/:id", h.DeleteCountry)
	e.POST("/country", h.CreateCountry)
	e.PATCH("/country/:id", h.UpdateCountry)
}
