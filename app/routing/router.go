package routing

import (
	"simple-wallet/app"
	"simple-wallet/app/registry"
	"simple-wallet/app/routing/handlers"
	"simple-wallet/app/routing/handlers/accounts"
	"simple-wallet/app/routing/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Router(fiberApp *fiber.App, domain *registry.Domain, config app.Config) {

	apiGroup := fiberApp.Group("/api")
	apiGroup.Use(logger.New())

	apiRouteGroup(apiGroup, domain, config)
}

func apiRouteGroup(g fiber.Router, domain *registry.Domain, config app.Config) {

	g.Post("/login", handlers.Authenticate(domain.User, config))
	g.Post("/user", handlers.Register(domain.User))

	g.Get("/account/balance", middleware.AuthByBearerToken(config.Secret), accounts.BalanceEnquiry(domain.Account))
	g.Post("/account/deposit", middleware.AuthByBearerToken(config.Secret), accounts.Deposit(domain.Account))
	g.Post("/account/withdrawal", middleware.AuthByBearerToken(config.Secret), accounts.Withdraw(domain.Account))
	g.Post("/account/withdraw", middleware.AuthByBearerToken(config.Secret), accounts.Withdraw(domain.Account))

	g.Get("/account/statement", middleware.AuthByBearerToken(config.Secret), accounts.MiniStatement(domain.Transaction))
}
