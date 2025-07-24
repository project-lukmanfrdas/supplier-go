package routes

import (
	"supplier-app/handlers"

	"github.com/labstack/echo/v4"
)

func SupplierRoutes(e *echo.Echo) {
	e.POST("/suppliers", handlers.CreateSupplier)
	e.GET("/suppliers", handlers.GetSuppliers)
}
