package routes

import (
	"log"

	"github.com/labstack/echo/v4"
)

type RouteInfo struct {
	Name string
	Func func(*echo.Echo)
}

var RouteRegistry = []RouteInfo{
	{Name: "SupplierRoutes", Func: SupplierRoutes},
	{Name: "ContactRoutes", Func: ContactRoutes},
	// Tambahkan route baru di sini
}

func RegisterRoute(name string, routeFunc func(*echo.Echo)) {
	RouteRegistry = append(RouteRegistry, RouteInfo{Name: name, Func: routeFunc})
}

func SetupAllRoutes(e *echo.Echo) {
	log.Println("🚀 Setting up all registered routes...")

	for i, route := range RouteRegistry {
		route.Func(e)
		log.Printf("✅ Route %d (%s) registered successfully", i+1, route.Name)
	}

	log.Println("🎯 All routes setup completed!")
}

func SetupSpecificRoutes(e *echo.Echo, routeFuncs ...func(*echo.Echo)) {
	log.Println("🎯 Setting up specific routes...")

	for _, routeFunc := range routeFuncs {
		routeFunc(e)
		log.Printf("✅ Specific route registered successfully")
	}

	log.Println("🎯 Specific routes setup completed!")
}

func GetRegisteredRouteNames() []string {
	var names []string
	for _, route := range RouteRegistry {
		names = append(names, route.Name)
	}
	return names
}

func PrintRegisteredRoutes() {
	log.Println("📋 Registered routes:")
	for i, route := range RouteRegistry {
		log.Printf("  %d. %s", i+1, route.Name)
	}
}
