package app

import (
	"context"
	"fmt"
	"gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/application"
	"gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/database"
	"gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/example/internal/repositories"
	"gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/logger"
	"time"
)

func Start(a application.App) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log := a.GetLogger()
	cfg := a.GetConfig()

	db, err := database.NewPostgresConn(cfg, log)
	if err != nil {
		return err
	}

	printRepoModels(ctx, db, log)

	return nil
}

func printRepoModels(ctx context.Context, db database.DB, log logger.Logger) {
	repo := repositories.NewSomeRepository(db.GetDB(), log)

	users, err := repo.GetUsers(ctx)
	if err != nil {
		log.Error("repo.GetUsers", "err", err)
		return
	}

	for _, u := range users {
		log.Info(
			fmt.Sprintf("USER #%d", u.Id),
			"Target", u.PhoneNumber,
			"HasPassword", u.HasPassword,
		)
	}
}
