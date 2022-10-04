package word

import (
	"time"
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

	//
	IsKnown bool

	// Число повторений, в зависимости от этого значения назначается следующее повторение
	RepetitionNumber int
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
