package helper

import (
	"testing"
)

func TestName(t *testing.T) {
	var tests = []struct {
		StartAt struct {
			Date, Hour, Minutes int
		}
		EndsAt struct {
			Date, Hour, Minutes int
		}
		Duration float64
	}{
		{
			StartAt: struct{ Date, Hour, Minutes int }{
				Date:    10,
				Hour:    1,
				Minutes: 10},
			EndsAt: struct{ Date, Hour, Minutes int }{
				Date:    10,
				Hour:    2,
				Minutes: 40},
			Duration: 90,
		},
		{
			StartAt: struct{ Date, Hour, Minutes int }{
				Date:    10,
				Hour:    1,
				Minutes: 10},
			EndsAt: struct{ Date, Hour, Minutes int }{
				Date:    10,
				Hour:    1,
				Minutes: 40},
			Duration: 30,
		},
		{
			StartAt: struct{ Date, Hour, Minutes int }{
				Date:    10,
				Hour:    1,
				Minutes: 10},
			EndsAt: struct{ Date, Hour, Minutes int }{
				Date:    12,
				Hour:    1,
				Minutes: 10},
			Duration: 2880,
		},
	}
	for _, v := range tests {
		t.Run("test duration", func(t *testing.T) {
			start := TaskTimeFormat(
				v.StartAt.Date,
				v.StartAt.Hour,
				v.StartAt.Minutes,
			)
			end := TaskTimeFormat(
				v.EndsAt.Date,
				v.EndsAt.Hour,
				v.EndsAt.Minutes,
			)
			if TaskDuration(start, end) != v.Duration {
				t.Errorf("XXXXX")
			}
		})
	}
}
