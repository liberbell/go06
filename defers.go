package main

import "fmt"

func main() {
	// defer fmt.Println("World.")

	// fmt.Println("Hello ")
	fmt.Println("counting")
	for x := 0; x < 5; x++ {
		defer fmt.Println(x)
	}
	fmt.Println("finished.")
}
