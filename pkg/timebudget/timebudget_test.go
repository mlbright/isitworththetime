package timebudget

import (
	"testing"
	"time"
)

func TestInvalidTimePeriod(t *testing.T) {
	saved := time.Duration(30 * time.Minute)
	frequency := 50
	interval := time.Duration(366 * DefaultWorkDay)
	period := time.Duration(DefaultYear)

	b, err := TimeBudget(&saved, uint(frequency), &interval, &period, DefaultUsers)
	if b != 0 || err != ErrNonsensicalTimePeriod {
		t.Log(err)
		t.Fatalf(`Expecting an error when the time period is smaller than the interval`)
	}
}

func TestTypical(t *testing.T) {
	saved := 30 * time.Minute
	frequency := uint(1)
	interval := DefaultWorkDay
	period := DefaultSpan

	want := (27 * 24 * time.Hour) + (2 * time.Hour)
	got, err := TimeBudget(&saved, frequency, &interval, &period, DefaultUsers)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if got != want {
		t.Fatalf("got %v, wanted %v", got, want)
	}
}

func FuzzTimeSaved(f *testing.F) {
	f.Add(uint(1), uint(1), uint(1), uint(1), uint(1))
	f.Fuzz(func(t *testing.T, saved uint, frequency uint, interval uint, period uint, users uint) {

		savedTime := time.Duration(saved)
		intervalTime := time.Duration(interval)
		periodTime := time.Duration(period)
		got, err := TimeBudget(&savedTime, frequency, &intervalTime, &periodTime, users)
		if err == ErrSavedLargerThanInterval || err == ErrNonsensicalTimePeriod || err == ErrDivideByZero {
			t.Skip()
		}
		if err != nil {
			t.Errorf("%q: %v", got, err)
		}
	})
}
