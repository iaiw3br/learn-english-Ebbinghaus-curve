package list

import "tg-bot-learning-english/internal/word"

type CreateList struct {
	Title string
	Words []word.CreateWord
}

type List struct {
	Title string
	Words []word.Word
}

func (l *List) AddWord(w word.Word) {
	l.Words = append(l.Words, w)
}
