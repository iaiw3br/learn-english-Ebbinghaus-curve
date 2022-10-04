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
				RepetitionNumber: ZeroRepetition,
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
				RepetitionNumber: FirstRepetition,
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
				RepetitionNumber: SecondRepetition,
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
				RepetitionNumber: ThirdRepetition,
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
				RepetitionNumber: FourRepetition,
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
				RepetitionNumber: FourRepetition,
			},
			want: fields{
				RepetitionNumber: FirstRepetition,
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
				RepetitionNumber: ThirdRepetition,
			},
			want: fields{
				RepetitionNumber: FourRepetition,
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
