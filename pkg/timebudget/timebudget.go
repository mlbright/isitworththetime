package timebudget

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

const (
	DefaultWorkDay = 8 * time.Hour
	DefaultYear    = DefaultWorkDay * 5 * 52
	DefaultSpan    = 5 * DefaultYear
	DefaultUsers   = 1
)

var (
	ErrNonsensicalTimePeriod   error = errors.New("time span is smaller than the interval")
	ErrSavedLargerThanInterval error = errors.New("time saved is larger than the interval/period")
	ErrDivideByZero            error = errors.New("time interval must be greater than zero")
)

func TimeBudget(saved *time.Duration, frequency uint, interval *time.Duration, period *time.Duration, users uint) (time.Duration, error) {
	if *period < *interval {
		return 0, ErrNonsensicalTimePeriod
	}

	if time.Duration(frequency)**saved >= *interval {
		return 0, ErrSavedLargerThanInterval
	}

	if *interval == 0 {
		return 0, ErrDivideByZero
	}

	return *period / *interval * time.Duration(frequency) * *saved * time.Duration(users), nil
}

// https://gist.github.com/harshavardhana/327e0577c4fed9211f65
func HumanizeDuration(duration time.Duration) string {
	days := int64(duration.Hours() / 24)
	hours := int64(math.Mod(duration.Hours(), 24))
	minutes := int64(math.Mod(duration.Minutes(), 60))
	seconds := int64(math.Mod(duration.Seconds(), 60))

	chunks := []struct {
		singularName string
		amount       int64
	}{
		{"day", days},
		{"hour", hours},
		{"minute", minutes},
		{"second", seconds},
	}

	parts := []string{}

	for _, chunk := range chunks {
		switch chunk.amount {
		case 0:
			continue
		case 1:
			parts = append(parts, fmt.Sprintf("%d %s", chunk.amount, chunk.singularName))
		default:
			parts = append(parts, fmt.Sprintf("%d %ss", chunk.amount, chunk.singularName))
		}
	}

	return strings.Join(parts, " ")
}
