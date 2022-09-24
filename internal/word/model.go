package word

import (
	"time"
)

type CreateWord struct {
	Name          string
	Example       []*Example
	DefinitionENG string
	DefinitionRUS string
	Date          time.Time
}

type Word struct {
	Name          string
	Example       []*Example
	DefinitionENG string
	DefinitionRUS string
	Date          time.Time
	IsKnown       bool
}

type Example struct {
	Sentence string
}

// Create returns new Word
func Create(cw CreateWord) Word {
	return Word{
		Name:          cw.Name,
		Example:       cw.Example,
		DefinitionRUS: cw.DefinitionRUS,
		DefinitionENG: cw.DefinitionENG,
		Date:          cw.Date,
		IsKnown:       false,
	}
}

// Know changes IsKnown = true
func (w *Word) Know() {
	w.IsKnown = true
}

// NotKnow changes IsKnown = false and Date
func (w *Word) NotKnow(now time.Time) {
	w.IsKnown = false
	w.Date = now
}
