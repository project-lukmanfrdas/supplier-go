package main

import (
	"supplier-app/config"
	_ "supplier-app/docs"
	"supplier-app/models"
	"supplier-app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Supplier API
// @version 1.0
// @description API untuk mengelola data suppliers dan contacts
// @host localhost:1323
// @BasePath /
// @schemes http
func main() {
	config.InitDB()
	config.DB.AutoMigrate(&models.Contact{}, &models.Supplier{})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.SupplierRoutes(e)
	routes.ContactRoutes(e)
	e.Static("/uploads", "uploads") // untuk akses gambar logo
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
