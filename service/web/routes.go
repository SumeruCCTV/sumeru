package web

import "github.com/gofiber/fiber/v2"

type RouteInitializer func(svc *Service, app *fiber.App)
type MiddlewareInitializer func(app *fiber.App)

var routes []RouteInitializer
var middlewares []MiddlewareInitializer

func Register(initializer RouteInitializer) {
	routes = append(routes, initializer)
}

func RegisterMiddleware(initializer MiddlewareInitializer) {
	middlewares = append(middlewares, initializer)
}

func registerRoutes(svc *Service, app *fiber.App) {
	for _, initializer := range routes {
		initializer(svc, app)
	}
}

func registerMiddleware(app *fiber.App) {
	for _, initializer := range middlewares {
		initializer(app)
	}
}
