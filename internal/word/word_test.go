package word

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

func TestWord_MarkUnknown(t *testing.T) {
	type fields struct {
		RepetitionNumber int
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			name: "Mark unknown",
			fields: fields{
				RepetitionNumber: 4,
			},
			want: fields{
				RepetitionNumber: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Word{
				RepetitionNumber: tt.fields.RepetitionNumber,
			}
			w.MarkUnknown()
			assert.Equal(t, tt.want.RepetitionNumber, w.RepetitionNumber)
		})
	}
}

func TestWord_MarkKnown(t *testing.T) {
	type fields struct {
		RepetitionNumber int
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			name: "Mark known",
			fields: fields{
				RepetitionNumber: 3,
			},
			want: fields{
				RepetitionNumber: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Word{
				RepetitionNumber: tt.fields.RepetitionNumber,
			}
			w.MarkKnown()
			assert.Equal(t, tt.want.RepetitionNumber, w.RepetitionNumber)
		})
	}
}
