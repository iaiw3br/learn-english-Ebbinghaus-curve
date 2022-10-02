package word

import (
	"time"
)

type CreateWord struct {
	Name           string
	Example        []*Example
	DefinitionENG  string
	DefinitionRUS  string
	RepetitionDate time.Time
}

type Word struct {
	Name           string
	Example        []*Example
	DefinitionENG  string
	DefinitionRUS  string
	RepetitionDate time.Time
	IsKnown        bool
}

type Example struct {
	Sentence string
}

// convertToWord returns Word from CreateWord
func convertToWord(cw CreateWord) Word {
	return Word{
		Name:           cw.Name,
		Example:        cw.Example,
		DefinitionRUS:  cw.DefinitionRUS,
		DefinitionENG:  cw.DefinitionENG,
		RepetitionDate: cw.RepetitionDate,
		IsKnown:        false,
	}
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
