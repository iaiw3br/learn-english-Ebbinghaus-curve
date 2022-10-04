package word

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWord_Create(t *testing.T) {
	type fields struct {
		Name           string
		Example        []*Example
		DefinitionENG  string
		DefinitionRUS  string
		RepetitionDate time.Time
		IsKnown        bool
	}
	type args struct {
		cw CreateWord
	}
	now := time.Now()
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Word
	}{
		{
			name: "Create word",
			fields: fields{
				Name: "car",
				Example: []*Example{
					{
						Sentence: "I go to work by car.",
					},
					{
						Sentence: "Where did you park your car?",
					},
				},
				DefinitionENG:  "a vehicle with an engine",
				DefinitionRUS:  "машина, автомобиль",
				RepetitionDate: now,
				IsKnown:        false,
			},
			want: Word{
				Name: "car",
				Example: []*Example{
					{
						Sentence: "I go to work by car.",
					},
					{
						Sentence: "Where did you park your car?",
					},
				},
				DefinitionENG:    "a vehicle with an engine",
				DefinitionRUS:    "машина, автомобиль",
				RepetitionDate:   now,
				IsKnown:          false,
				RepetitionNumber: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := CreateWord{
				Name:           tt.fields.Name,
				Example:        tt.fields.Example,
				DefinitionENG:  tt.fields.DefinitionENG,
				DefinitionRUS:  tt.fields.DefinitionRUS,
				RepetitionDate: tt.fields.RepetitionDate,
			}
			word := convertToWord(w)

			assert.Equal(t, tt.want.Name, word.Name)
			assert.Equal(t, tt.want.Example, word.Example)
			assert.Equal(t, tt.want.DefinitionENG, word.DefinitionENG)
			assert.Equal(t, tt.want.DefinitionRUS, word.DefinitionRUS)
			assert.Equal(t, tt.want.RepetitionDate, word.RepetitionDate)
			assert.Equal(t, tt.want.IsKnown, word.IsKnown)
			assert.Equal(t, tt.want.RepetitionNumber, word.RepetitionNumber)
		})
	}
}

func TestWord_Know(t *testing.T) {
	type fields struct {
		Name    string
		IsKnown bool
	}
	tests := []struct {
		name     string
		fields   fields
		expected fields
	}{
		{
			name: "make is known",
			fields: fields{
				IsKnown: false,
			},
			expected: fields{
				IsKnown: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Word{
				Name:    tt.fields.Name,
				IsKnown: tt.fields.IsKnown,
			}
			w.Know()
			assert.Equal(t, tt.expected.IsKnown, w.IsKnown)
		})
	}
}

func TestWord_NotKnow(t *testing.T) {
	type fields struct {
		Date    time.Time
		IsKnown bool
	}
	now := time.Now()
	type args struct {
		now time.Time
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected fields
	}{
		{
			name: "is not known",
			fields: fields{
				IsKnown: true,
			},
			args: args{now: now},
			expected: fields{
				IsKnown: false,
				Date:    now,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Word{
				RepetitionDate: tt.fields.Date,
				IsKnown:        tt.fields.IsKnown,
			}
			w.NotKnow(tt.args.now)
			assert.Equal(t, tt.expected.IsKnown, w.IsKnown)
			assert.Equal(t, tt.expected.Date, w.RepetitionDate)
		})
	}
}
