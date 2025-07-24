package routes

import (
	"supplier-app/handlers"

	"github.com/labstack/echo/v4"
)

func ContactRoutes(e *echo.Echo) {
	e.POST("/contacts", handlers.CreateContact)
	e.GET("/contacts/supplier/:supplier_id", handlers.GetContactsBySupplier)
	e.DELETE("/contacts/:id", handlers.DeleteContact)
}
