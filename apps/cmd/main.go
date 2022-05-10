package main

import (
	"fmt"
	"time"

	"github.com/theovassiliou/gtime"
)

func main() {
	today, _ := time.Parse("2006-01-02", "2022-05-22")
	timestamp1, _ := time.Parse("2006-01-02", "2022-05-21")
	fmt.Println("Timestamp is from " + gtime.HFFDistanceApart(timestamp1, today))

	// Timestamp is from yesterday

	timestamp2, _ := time.Parse("2006-01-02", "2022-05-19")
	fmt.Println("Timestamp is from " + gtime.HFFDistanceApart(timestamp2, today))

	// Timestamp is from 3 days ago

}
