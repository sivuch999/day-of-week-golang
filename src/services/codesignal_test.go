package services

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		param1 int
		param2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "successfully#01",
			args: args{
				param1: 1,
				param2: 3,
			},
			want: 4,
		},
		{
			name: "successfully#02",
			args: args{
				param1: 0,
				param2: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.param1, tt.args.param2); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCenturyFromYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "successfully#01",
			args: args{
				year: 999,
			},
			want: 10,
		},
		{
			name: "successfully#02",
			args: args{
				year: 1011,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CenturyFromYear(tt.args.year); got != tt.want {
				t.Errorf("CenturyFromYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPalindrome(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "successfully#01",
			args: args{
				inputString: "b",
			},
			want: true,
		},
		{
			name: "successfully#02",
			args: args{
				inputString: "baaba",
			},
			want: false,
		},
		{
			name: "successfully#03",
			args: args{
				inputString: "abcdadcba",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPalindrome(tt.args.inputString); got != tt.want {
				t.Errorf("CheckPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdjacentElementsProduct(t *testing.T) {
	type args struct {
		inputArray []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "successfully#01",
			args: args{
				inputArray: []int{1, 2, 3, 4, 5, 6},
			},
			want: 30,
		},
		{
			name: "successfully#02",
			args: args{
				inputArray: []int{6, 5, 4, 10, 5},
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdjacentElementsProduct(tt.args.inputArray); got != tt.want {
				t.Errorf("AdjacentElementsProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShapeArea(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "successfully#01",
			args: args{
				number: 2,
			},
			want: 5,
		},
		{
			name: "successfully#02",
			args: args{
				number: 9999,
			},
			want: 199940005,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShapeArea(tt.args.number); got != tt.want {
				t.Errorf("ShapeArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
