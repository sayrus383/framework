package repositories

import (
	"context"
	"database/sql"
	"gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/example/internal/models"
	"gitlab.bcc.kz/digital-banking-platform/microservices/currency-market/modules/framework.git/logger"
)

type SomeRepository struct {
	db  *sql.DB
	log logger.Logger
}

func NewSomeRepository(db *sql.DB, log logger.Logger) *SomeRepository {
	return &SomeRepository{
		db:  db,
		log: log,
	}
}

func (r *SomeRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	rows, err := r.db.QueryContext(ctx, "select id, phone_number, has_password from public.users")
	if err != nil {
		r.log.Error("SomeRepository.GetUsers QueryContext error", "err", err)
		return nil, err
	}
	defer rows.Close()

	var mapping []models.User

	for rows.Next() {
		var m models.User
		err := rows.Scan(&m.Id, &m.PhoneNumber, &m.HasPassword)
		if err != nil {
			r.log.Error("SomeRepository.GetUsers rows.Scan error", "err", err)
			return nil, err
		}
		mapping = append(mapping, m)
	}
	return mapping, nil
}

func (r *SomeRepository) GetMapping(ctx context.Context) ([]models.EndpointMapping, error) {
	rows, err := r.db.QueryContext(ctx, "select id, source, target, disabled, isauth from currency.endpoint_mapping")
	if err != nil {
		r.log.Error("SomeRepository.GetMapping QueryContext", "err", err)
		return nil, err
	}
	defer rows.Close()

	var mapping []models.EndpointMapping

	for rows.Next() {
		var m models.EndpointMapping
		err := rows.Scan(&m.Id, &m.Source, &m.Target, &m.Disabled, &m.IsAuth)
		if err != nil {
			r.log.Error("SomeRepository.GetMapping rows.Scan", "err", err)
			return nil, err
		}
		mapping = append(mapping, m)
	}
	return mapping, nil
}
