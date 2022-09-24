package list

import "tg-bot-learning-english/internal/word"

type CreateList struct {
	Name  string
	Words []word.Word
}

type List struct {
	Name  string
	Words []word.Word
}

func Create(cl CreateList) List {
	return List{
		Name:  cl.Name,
		Words: cl.Words,
	}
}

func (l *List) AddWord(w word.Word) {
	l.Words = append(l.Words, w)
}
