package word

func Create(cw CreateWord) Word {
	word := convertToWord(cw)

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
