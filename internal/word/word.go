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

type CreateWord struct {
	Name          string
	Sentences     []string
	DefinitionENG string
	DefinitionRUS string
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
}

// SetNextRepetition Устанавливает дату следующего повторения
func (w *Word) SetNextRepetition(now time.Time) {
	switch w.RepetitionNumber {
	case 0:
		w.RepetitionDate = now
	case 1:
		w.RepetitionDate = now.Add(time.Minute * minutesToAdd)
	case 2:
		w.RepetitionDate = now.AddDate(0, 0, daysToAdd)
	case 3:
		w.RepetitionDate = now.AddDate(0, 0, oneWeekDays*weeksToAdd)
	case 4:
		w.RepetitionDate = now.AddDate(0, monthsToAdd, 0)
	}
}

// MarkKnown устанавливает значение для следующего повторения
func (w *Word) MarkKnown() {
	w.RepetitionNumber += 1
}

// MarkUnknown сбрасывает значение для повторного изучения слова
func (w *Word) MarkUnknown() {
	w.RepetitionNumber = 1
}
