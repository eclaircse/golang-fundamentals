package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime) // complext date + time

	fmt.Println(presentTime.Format("01-02-2006"))
	// to print just the date, '-' ki jagah kuchh aur rakhenge
	// toh woh ayega output mein numbers ke bich mein

	fmt.Println(presentTime.Format("15:04:05"))
	// to print just the time, ':' ki jagah kuchh aur rakhenge
	// toh woh ayega output mein numbers ke bich mein
	
	fmt.Println(presentTime.Format("Monday"))

	fmt.Println(presentTime.Format("01-02-2006 15:04:06 Monday")) // all in one

	// in all these formatting, not necessary saare use kare
	// like, time mein seconds use nahi karenge toh chalega

	// creating a date

	createdDate := time.Date(2020, time.August, 12, 22, 23, 5, 0, time.UTC)
	// year, month, day, hour, minute, second, location

	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}