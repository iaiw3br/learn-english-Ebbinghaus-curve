package word

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	type args struct {
		cw CreateWord
	}
	now := time.Now()
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
				RepetitionNumber: ZeroRepetition,
				RepetitionDate:   now,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Create(tt.args.cw, now), "Create(%v)", tt.args.cw)
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
				RepetitionNumber: ZeroRepetition,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, convertToWord(tt.args.cw), "convertToWord(%v)", tt.args.cw)
		})
	}
}

func TestMarkKnown(t *testing.T) {
	type args struct {
		w Word
	}
	tests := []struct {
		name string
		args args
		want Word
	}{
		{
			name: "Mark known",
			args: args{
				w: Word{
					RepetitionNumber: ThirdRepetition,
				},
			},
			want: Word{
				RepetitionNumber: FourRepetition,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MarkKnown(tt.args.w), "MarkKnown(%v)", tt.args.w)
		})
	}
}

func TestMarkUnknown(t *testing.T) {
	type args struct {
		w Word
	}
	tests := []struct {
		name string
		args args
		want Word
	}{
		{
			name: "Mark unknown",
			args: args{
				w: Word{
					RepetitionNumber: SecondRepetition,
				},
			},
			want: Word{
				RepetitionNumber: FirstRepetition,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MarkUnknown(tt.args.w), "MarkUnknown(%v)", tt.args.w)
		})
	}
}
