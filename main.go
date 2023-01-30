package main

import (
	"day_of_week_test/src/services"
	"fmt"
)

func main() {
	ExamDayOfWeek()
	ExamCodesignal()
}

func ExamDayOfWeek() {
	data := []services.Day{
		{StartDate: "1900-01-01", FullDate: "1900-01-01"}, // Monday
		{StartDate: "1900-01-01", FullDate: "1900-01-02"}, // Tuesday
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

func ExamCodesignal() {
	// Add
	fmt.Printf("Add: %v\n", services.Add(1, 3))   // 4
	fmt.Printf("Add: %v\n", services.Add(10, 10)) // 20
	fmt.Printf("Add: %v\n", services.Add(0, 1))   // 1

	// Century From Year
	fmt.Printf("centuryFromYear: %v\n", services.CenturyFromYear(1905)) // 20
	fmt.Printf("centuryFromYear: %v\n", services.CenturyFromYear(45))   // 1
	fmt.Printf("centuryFromYear: %v\n", services.CenturyFromYear(1900)) // 19

	// Check Palindrome
	fmt.Printf("isPalindrome: %v\n", services.CheckPalindrome("a"))       // true
	fmt.Printf("isPalindrome: %v\n", services.CheckPalindrome("aaba"))    // false
	fmt.Printf("isPalindrome: %v\n", services.CheckPalindrome("aa"))      // true
	fmt.Printf("isPalindrome: %v\n", services.CheckPalindrome("abacaba")) // true

	// Adjacent Elements Product
	fmt.Printf("adjacentElementsProduct: %v\n", services.AdjacentElementsProduct([]int{3, 6, -2, -5, 7, 3})) // 21
	fmt.Printf("adjacentElementsProduct: %v\n", services.AdjacentElementsProduct([]int{1, 2, 3, 0}))         // 6
	fmt.Printf("adjacentElementsProduct: %v\n", services.AdjacentElementsProduct([]int{1, 0, 1, 0, 1000}))   // 0

	//Shape Area
	fmt.Printf("shapeArea: %v\n", services.ShapeArea(1))    // 1
	fmt.Printf("shapeArea: %v\n", services.ShapeArea(2))    // 5
	fmt.Printf("shapeArea: %v\n", services.ShapeArea(3))    // 13
	fmt.Printf("shapeArea: %v\n", services.ShapeArea(4))    // 25
	fmt.Printf("ShapeArea: %v\n", services.ShapeArea(100))  // 19801
	fmt.Printf("shapeArea: %v\n", services.ShapeArea(8999)) // 161946005
}
