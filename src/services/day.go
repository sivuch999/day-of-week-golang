package services

import (
	"time"
)

type Day struct {
	StartDate string
	FullDate  string
}

type diffType string

const (
	DIFF_TYPE_HOUR diffType = "hour"
	DIFF_TYPE_DAY  diffType = "day"
)

func (day Day) DayOfWeekService() (string, error) {
	if day.StartDate == "" {
		day.StartDate = "1900-01-01"
	}

	// difference between date
	diffDay, err := DiffDateService(day.StartDate, day.FullDate, "day")
	if err != nil {
		return "", err
	}

	// since start date with Monday, Array must start with Monday as a result of modular 7
	dayIndex := diffDay % 7
	daysOfWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	return daysOfWeek[dayIndex], nil
}

func DiffDateService(startDate string, stopDate string, diffType diffType) (int, error) {
	d1, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return 0, err
	}
	d2, err := time.Parse("2006-01-02", stopDate)
	if err != nil {
		return 0, err
	}
	switch diffType {
	case DIFF_TYPE_HOUR:
		return int(d2.Sub(d1).Milliseconds() / (1000 * 60 * 60)), nil
	case DIFF_TYPE_DAY:
		return int(d2.Sub(d1).Milliseconds() / (1000 * 60 * 60 * 24)), nil
	default:
		return int(d2.Sub(d1).Milliseconds() / (1000 * 60 * 60 * 24)), nil
	}
}
