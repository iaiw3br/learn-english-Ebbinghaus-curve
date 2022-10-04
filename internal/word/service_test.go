package word

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	type args struct {
		cw CreateWord
	}
	tests := []struct {
		name string
		args args
		want Word
	}{
		{
			name: "Create word",
			args: args{CreateWord{
				Name: "car",
				Sentences: []string{
					"I go to work by car",
					"Where did you park your car?",
				},
				DefinitionENG: "a vehicle with an engine",
				DefinitionRUS: "машина, автомобиль",
			}},
			want: Word{
				Name: "car",
				Sentences: []string{
					"I go to work by car",
					"Where did you park your car?",
				},
				DefinitionENG:    "a vehicle with an engine",
				DefinitionRUS:    "машина, автомобиль",
				IsKnown:          false,
				RepetitionNumber: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Create(tt.args.cw), "Create(%v)", tt.args.cw)
		})
	}
}

func Test_convertToWord(t *testing.T) {
	type args struct {
		cw CreateWord
	}
	tests := []struct {
		name string
		args args
		want Word
	}{
		{
			name: "Create word",
			args: args{CreateWord{
				Name: "car",
				Sentences: []string{
					"I go to work by car",
					"Where did you park your car?",
				},
				DefinitionENG: "a vehicle with an engine",
				DefinitionRUS: "машина, автомобиль",
			}},
			want: Word{
				Name: "car",
				Sentences: []string{
					"I go to work by car",
					"Where did you park your car?",
				},
				DefinitionENG:    "a vehicle with an engine",
				DefinitionRUS:    "машина, автомобиль",
				IsKnown:          false,
				RepetitionNumber: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, convertToWord(tt.args.cw), "convertToWord(%v)", tt.args.cw)
		})
	}
}
