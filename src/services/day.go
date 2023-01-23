package services

import (
	"strconv"
	"strings"
)

type ExtractDate struct {
	Year  int
	Month int
	Day   int
	// Hour   string
	// Minute string
	// Second string
}

type Day struct {
	StartDate string
	FullDate  string
}

type diffType string

const (
	DIFF_TYPE_MILLISECOND diffType = "millisecond"
	DIFF_TYPE_SECOND      diffType = "second"
	DIFF_TYPE_MINUTE      diffType = "minute"
	DIFF_TYPE_HOUR        diffType = "hour"
	DIFF_TYPE_DAY         diffType = "day"
)

func (day Day) DayOfWeekService() (string, error) {
	if day.StartDate == "" {
		day.StartDate = "1900-01-01"
	}

	// Difference between date
	diffDay, err := DiffDateService(day.StartDate, day.FullDate, "day")
	if err != nil {
		return "", err
	}

	// Since start date with Monday, Array must start with Monday as a result of modular 7
	dayIndex := diffDay % 7
	daysOfWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	return daysOfWeek[dayIndex], nil
}

func DiffDateService(startDate string, stopDate string, diffType diffType) (int, error) {
	extractD1, extractD2, err := splitAndVerifyDate(startDate, stopDate)
	if err != nil {
		return 0, err
	}

	days := 0
	// find day in year range
	for i := extractD1.Year; i < extractD2.Year; i++ {
		if isLeapYear(i) {
			days += 366
		} else {
			days += 365
		}
	}
	// find day in month range
	for i := extractD1.Month; i < extractD2.Month; i++ {
		days += daysInMonth(i, extractD2.Year)
	}
	// find day range
	days += extractD2.Day - extractD1.Day

	switch diffType {
	case DIFF_TYPE_MILLISECOND:
		return days * (1000 * 60 * 60 * 24), nil
	case DIFF_TYPE_SECOND:
		return days * (60 * 60 * 24), nil
	case DIFF_TYPE_MINUTE:
		return days * (60 * 24), nil
	case DIFF_TYPE_HOUR:
		return days * 24, nil
	case DIFF_TYPE_DAY:
		return days, nil
	default:
		return days * (1000 * 60 * 60 * 24), nil
	}
}

func splitAndVerifyDate(startDate string, stopDate string) (extractD1 ExtractDate, extractD2 ExtractDate, err error) {
	splitD1 := strings.Split(startDate, "-")
	splitD2 := strings.Split(stopDate, "-")

	year1, err := strconv.Atoi(splitD1[0])
	if err != nil {
		return extractD1, extractD2, err
	}
	month1, err := strconv.Atoi(splitD1[1])
	if err != nil {
		return extractD1, extractD2, err
	}
	day1, err := strconv.Atoi(splitD1[2])
	if err != nil {
		return extractD1, extractD2, err
	}

	year2, err := strconv.Atoi(splitD2[0])
	if err != nil {
		return extractD1, extractD2, err
	}
	month2, err := strconv.Atoi(splitD2[1])
	if err != nil {
		return extractD1, extractD2, err
	}
	day2, err := strconv.Atoi(splitD2[2])
	if err != nil {
		return extractD1, extractD2, err
	}

	extractD1 = ExtractDate{
		Year:  year1,
		Month: month1,
		Day:   day1,
	}

	extractD2 = ExtractDate{
		Year:  year2,
		Month: month2,
		Day:   day2,
	}

	return extractD1, extractD2, nil
}

func isLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}

func daysInMonth(month, year int) int {
	days := 31
	if month == 2 {
		if isLeapYear(year) {
			days = 29
		} else {
			days = 28
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		days = 30
	}
	return days
}
