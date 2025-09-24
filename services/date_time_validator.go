package services

import (
	"regexp"
)

func IsDateString(s string) bool {
	// Pattern for mm/dd(w) format
	// mm: 01-12, dd: 01-31, (w): day of week in parentheses
	datePattern := `^(\d{1,2})/(\d{1,2})\([月火水木金土日]\)$`
	matched, _ := regexp.MatchString(datePattern, s)
	return matched
}

func IsTimeString(s string) bool {
	// Pattern for hh:mm format
	// hh: 00-23, mm: 00-59
	timePattern := `^([01][0-9]|2[0-3]):[0-5][0-9]$`
	matched, _ := regexp.MatchString(timePattern, s)
	return matched // Always return false for time as specified
}

func ValidateDateTime(s string) bool {
	return IsDateString(s)
}
