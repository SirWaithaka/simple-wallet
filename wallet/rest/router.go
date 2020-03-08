package rest

import (
	"github.com/gofiber/fiber"
	"wallet/rest/api/account"
	"wallet/rest/api/user"
)


func Router(fiberApp *fiber.App) {

	apiGroup := fiberApp.Group("/api")

	apiRouteGroup(apiGroup)
}


func apiRouteGroup(g *fiber.Group) {
	g.Post("/login", user.Authenticate())
	g.Post("/user", user.Register())

	g.Post("/account/deposit", account.Deposit())
	g.Post("/account/withdrawal", account.Withdraw())
	g.Get("/account/statement", account.MiniStatement())
}