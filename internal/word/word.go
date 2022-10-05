package word

import (
	"time"
)

const (
	oneWeekDays  = 7
	minutesToAdd = 30
	daysToAdd    = 1
	weeksToAdd   = 2
	monthsToAdd  = 2
)

const (
	ZeroRepetition = iota
	FirstRepetition
	SecondRepetition
	ThirdRepetition
	FourRepetition
)

type CreateWord struct {
	Name          string   `json:"name"`
	Sentences     []string `json:"sentences"`
	DefinitionENG string   `json:"definitionENG"`
	DefinitionRUS string   `json:"definitionRUS"`
	ListID        int      `json:"listID"`
}

type Word struct {
	// Наименование слова
	Name string

	// Примеры предложений, в которых используется слово
	Sentences []string

	// Описание слова на английском
	DefinitionENG string

	// Описание слова на русском
	DefinitionRUS string

	// Дата следующего повторения
	RepetitionDate time.Time

	// Число повторений, в зависимости от этого значения назначается следующее повторение
	RepetitionNumber int

	// Идентификатор списка
	ListID int
}

// SetNextRepetition устанавливает дату следующего повторения
func (w *Word) SetNextRepetition(now time.Time) {
	switch w.RepetitionNumber {
	case ZeroRepetition:
		w.RepetitionDate = now
	case FirstRepetition:
		w.RepetitionDate = now.Add(time.Minute * minutesToAdd)
	case SecondRepetition:
		w.RepetitionDate = now.AddDate(0, 0, daysToAdd)
	case ThirdRepetition:
		w.RepetitionDate = now.AddDate(0, 0, oneWeekDays*weeksToAdd)
	case FourRepetition:
		w.RepetitionDate = now.AddDate(0, monthsToAdd, 0)
	}
}

// MarkKnown устанавливает значение для следующего повторения
func (w *Word) MarkKnown() {
	w.RepetitionNumber += 1
}

// MarkUnknown сбрасывает значение для повторного изучения слова
func (w *Word) MarkUnknown() {
	w.RepetitionNumber = FirstRepetition
}
