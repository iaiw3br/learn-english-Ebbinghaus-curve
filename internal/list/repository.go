package list

import (
	"context"
	"time"

	"tg-bot-learning-english/internal/word"
	"tg-bot-learning-english/pkg/client/postgresql"
)

type Store interface {
	Create(ctx context.Context, list List) (int, error)
	Repeat(ctx context.Context, date time.Time) (*[]List, error)
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

func (r *repository) Repeat(ctx context.Context, date time.Time) (*[]List, error) {
	query := `
	SELECT l.title, 
	       w.name, w.sentences, w.definition_eng, w.definition_rus, 
	       w.repetition_number, w.repetition_date
	FROM lists l
	JOIN words w on l.id = w.list_id
	WHERE w.repetition_date <= $1
	`

	rows, err := r.client.Query(ctx, query, date)
	if err != nil {
		return nil, err
	}

	var list []List
	for rows.Next() {
		var l List
		var w word.Word

		err = rows.Scan(&l.Title,
			&w.Name, &w.Sentences, &w.DefinitionENG, &w.DefinitionRUS,
			&w.RepetitionNumber, &w.RepetitionDate,
		)
		if err != nil {
			return nil, err
		}

		var isWordAdded bool
		for i, nl := range list {
			if nl.Title == l.Title {
				list[i].Words = append(list[i].Words, w)
				isWordAdded = true
				break
			}
		}
		if !isWordAdded {
			l.Words = []word.Word{w}
			list = append(list, l)
		}
	}
	return &list, nil
}
