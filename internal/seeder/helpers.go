package seeder

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func strPtr(s string) *string {
	s = strings.TrimSpace(s)

	if s == "" {
		return nil
	}

	return &s
}

func floatPtr(s string) *float64 {
	s = strings.TrimSpace(s)

	if s == "" {
		return nil
	}

	s = strings.ReplaceAll(s, ",", "")

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil
	}

	return &f
}

func intPtr(s string) *int {
	s = strings.TrimSpace(s)

	if s == "" {
		return nil
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}

	return &i
}

func parseFloat(s string) float64 {
	s = strings.TrimSpace(s)

	if s == "" {
		return 0
	}

	s = strings.ReplaceAll(s, ",", "")

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}

	return f
}

func parseInt(s string) int {
	s = strings.TrimSpace(s)

	if s == "" {
		return 0
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return i
}

func parseDate(s string) time.Time {
	s = strings.TrimSpace(s)

	if s == "" {
		return time.Time{}
	}

	layouts := []string{
		"2006-01-02T15:04Z07:00",
		time.RFC3339,
		"2006-01-02T15:04:05.000Z",
		"2006-01-02 15:04:05",
		"2006-01-02",
		"02-01-2006",
		"02/01/2006",
		"1/2/2006",
		"02-Jan-2006",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t
		}
	}

	panic(fmt.Sprintf("cannot parse date: %q", s))
}

func datePtr(s string) *time.Time {

	s = strings.TrimSpace(s)

	if s == "" {
		return nil
	}

	layouts := []string{
		time.RFC3339,
		"2006-01-02",
		"02-01-2006",
		"02/01/2006",
		"1/2/2006",
		"02-Jan-2006",
	}

	for _, layout := range layouts {

		t, err := time.Parse(layout, s)

		if err == nil {
			return &t
		}
	}

	return nil
}
