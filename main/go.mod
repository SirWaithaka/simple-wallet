module github.com/sirwaithaka/simple-wallet/main

go 1.13

require (
	github.com/gofiber/fiber/v2 v2.0.2
	wallet v1.0.0
)

replace wallet => ../wallet
