package word

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
