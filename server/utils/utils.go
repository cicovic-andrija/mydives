package utils

import (
	"strconv"
	"strings"
	"time"
)

func ConvertAndCheckID(strid string, max int) int {
	id, err := strconv.Atoi(strid)
	if err != nil || id < 1 || id > max {
		return 0
	}
	return id
}

func IsSpecialTag(tag string) bool {
	return strings.HasPrefix(tag, "_")
}

// ParseSpecialTag parses a special tag in the format "_key_value" and returns
// the key and value as two strings. If the tag doesn't match the format,
// returns empty strings for both key and value.
func ParseSpecialTag(tag string) (key string, value string) {
	if !strings.HasPrefix(tag, "_") {
		return "", ""
	}

	rest := tag[1:]

	idx := strings.Index(rest, "_")
	if idx == -1 || idx == 0 {
		return "", ""
	}

	key = rest[:idx]
	value = rest[idx+1:]
	return key, value
}

// DurationToYMD calculates the years, months, and days between two time points.
// Not super precise, works better for UTC.
func DurationToYMD(start time.Time, end time.Time) (years int, months int, days int) {
	if end.Before(start) {
		start, end = end, start
	}

	y1, m1, d1 := start.Date()
	y2, m2, d2 := end.Date()
	years = y2 - y1

	if m2 < m1 || (m2 == m1 && d2 < d1) {
		years--
	}

	months = int(end.Month()) - int(start.Month())
	if d2 < d1 {
		months--
	}
	if months < 0 {
		months += 12
	}

	newStart := start.AddDate(years, months, 0)
	days = int(end.Sub(newStart).Hours() / 24)

	return
}
