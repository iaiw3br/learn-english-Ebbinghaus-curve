package word

import "time"

func Create(cw CreateWord, now time.Time) Word {
	word := convertToWord(cw)
	word.SetNextRepetition(now)

	return word
}

// convertToWord returns Word from CreateWord
func convertToWord(cw CreateWord) Word {
	return Word{
		Name:          cw.Name,
		Sentences:     cw.Sentences,
		DefinitionRUS: cw.DefinitionRUS,
		DefinitionENG: cw.DefinitionENG,
	}
}
