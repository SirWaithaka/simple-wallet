package rest

import (
	"github.com/gofiber/fiber"
	fibermiddleware "github.com/gofiber/fiber/middleware"
	"wallet"

	"wallet/registry"
	"wallet/rest/api/accounts"
	"wallet/rest/api/middleware"
	"wallet/rest/api/users"
)

func Router(fiberApp *fiber.App, domain *registry.Domain, config wallet.Config) {

	apiGroup := fiberApp.Group("/api")
	apiGroup.Use(fibermiddleware.Logger())

	apiRouteGroup(apiGroup, domain, config)
}

func apiRouteGroup(g *fiber.Group, domain *registry.Domain, config wallet.Config) {
	g.Post("/login", users.Authenticate(domain.User))
	g.Post("/user", users.Register(domain.User))

	g.Get("/account/balance", middleware.AuthByBearerToken(config.Secret), accounts.BalanceEnquiry(domain.Account))
	g.Post("/account/deposit", middleware.AuthByBearerToken(config.Secret), accounts.Deposit(domain.Account))
	g.Post("/account/withdrawal", middleware.AuthByBearerToken(config.Secret), accounts.Withdraw(domain.Account))

	g.Get("/account/statement", middleware.AuthByBearerToken(config.Secret), accounts.MiniStatement(domain.Transaction))
}
