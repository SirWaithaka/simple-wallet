package rest

import (
	"github.com/gofiber/fiber"
	"wallet/registry"
	"wallet/rest/api/account"
	"wallet/rest/api/user"
)


func Router(fiberApp *fiber.App, domain *registry.Domain) {

	apiGroup := fiberApp.Group("/api")

	apiRouteGroup(apiGroup, domain)
}


func apiRouteGroup(g *fiber.Group, domain *registry.Domain) {
	g.Post("/login", user.Authenticate(domain.User))
	g.Post("/user", user.Register(domain.User))

	g.Post("/account/deposit", account.Deposit())
	g.Post("/account/withdrawal", account.Withdraw())
	g.Get("/account/statement", account.MiniStatement())
}