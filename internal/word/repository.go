package word

import (
	"context"

	"tg-bot-learning-english/pkg/client/postgresql"
)

type Store interface {
	Create(ctx context.Context, word Word) error
}

type repository struct {
	client postgresql.Client
}

func NewStore(client postgresql.Client) Store {
	return &repository{
		client: client,
	}
}

func (r *repository) Create(ctx context.Context, w Word) error {
	query := `
		INSERT INTO words (
		                   name, sentences, 
		                   definition_eng, definition_rus, 
		                   repetition_date, repetition_number
		                   )
		VALUES (
		        $1, $2, 
		        $3, $4, 
		        $5, $6)
	`

	_, err := r.client.Exec(ctx, query,
		w.Name, w.Sentences,
		w.DefinitionENG, w.DefinitionRUS,
		w.RepetitionDate, w.RepetitionNumber,
	)
	if err != nil {
		return err
	}
	return nil
}
