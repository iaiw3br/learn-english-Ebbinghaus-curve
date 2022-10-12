package word

import (
	"context"
	"time"
)

type Service struct {
	wordStore Store
}

func NewService(wordStore Store) Service {
	return Service{
		wordStore: wordStore,
	}
}

// Create создаёт объект Word
func (s *Service) Create(ctx context.Context, cw CreateWord, now time.Time) error {
	word := cw.Create(now)

	err := s.wordStore.Create(ctx, word)
	if err != nil {
		return err
	}

	return nil
}

func MarkKnown(w Word) Word {
	w.MarkKnown()
	return w
}

func MarkUnknown(w Word) Word {
	w.MarkUnknown()
	return w
}

// ConvertCreateToWord возвращает Word из CreateWord
func ConvertCreateToWord(cw CreateWord) Word {
	return Word{
		Name:          cw.Name,
		Sentences:     cw.Sentences,
		DefinitionRUS: cw.DefinitionRUS,
		DefinitionENG: cw.DefinitionENG,
		ListID:        cw.ListID,
	}
}
