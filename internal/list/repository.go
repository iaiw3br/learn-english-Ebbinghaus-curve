package list

import (
	"context"

	"tg-bot-learning-english/pkg/client/postgresql"
)

type Store interface {
	Create(ctx context.Context, list List) (int, error)
}

type repository struct {
	client postgresql.Client
}

func NewStore(client postgresql.Client) Store {
	return &repository{
		client: client,
	}
}

func (r *repository) Create(ctx context.Context, list List) (int, error) {
	query := `
		INSERT INTO lists (title)
		VALUES ($1)
		RETURNING id
	`

	var listID int
	err := r.client.QueryRow(ctx, query, list.Title).Scan(&listID)
	return listID, err
}
