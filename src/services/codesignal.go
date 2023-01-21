package services

import (
	"math"
)

// 1) add
// Write a function that returns the sum of two numbers.
func Add(param1 int, param2 int) int {
	return param1 + param2
}

// 2) centuryFromYear
// Given a year, return the century it is in.
// The first century spans from the year 1 up to and including the year 100,
// the second - from the year 101 up to and including the year 200, etc.
func CenturyFromYear(year int) int {
	return int(math.Ceil(float64(year) / 100))
}

// 3) checkPalindrome
// Given the string, check if it is a palindrome.
func CheckPalindrome(inputString string) bool {
	// Convert string to array rune numbers
	r := []rune(inputString)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i] // Switch rune numbers
	}
	reversed := string(r)

	return inputString == reversed
}

// 4) adjacentElementsProduct
// Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.
func AdjacentElementsProduct(inputArray []int) int {
	maxNumber := 0

	for i := 0; i < len(inputArray)-1; i++ {
		compareNumber := inputArray[i] * inputArray[i+1]
		if i == 0 {
			maxNumber = compareNumber
			continue
		}
		if compareNumber > maxNumber {
			maxNumber = compareNumber
		}
	}

	return maxNumber
}

// 5) ShapeArea
// Below we will define an n-interesting polygon. Your task is to find the area of a polygon for a given n.
// A 1-interesting polygon is just a square with a side of length 1.
// An n-interesting polygon is obtained by taking the n - 1-interesting polygon and appending 1-interesting polygons to its rim, side by side. You can see the 1-, 2-, 3- and 4-interesting polygons in the picture below
func ShapeArea(number int) int {
	return int(math.Pow(float64(number), 2) + math.Pow((float64(number)-1), 2))
}
