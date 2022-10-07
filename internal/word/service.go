package word

import "time"

// Create создаёт объект Word
func Create(cw CreateWord, now time.Time) Word {
	word := ConvertCreateToWord(cw)
	word.SetNextRepetition(now)

	return word
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
