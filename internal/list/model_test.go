package list

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"tg-bot-learning-english/internal/word"
)

func TestCreate(t *testing.T) {
	type args struct {
		cl CreateList
	}
	now := time.Now()
	tests := []struct {
		name string
		args args
		want List
	}{
		{
			name: "Create list",
			args: args{
				cl: CreateList{
					Name: "Transport",
					Words: []word.Word{
						{
							Name: "car",
							Example: []*word.Example{
								{
									Sentence: "I go to work by car.",
								},
								{
									Sentence: "Where did you park your car?",
								},
							},
							DefinitionENG: "a vehicle with an engine",
							DefinitionRUS: "машина, автомобиль",
							Date:          now,
							IsKnown:       false,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := Create(tt.args.cl)
			assert.Equal(t, list.Name, tt.args.cl.Name)
			assert.Equal(t, list.Words, tt.args.cl.Words)
		})
	}
}

func TestList_AddWord(t *testing.T) {
	type fields struct {
		Name  string
		Words []word.Word
	}
	type args struct {
		w word.Word
	}
	now := time.Now()
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name: "add new word to list",
			fields: fields{
				Name: "Transport",
				Words: []word.Word{
					{
						Name: "car",
						Example: []*word.Example{
							{
								Sentence: "I go to work by car",
							},
							{
								Sentence: "Where did you park your car?",
							},
						},
						DefinitionENG: "a vehicle with an engine",
						DefinitionRUS: "машина, автомобиль",
						Date:          time.Date(2022, 9, 24, 0, 0, 0, 0, time.Local),
						IsKnown:       false,
					},
				},
			},
			args: args{
				w: word.Word{
					Name: "bus",
					Example: []*word.Example{
						{
							Sentence: "a school bus",
						},
					},
					DefinitionENG: "a large vehicle that carries passengers by road, usually along a fixed route",
					DefinitionRUS: "автобус",
					Date:          now,
					IsKnown:       false,
				},
			},
			want: fields{
				Name: "Transport",
				Words: []word.Word{
					{
						Name: "car",
						Example: []*word.Example{
							{
								Sentence: "I go to work by car",
							},
							{
								Sentence: "Where did you park your car?",
							},
						},
						DefinitionENG: "a vehicle with an engine",
						DefinitionRUS: "машина, автомобиль",
						Date:          time.Date(2022, 9, 24, 0, 0, 0, 0, time.Local),
						IsKnown:       false,
					},
					{
						Name: "bus",
						Example: []*word.Example{
							{
								Sentence: "a school bus",
							},
						},
						DefinitionENG: "a large vehicle that carries passengers by road, usually along a fixed route",
						DefinitionRUS: "автобус",
						Date:          now,
						IsKnown:       false,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				Name:  tt.fields.Name,
				Words: tt.fields.Words,
			}
			l.AddWord(tt.args.w)
			assert.Equal(t, l.Words, tt.want.Words)
		})
	}
}
