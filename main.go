package main

import (
	"day_of_week_test/src/services"
	"fmt"
)

func main() {
	data := []services.Day{
		{StartDate: "1900-01-01", FullDate: "1901-09-06"}, // Friday
		{StartDate: "1900-01-01", FullDate: "1981-11-12"}, // Thursday
		{StartDate: "1900-01-01", FullDate: "2004-01-23"}, // Friday
		{StartDate: "1900-01-01", FullDate: "2010-01-28"}, // Thursday
		{StartDate: "1900-01-01", FullDate: "2023-06-19"}, // Monday
		{StartDate: "1900-01-01", FullDate: "2023-07-02"}, // Sunday
		{StartDate: "1900-01-01", FullDate: "2023-01-20"}, // Friday
	}
	for _, v := range data {
		day, err := v.DayOfWeekService()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s: %s\n", v.FullDate, day)
	}
}
