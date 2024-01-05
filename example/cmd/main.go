package main

import (
	"gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/application"
	internapApp "gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/example/internal/app"
)

func main() {
	app := application.New()

	if err := app.Run(internapApp.Start); err != nil {
		app.GetLogger().Fatal("app error", "err", err)
	}
}
