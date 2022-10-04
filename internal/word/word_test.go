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

func TestWord_SetNextRepetition(t *testing.T) {
	type fields struct {
		RepetitionDate   time.Time
		RepetitionNumber int
	}
	type args struct {
		now time.Time
	}
	now := time.Now()
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name: "Zero RepetitionNumber",
			args: args{now: now},
			fields: fields{
				RepetitionDate:   now,
				RepetitionNumber: 0,
			},
			want: fields{
				RepetitionDate: now,
			},
		},
		{
			name: "One RepetitionNumber",
			args: args{now: now},
			fields: fields{
				RepetitionDate:   now,
				RepetitionNumber: 1,
			},
			want: fields{
				RepetitionDate: now.Add(time.Minute * minutesToAdd),
			},
		},
		{
			name: "Two RepetitionNumber",
			args: args{now: now},
			fields: fields{
				RepetitionDate:   now,
				RepetitionNumber: 2,
			},
			want: fields{
				RepetitionDate: now.AddDate(0, 0, daysToAdd),
			},
		},
		{
			name: "Three RepetitionNumber",
			args: args{now: now},
			fields: fields{
				RepetitionDate:   now,
				RepetitionNumber: 3,
			},
			want: fields{
				RepetitionDate: now.AddDate(0, 0, oneWeekDays*weeksToAdd),
			},
		},
		{
			name: "Four RepetitionNumber",
			args: args{now: now},
			fields: fields{
				RepetitionDate:   now,
				RepetitionNumber: 4,
			},
			want: fields{
				RepetitionDate: now.AddDate(0, monthsToAdd, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Word{
				RepetitionDate:   tt.fields.RepetitionDate,
				RepetitionNumber: tt.fields.RepetitionNumber,
			}
			w.SetNextRepetition(tt.args.now)
			assert.Equal(t, w.RepetitionDate, tt.want.RepetitionDate)
		})
	}
}
