package services

import (
	"testing"
)

func TestDay_DayOfWeekService(t *testing.T) {
	type args struct {
		StartDate string
		FullDate  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "successfully#01",
			args: args{
				StartDate: "1900-01-01",
				FullDate:  "1901-09-06",
			},
			want:    "Friday",
			wantErr: false,
		},
		{
			name: "successfully#02",
			args: args{
				StartDate: "1900-01-01",
				FullDate:  "2023-07-02",
			},
			want:    "Sunday",
			wantErr: false,
		},
		{
			name: "successfully#03",
			args: args{
				StartDate: "1900-01-01",
				FullDate:  "2023-06-19",
			},
			want:    "Monday",
			wantErr: false,
		},
		{
			name: "invalid date format",
			args: args{
				StartDate: "1900/01/01",
				FullDate:  "2023-06-19",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day := &Day{
				StartDate: tt.args.StartDate,
				FullDate:  tt.args.FullDate,
			}
			got, err := day.DayOfWeekService()

			if (err != nil) != tt.wantErr {
				t.Errorf("Day.DayOfWeekService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day.DayOfWeekService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffDateService(t *testing.T) {
	type args struct {
		startDate string
		stopDate  string
		diffType  string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "successfully#01",
			args: args{
				startDate: "1900-01-01",
				stopDate:  "1901-09-06",
				diffType:  "day",
			},
			want:    613,
			wantErr: false,
		},
		{
			name: "successfully#02",
			args: args{
				startDate: "1900-01-01",
				stopDate:  "2023-07-02",
				diffType:  "day",
			},
			want:    45107,
			wantErr: false,
		},
		{
			name: "successfully#03",
			args: args{
				startDate: "1900-01-01",
				stopDate:  "2023-06-19",
				diffType:  "day",
			},
			want:    45094,
			wantErr: false,
		},
		{
			name: "invalid date format",
			args: args{
				startDate: "1900/01/01",
				stopDate:  "2023-06-19",
				diffType:  "day",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DiffDateService(tt.args.startDate, tt.args.stopDate, "day")
			if (err != nil) != tt.wantErr {
				t.Errorf("DiffDateService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DiffDateService() = %v, want %v", got, tt.want)
			}
		})
	}
}
