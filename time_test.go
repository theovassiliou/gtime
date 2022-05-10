package gtime

import (
	"fmt"
	"testing"
	"time"
)

var today = time.Now()
var yesterday = today.AddDate(0, 0, -1)
var dby = yesterday.AddDate(0, 0, -1)

func TestIsToday(t *testing.T) {

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "today",
			args: args{
				t: today,
			},
			want: true,
		},
		{
			name: "yesterday",
			args: args{
				t: yesterday,
			},
			want: false,
		},
		{
			name: "day before yesterday",
			args: args{
				t: dby,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsToday(tt.args.t); got != tt.want {
				t.Errorf("IsToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsToday() {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	if IsToday(today) {
		fmt.Println("Is today")
	}

	if !IsToday(yesterday) {
		fmt.Println("Is not today")
	}

	// Output:
	// Is today
	// Is not today

}

func TestIsYesterday(t *testing.T) {
	type args struct {
		ref time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "today",
			args: args{
				ref: today,
			},
			want: false,
		},
		{
			name: "yesterday",
			args: args{
				ref: yesterday,
			},
			want: true,
		},
		{
			name: "day before yesterday",
			args: args{
				ref: dby,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsYesterday(tt.args.ref); got != tt.want {
				t.Errorf("IsYesterday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsYesterday() {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	if !IsYesterday(today) {
		fmt.Println("Today is not yesterday")
	}

	if IsYesterday(yesterday) {
		fmt.Println("Yesterday was yesterday")
	}

	// Output:
	// Today is not yesterday
	// Yesterday was yesterday

}

func TestIsDaysBefore(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
		n  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same day",
			args: args{
				t1: today,
				t2: today,
				n:  0,
			},
			want: true,
		},
		{
			name: "same day",
			args: args{
				t1: yesterday,
				t2: yesterday,
				n:  0,
			},
			want: true,
		},
		{
			name: "not before",
			args: args{
				t1: yesterday,
				t2: yesterday,
				n:  1,
			},
			want: false,
		},
		{
			name: "not before",
			args: args{
				t1: yesterday,
				t2: yesterday,
				n:  2,
			},
			want: false,
		},
		{
			name: "before",
			args: args{
				t1: yesterday,
				t2: today,
				n:  1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDaysBefore(tt.args.t1, tt.args.t2, tt.args.n); got != tt.want {
				t.Errorf("IsDaysBefore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsDaysBefore() {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	if IsDaysBefore(yesterday, today, 1) {
		fmt.Println("Yesterday is one day before today")
	}

	if IsDaysBefore(today, yesterday, -1) {
		fmt.Println("Yesterday is also one day after today")
	}

	// Output:
	// Yesterday is one day before today
	// Yesterday is also one day after today
}

func TestIsDaysAfter(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
		n  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same day",
			args: args{
				t1: today,
				t2: today,
				n:  0,
			},
			want: true,
		},
		{
			name: "same day",
			args: args{
				t1: yesterday,
				t2: yesterday,
				n:  0,
			},
			want: true,
		},
		{
			name: "after",
			args: args{
				t1: today,
				t2: yesterday,
				n:  1,
			},
			want: true,
		},
		{
			name: "not after",
			args: args{
				t1: yesterday,
				t2: today,
				n:  1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDaysAfter(tt.args.t1, tt.args.t2, tt.args.n); got != tt.want {
				t.Errorf("IsDaysAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsDaysAfter() {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	if IsDaysAfter(today, yesterday, 1) {
		fmt.Println("Today is one day afer yesterday")
	}

	if IsDaysAfter(yesterday, today, -1) {
		fmt.Println("Today is also one day before yesterday")
	}

	// Output:
	// Today is one day afer yesterday
	// Today is also one day before yesterday

}

func TestHFFDistanceApart(t *testing.T) {

	today, _ := time.Parse("2006-01-02", "2022-04-23")

	tests := []struct {
		want string
		ref  time.Time
	}{
		{"today", today.AddDate(0, 0, 0)},
		{"tomorrow", today.AddDate(0, 0, 1)},
		{"yesterday", today.AddDate(0, 0, -1)},
		{"day after tomorrow", today.AddDate(0, 0, 2)},
		{"day before yesterday", today.AddDate(0, 0, -2)},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := HFFDistanceApart(tt.ref, today); got != tt.want {
				t.Errorf("HFFDistanceApart() = %v, want %v", got, tt.want)
			}
		})
	}
}
func ExampleHFDistanceToday() {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	tomorrow := today.AddDate(0, 0, 1)

	fmt.Println(HFDistanceToday(yesterday))
	fmt.Println(HFDistanceToday(today))
	fmt.Println(HFDistanceToday(tomorrow))
	fmt.Println(HFDistanceToday(today.AddDate(0, 0, 2)))
	fmt.Println(HFDistanceToday(today.AddDate(0, 0, -2)))
	fmt.Println(HFDistanceToday(today.AddDate(0, 0, -3)))

	// Output: yesterday
	// today
	// tomorrow
	// day after tomorrow
	// day before yesterday
	// 3 days ago
}

func ExampleDaysApart() {

	today, _ := time.Parse("2006-01-02", "2022-04-23")
	yesterday := today.AddDate(0, 0, -1)
	tomorrow := today.AddDate(0, 0, 1)
	nextYear := today.AddDate(1, 0, 0)
	prevYear := today.AddDate(-1, 0, 0)

	fmt.Println(DaysApart(today, yesterday))
	fmt.Println(DaysApart(today, today))
	fmt.Println(DaysApart(today, tomorrow))
	fmt.Println(DaysApart(today, nextYear))
	fmt.Println(DaysApart(today, prevYear))

	// Output:
	// -1
	// 0
	// 1
	// 365
	// -365
}

func TestDaysApart(t *testing.T) {

	today, _ := time.Parse("2006-01-02", "2022-05-09 02:44")

	tests := []struct {
		name string
		ref  time.Time
		want int
	}{
		{"today", today, 0},
		{"tomorrow", today.AddDate(0, 0, 1), 1},
		{"yesterday", today.AddDate(0, 0, -1), -1},
		{"yesterday", today.AddDate(0, 0, -1).Add(14 * time.Hour), -1},
		{"next year", today.AddDate(1, 0, 0), 365},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DaysApart(today, tt.ref); got != tt.want {
				t.Errorf("DaysApart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHFDistanceToday(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HFDistanceToday(tt.args.t); got != tt.want {
				t.Errorf("HFDistanceToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absInt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := absInt(tt.args.x); got != tt.want {
				t.Errorf("absInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absDiffInt(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "negativ",
			args: args{
				x: 1,
				y: 4,
			},
			want: 3,
		},
		{
			name: "positiv",
			args: args{
				x: 4,
				y: 1,
			},
			want: 3,
		},
		{
			name: "negativ",
			args: args{
				x: -4,
				y: 2,
			},
			want: 6,
		},
		{
			name: "negativ",
			args: args{
				x: 2,
				y: -4,
			},
			want: 6,
		},
		{
			name: "zero",
			args: args{
				x: 2,
				y: 2,
			},
			want: 0,
		},
		{
			name: "zero",
			args: args{
				x: -2,
				y: -2,
			},
			want: 0,
		},
		{
			name: "zero",
			args: args{
				x: 0,
				y: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := absDiffInt(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("absDiffInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleHFFDistanceApart() {
	today, _ := time.Parse("2006-01-02", "2022-05-22")
	timestamp1, _ := time.Parse("2006-01-02", "2022-05-21")
	timestamp2, _ := time.Parse("2006-01-02", "2022-05-19")
	event1, _ := time.Parse("2006-01-02", "2022-05-23")
	event2, _ := time.Parse("2006-01-02", "2022-06-22")

	fmt.Println("timestamp1 is from " + HFFDistanceApart(timestamp1, today))
	fmt.Println("timestamp2 is from " + HFFDistanceApart(timestamp2, today))

	fmt.Println("event1 will happen " + HFFDistanceApart(event1, today))
	fmt.Println("event2 will happen " + HFFDistanceApart(event2, today))

	// Output:
	// timestamp1 is from yesterday
	// timestamp2 is from 3 days ago
	// event1 will happen tomorrow
	// event2 will happen 31 days
}
