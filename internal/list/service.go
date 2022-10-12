package list

import (
	"context"
	"time"

	"tg-bot-learning-english/internal/word"
)

type Service struct {
	listStore Store
	wordStore word.Store
}

func NewService(listStore Store, wordStore word.Store) Service {
	return Service{
		listStore: listStore,
		wordStore: wordStore,
	}
}

func (s *Service) Create(cl CreateList) error {
	ctx := context.Background()

	list := cl.convertToList()
	listID, err := s.listStore.Create(ctx, list)
	if err != nil {
		return err
	}

	now := time.Now()
	for _, cw := range cl.Words {
		cw.ListID = listID

		w := cw.Create(now)
		err = s.wordStore.Create(ctx, w)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) Repeat(dateRepeat time.Time) (*[]List, error) {
	list, err := s.listStore.Repeat(context.Background(), dateRepeat)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (cw *CreateList) convertToList() List {
	return List{
		Title: cw.Title,
	}
}
