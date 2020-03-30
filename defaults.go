package main

import (
	"fmt"
	"time"
)

// func main() {
// 	switch time.Now().Weekday() {
// 	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
// 		fmt.Println("It is a weekday.")
// 	default:
// 		fmt.Println("It is a weekend.")
// 	}
// }

func main() {
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("Before noon.")
	default:
		fmt.Println("Afternoon.")
	}
}
