package word

import (
	"time"
)

type CreateWord struct {
	Name          string
	Example       []*Example
	DefinitionENG string
	DefinitionRUS string
}

type Word struct {
	// Наименование слова
	Name string

	Example []*Example

	// Описание слова на английском
	DefinitionENG string

	// Описание слова на русском
	DefinitionRUS string

	// Дата следующего повторения
	RepetitionDate time.Time

	//
	IsKnown bool

	// Число повторений, в зависимости от этого значения назначается следующее повторение
	RepetitionNumber int
}

// Примеры предложений, где используется слово
type Example struct {
	Sentence string
}

// Know changes IsKnown = true
func (w *Word) Know() {
	w.IsKnown = true
}

// NotKnow changes IsKnown = false and Date
func (w *Word) NotKnow(now time.Time) {
	w.IsKnown = false
	w.RepetitionDate = now
}
